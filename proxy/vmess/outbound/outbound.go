package outbound

import (
	"v2ray.com/core/app"
	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"
	"v2ray.com/core/common/bufio"
	"v2ray.com/core/common/log"
	v2net "v2ray.com/core/common/net"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/retry"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/common/signal"
	"v2ray.com/core/proxy"
	"v2ray.com/core/proxy/vmess"
	"v2ray.com/core/proxy/vmess/encoding"
	"v2ray.com/core/transport/internet"
	"v2ray.com/core/transport/ray"
)

// VMessOutboundHandler is an outbound connection handler for VMess protocol.
type VMessOutboundHandler struct {
	serverList   *protocol.ServerList
	serverPicker protocol.ServerPicker
	meta         *proxy.OutboundHandlerMeta
}

// Dispatch implements OutboundHandler.Dispatch().
func (v *VMessOutboundHandler) Dispatch(target v2net.Destination, payload *buf.Buffer, ray ray.OutboundRay) {
	defer payload.Release()
	defer ray.OutboundInput().ForceClose()
	defer ray.OutboundOutput().Close()

	var rec *protocol.ServerSpec
	var conn internet.Connection

	err := retry.ExponentialBackoff(5, 100).On(func() error {
		rec = v.serverPicker.PickServer()
		rawConn, err := internet.Dial(v.meta.Address, rec.Destination(), v.meta.GetDialerOptions())
		if err != nil {
			return err
		}
		conn = rawConn

		return nil
	})
	if err != nil {
		log.Warning("VMess|Outbound: Failed to find an available destination:", err)
		return
	}
	log.Info("VMess|Outbound: Tunneling request to ", target, " via ", rec.Destination())

	command := protocol.RequestCommandTCP
	if target.Network == v2net.Network_UDP {
		command = protocol.RequestCommandUDP
	}
	request := &protocol.RequestHeader{
		Version: encoding.Version,
		User:    rec.PickUser(),
		Command: command,
		Address: target.Address,
		Port:    target.Port,
		Option:  protocol.RequestOptionChunkStream,
	}

	rawAccount, err := request.User.GetTypedAccount()
	if err != nil {
		log.Warning("VMess|Outbound: Failed to get user account: ", err)
	}
	account := rawAccount.(*vmess.InternalAccount)
	request.Security = account.Security

	defer conn.Close()

	conn.SetReusable(true)
	if conn.Reusable() { // Conn reuse may be disabled on transportation layer
		request.Option.Set(protocol.RequestOptionConnectionReuse)
	}

	input := ray.OutboundInput()
	output := ray.OutboundOutput()

	session := encoding.NewClientSession(protocol.DefaultIDHash)

	requestDone := signal.ExecuteAsync(func() error {
		defer input.ForceClose()

		writer := bufio.NewWriter(conn)
		defer writer.Release()

		session.EncodeRequestHeader(request, writer)

		bodyWriter := session.EncodeRequestBody(request, writer)
		defer bodyWriter.Release()

		if !payload.IsEmpty() {
			if err := bodyWriter.Write(payload); err != nil {
				return err
			}
		}
		writer.SetBuffered(false)

		if err := buf.PipeUntilEOF(input, bodyWriter); err != nil {
			return err
		}

		if request.Option.Has(protocol.RequestOptionChunkStream) {
			if err := bodyWriter.Write(buf.NewLocal(8)); err != nil {
				return err
			}
		}
		return nil
	})

	responseDone := signal.ExecuteAsync(func() error {
		defer output.Close()

		reader := bufio.NewReader(conn)
		defer reader.Release()

		header, err := session.DecodeResponseHeader(reader)
		if err != nil {
			return err
		}
		v.handleCommand(rec.Destination(), header.Command)

		conn.SetReusable(header.Option.Has(protocol.ResponseOptionConnectionReuse))

		reader.SetBuffered(false)
		bodyReader := session.DecodeResponseBody(request, reader)
		defer bodyReader.Release()

		if err := buf.PipeUntilEOF(bodyReader, output); err != nil {
			return err
		}

		return nil
	})

	if err := signal.ErrorOrFinish2(requestDone, responseDone); err != nil {
		log.Info("VMess|Outbound: Connection ending with ", err)
		conn.SetReusable(false)
	}

	return
}

// Factory is a proxy factory for VMess outbound.
type Factory struct{}

func (v *Factory) StreamCapability() v2net.NetworkList {
	return v2net.NetworkList{
		Network: []v2net.Network{v2net.Network_TCP, v2net.Network_KCP, v2net.Network_WebSocket},
	}
}

func (v *Factory) Create(space app.Space, rawConfig interface{}, meta *proxy.OutboundHandlerMeta) (proxy.OutboundHandler, error) {
	vOutConfig := rawConfig.(*Config)

	serverList := protocol.NewServerList()
	for _, rec := range vOutConfig.Receiver {
		serverList.AddServer(protocol.NewServerSpecFromPB(*rec))
	}
	handler := &VMessOutboundHandler{
		serverList:   serverList,
		serverPicker: protocol.NewRoundRobinServerPicker(serverList),
		meta:         meta,
	}

	return handler, nil
}

func init() {
	common.Must(proxy.RegisterOutboundHandlerCreator(serial.GetMessageType(new(Config)), new(Factory)))
}
