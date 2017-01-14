package freedom_test

import (
	"testing"

	"context"

	"v2ray.com/core/app"
	"v2ray.com/core/app/dispatcher"
	_ "v2ray.com/core/app/dispatcher/impl"
	"v2ray.com/core/app/dns"
	_ "v2ray.com/core/app/dns/server"
	"v2ray.com/core/app/proxyman"
	_ "v2ray.com/core/app/proxyman/outbound"
	"v2ray.com/core/app/router"
	"v2ray.com/core/common/buf"
	v2net "v2ray.com/core/common/net"
	"v2ray.com/core/proxy"
	. "v2ray.com/core/proxy/freedom"
	"v2ray.com/core/testing/assert"
	"v2ray.com/core/testing/servers/tcp"
	"v2ray.com/core/transport/internet"
	_ "v2ray.com/core/transport/internet/tcp"
	"v2ray.com/core/transport/ray"
)

func TestSinglePacket(t *testing.T) {
	assert := assert.On(t)

	tcpServer := &tcp.Server{
		MsgProcessor: func(data []byte) []byte {
			buffer := make([]byte, 0, 2048)
			buffer = append(buffer, []byte("Processed: ")...)
			buffer = append(buffer, data...)
			return buffer
		},
	}
	tcpServerAddr, err := tcpServer.Start()
	assert.Error(err).IsNil()

	space := app.NewSpace()
	ctx := app.ContextWithSpace(context.Background(), space)
	ctx = proxy.ContextWithOutboundMeta(ctx, &proxy.OutboundHandlerMeta{
		Address: v2net.AnyIP,
		StreamSettings: &internet.StreamConfig{
			Protocol: internet.TransportProtocol_TCP,
		},
	})
	freedom, err := New(ctx, &Config{})
	assert.Error(err).IsNil()
	assert.Error(space.Initialize()).IsNil()

	traffic := ray.NewRay()
	data2Send := "Data to be sent to remote"
	payload := buf.NewLocal(2048)
	payload.Append([]byte(data2Send))
	traffic.InboundInput().Write(payload)

	go freedom.Dispatch(tcpServerAddr, traffic)
	traffic.InboundInput().Close()

	respPayload, err := traffic.InboundOutput().Read()
	assert.Error(err).IsNil()
	assert.String(respPayload.String()).Equals("Processed: Data to be sent to remote")

	tcpServer.Close()
}

func TestIPResolution(t *testing.T) {
	assert := assert.On(t)

	space := app.NewSpace()
	ctx := app.ContextWithSpace(context.Background(), space)
	assert.Error(app.AddApplicationToSpace(ctx, new(proxyman.OutboundConfig))).IsNil()
	assert.Error(app.AddApplicationToSpace(ctx, new(dispatcher.Config))).IsNil()
	assert.Error(app.AddApplicationToSpace(ctx, new(router.Config))).IsNil()
	assert.Error(app.AddApplicationToSpace(ctx, &dns.Config{
		Hosts: map[string]*v2net.IPOrDomain{
			"v2ray.com": v2net.NewIPOrDomain(v2net.LocalHostIP),
		},
	})).IsNil()

	ctx = proxy.ContextWithOutboundMeta(ctx, &proxy.OutboundHandlerMeta{
		Address: v2net.AnyIP,
		StreamSettings: &internet.StreamConfig{
			Protocol: internet.TransportProtocol_TCP,
		},
	})
	freedom, err := New(ctx, &Config{DomainStrategy: Config_USE_IP})
	assert.Error(err).IsNil()

	assert.Error(space.Initialize()).IsNil()

	ipDest := freedom.ResolveIP(v2net.TCPDestination(v2net.DomainAddress("v2ray.com"), v2net.Port(80)))
	assert.Destination(ipDest).IsTCP()
	assert.Address(ipDest.Address).Equals(v2net.LocalHostIP)
}
