package socks

import (
	"context"
	"io"
	"runtime"
	"time"

	"v2ray.com/core/app"
	"v2ray.com/core/app/dispatcher"
	"v2ray.com/core/app/log"
	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"
	"v2ray.com/core/common/errors"
	"v2ray.com/core/common/net"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/signal"
	"v2ray.com/core/proxy"
	"v2ray.com/core/transport/internet"
	"v2ray.com/core/transport/internet/udp"
)

// Server is a SOCKS 5 proxy server
type Server struct {
	config *ServerConfig
}

// NewServer creates a new Server object.
func NewServer(ctx context.Context, config *ServerConfig) (*Server, error) {
	space := app.SpaceFromContext(ctx)
	if space == nil {
		return nil, errors.New("Socks|Server: No space in context.")
	}
	s := &Server{
		config: config,
	}
	return s, nil
}

func (s *Server) Network() net.NetworkList {
	list := net.NetworkList{
		Network: []net.Network{net.Network_TCP},
	}
	if s.config.UdpEnabled {
		list.Network = append(list.Network, net.Network_UDP)
	}
	return list
}

func (s *Server) Process(ctx context.Context, network net.Network, conn internet.Connection, dispatcher dispatcher.Interface) error {
	conn.SetReusable(false)

	switch network {
	case net.Network_TCP:
		return s.processTCP(ctx, conn, dispatcher)
	case net.Network_UDP:
		return s.handleUDPPayload(ctx, conn, dispatcher)
	default:
		return errors.New("Socks|Server: Unknown network: ", network)
	}
}

func (s *Server) processTCP(ctx context.Context, conn internet.Connection, dispatcher dispatcher.Interface) error {
	conn.SetReadDeadline(time.Now().Add(time.Second * 8))
	reader := buf.NewBufferedReader(conn)

	inboundDest, ok := proxy.InboundEntryPointFromContext(ctx)
	if !ok {
		return errors.New("Socks|Server: inbound entry point not specified.")
	}
	session := &ServerSession{
		config: s.config,
		port:   inboundDest.Port,
	}

	request, err := session.Handshake(reader, conn)
	if err != nil {
		if source, ok := proxy.SourceFromContext(ctx); ok {
			log.Access(source, "", log.AccessRejected, err)
		}
		log.Info("Socks|Server: Failed to read request: ", err)
		return err
	}
	conn.SetReadDeadline(time.Time{})

	if request.Command == protocol.RequestCommandTCP {
		dest := request.Destination()
		log.Info("Socks|Server: TCP Connect request to ", dest)
		if source, ok := proxy.SourceFromContext(ctx); ok {
			log.Access(source, dest, log.AccessAccepted, "")
		}

		return s.transport(ctx, reader, conn, dest, dispatcher)
	}

	if request.Command == protocol.RequestCommandUDP {
		return s.handleUDP()
	}

	return nil
}

func (*Server) handleUDP() error {
	// The TCP connection closes after v method returns. We need to wait until
	// the client closes it.
	// TODO: get notified from UDP part
	<-time.After(5 * time.Minute)

	return nil
}

func (v *Server) transport(ctx context.Context, reader io.Reader, writer io.Writer, dest net.Destination, dispatcher dispatcher.Interface) error {
	ctx, cancel := context.WithCancel(ctx)
	timeout := time.Second * time.Duration(v.config.Timeout)
	if timeout == 0 {
		timeout = time.Minute * 2
	}
	timer := signal.CancelAfterInactivity(ctx, cancel, timeout)

	ray, err := dispatcher.Dispatch(ctx, dest)
	if err != nil {
		return err
	}

	input := ray.InboundInput()
	output := ray.InboundOutput()

	requestDone := signal.ExecuteAsync(func() error {
		defer input.Close()

		v2reader := buf.NewReader(reader)
		if err := buf.PipeUntilEOF(timer, v2reader, input); err != nil {
			log.Info("Socks|Server: Failed to transport all TCP request: ", err)
			return err
		}
		return nil
	})

	responseDone := signal.ExecuteAsync(func() error {
		v2writer := buf.NewWriter(writer)
		if err := buf.PipeUntilEOF(timer, output, v2writer); err != nil {
			log.Info("Socks|Server: Failed to transport all TCP response: ", err)
			return err
		}
		return nil
	})

	if err := signal.ErrorOrFinish2(ctx, requestDone, responseDone); err != nil {
		log.Info("Socks|Server: Connection ends with ", err)
		input.CloseError()
		output.CloseError()
		return err
	}

	runtime.KeepAlive(timer)

	return nil
}

func (v *Server) handleUDPPayload(ctx context.Context, conn internet.Connection, dispatcher dispatcher.Interface) error {
	udpServer := udp.NewDispatcher(dispatcher)

	if source, ok := proxy.SourceFromContext(ctx); ok {
		log.Info("Socks|Server: Client UDP connection from ", source)
	}

	reader := buf.NewReader(conn)
	for {
		payload, err := reader.Read()
		if err != nil {
			return err
		}
		request, data, err := DecodeUDPPacket(payload.Bytes())

		if err != nil {
			log.Info("Socks|Server: Failed to parse UDP request: ", err)
			continue
		}

		if len(data) == 0 {
			continue
		}

		log.Info("Socks: Send packet to ", request.Destination(), " with ", len(data), " bytes")
		if source, ok := proxy.SourceFromContext(ctx); ok {
			log.Access(source, request.Destination, log.AccessAccepted, "")
		}

		dataBuf := buf.NewSmall()
		dataBuf.Append(data)
		udpServer.Dispatch(ctx, request.Destination(), dataBuf, func(payload *buf.Buffer) {
			defer payload.Release()

			log.Info("Socks|Server: Writing back UDP response with ", payload.Len(), " bytes")

			udpMessage := EncodeUDPPacket(request, payload.Bytes())
			defer udpMessage.Release()

			conn.Write(udpMessage.Bytes())
		})
	}
}

func init() {
	common.Must(common.RegisterConfig((*ServerConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return NewServer(ctx, config.(*ServerConfig))
	}))
}
