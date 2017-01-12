package scenarios

import (
	"net"
	"testing"

	"v2ray.com/core"
	"v2ray.com/core/app/router"
	v2net "v2ray.com/core/common/net"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/common/uuid"
	"v2ray.com/core/proxy/blackhole"
	"v2ray.com/core/proxy/dokodemo"
	"v2ray.com/core/proxy/freedom"
	"v2ray.com/core/proxy/vmess"
	"v2ray.com/core/proxy/vmess/inbound"
	"v2ray.com/core/proxy/vmess/outbound"
	"v2ray.com/core/testing/assert"
	"v2ray.com/core/testing/servers/tcp"
	"v2ray.com/core/transport/internet"
)

func TestPassiveConnection(t *testing.T) {
	assert := assert.On(t)

	tcpServer := tcp.Server{
		MsgProcessor: xor,
		SendFirst:    []byte("send first"),
	}
	dest, err := tcpServer.Start()
	assert.Error(err).IsNil()
	defer tcpServer.Close()

	serverPort := pickPort()
	serverConfig := &core.Config{
		Inbound: []*core.InboundConnectionConfig{
			{
				PortRange:              v2net.SinglePortRange(serverPort),
				ListenOn:               v2net.NewIPOrDomain(v2net.LocalHostIP),
				AllowPassiveConnection: true,
				Settings: serial.ToTypedMessage(&dokodemo.Config{
					Address: v2net.NewIPOrDomain(dest.Address),
					Port:    uint32(dest.Port),
					NetworkList: &v2net.NetworkList{
						Network: []v2net.Network{v2net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*core.OutboundConnectionConfig{
			{
				Settings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	assert.Error(InitializeServerConfig(serverConfig)).IsNil()

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: int(serverPort),
	})
	assert.Error(err).IsNil()

	{
		response := make([]byte, 1024)
		nBytes, err := conn.Read(response)
		assert.Error(err).IsNil()
		assert.String(string(response[:nBytes])).Equals("send first")
	}

	payload := "dokodemo request."
	{

		nBytes, err := conn.Write([]byte(payload))
		assert.Error(err).IsNil()
		assert.Int(nBytes).Equals(len(payload))
	}

	{
		response := make([]byte, 1024)
		nBytes, err := conn.Read(response)
		assert.Error(err).IsNil()
		assert.Bytes(response[:nBytes]).Equals(xor([]byte(payload)))
	}

	assert.Error(conn.Close()).IsNil()

	CloseAllServers()
}

func TestProxy(t *testing.T) {
	assert := assert.On(t)

	tcpServer := tcp.Server{
		MsgProcessor: xor,
	}
	dest, err := tcpServer.Start()
	assert.Error(err).IsNil()
	defer tcpServer.Close()

	serverUserID := protocol.NewID(uuid.New())
	serverPort := pickPort()
	serverConfig := &core.Config{
		Inbound: []*core.InboundConnectionConfig{
			{
				PortRange: v2net.SinglePortRange(serverPort),
				ListenOn:  v2net.NewIPOrDomain(v2net.LocalHostIP),
				Settings: serial.ToTypedMessage(&inbound.Config{
					User: []*protocol.User{
						{
							Account: serial.ToTypedMessage(&vmess.Account{
								Id: serverUserID.String(),
							}),
						},
					},
				}),
			},
		},
		Outbound: []*core.OutboundConnectionConfig{
			{
				Settings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	proxyUserID := protocol.NewID(uuid.New())
	proxyPort := pickPort()
	proxyConfig := &core.Config{
		Inbound: []*core.InboundConnectionConfig{
			{
				PortRange: v2net.SinglePortRange(proxyPort),
				ListenOn:  v2net.NewIPOrDomain(v2net.LocalHostIP),
				Settings: serial.ToTypedMessage(&inbound.Config{
					User: []*protocol.User{
						{
							Account: serial.ToTypedMessage(&vmess.Account{
								Id: proxyUserID.String(),
							}),
						},
					},
				}),
			},
		},
		Outbound: []*core.OutboundConnectionConfig{
			{
				Settings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	clientPort := pickPort()
	clientConfig := &core.Config{
		Inbound: []*core.InboundConnectionConfig{
			{
				PortRange: v2net.SinglePortRange(clientPort),
				ListenOn:  v2net.NewIPOrDomain(v2net.LocalHostIP),
				Settings: serial.ToTypedMessage(&dokodemo.Config{
					Address: v2net.NewIPOrDomain(dest.Address),
					Port:    uint32(dest.Port),
					NetworkList: &v2net.NetworkList{
						Network: []v2net.Network{v2net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*core.OutboundConnectionConfig{
			{
				Settings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: v2net.NewIPOrDomain(v2net.LocalHostIP),
							Port:    uint32(serverPort),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: serverUserID.String(),
									}),
								},
							},
						},
					},
				}),
				ProxySettings: &internet.ProxyConfig{
					Tag: "proxy",
				},
			},
			{
				Tag: "proxy",
				Settings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: v2net.NewIPOrDomain(v2net.LocalHostIP),
							Port:    uint32(proxyPort),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: proxyUserID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	assert.Error(InitializeServerConfig(serverConfig)).IsNil()
	assert.Error(InitializeServerConfig(proxyConfig)).IsNil()
	assert.Error(InitializeServerConfig(clientConfig)).IsNil()

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: int(clientPort),
	})
	assert.Error(err).IsNil()

	payload := "dokodemo request."
	nBytes, err := conn.Write([]byte(payload))
	assert.Error(err).IsNil()
	assert.Int(nBytes).Equals(len(payload))

	response := make([]byte, 1024)
	nBytes, err = conn.Read(response)
	assert.Error(err).IsNil()
	assert.Bytes(response[:nBytes]).Equals(xor([]byte(payload)))
	assert.Error(conn.Close()).IsNil()

	CloseAllServers()
}

func TestBlackhole(t *testing.T) {
	assert := assert.On(t)

	tcpServer := tcp.Server{
		MsgProcessor: xor,
	}
	dest, err := tcpServer.Start()
	assert.Error(err).IsNil()
	defer tcpServer.Close()

	tcpServer2 := tcp.Server{
		MsgProcessor: xor,
	}
	dest2, err := tcpServer2.Start()
	assert.Error(err).IsNil()
	defer tcpServer2.Close()

	serverPort := pickPort()
	serverPort2 := pickPort()
	serverConfig := &core.Config{
		Inbound: []*core.InboundConnectionConfig{
			{
				PortRange: v2net.SinglePortRange(serverPort),
				ListenOn:  v2net.NewIPOrDomain(v2net.LocalHostIP),
				Settings: serial.ToTypedMessage(&dokodemo.Config{
					Address: v2net.NewIPOrDomain(dest.Address),
					Port:    uint32(dest.Port),
					NetworkList: &v2net.NetworkList{
						Network: []v2net.Network{v2net.Network_TCP},
					},
				}),
			},
			{
				PortRange: v2net.SinglePortRange(serverPort2),
				ListenOn:  v2net.NewIPOrDomain(v2net.LocalHostIP),
				Settings: serial.ToTypedMessage(&dokodemo.Config{
					Address: v2net.NewIPOrDomain(dest2.Address),
					Port:    uint32(dest2.Port),
					NetworkList: &v2net.NetworkList{
						Network: []v2net.Network{v2net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*core.OutboundConnectionConfig{
			{
				Tag:      "direct",
				Settings: serial.ToTypedMessage(&freedom.Config{}),
			},
			{
				Tag:      "blocked",
				Settings: serial.ToTypedMessage(&blackhole.Config{}),
			},
		},
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&router.Config{
				Rule: []*router.RoutingRule{
					{
						Tag:       "blocked",
						PortRange: v2net.SinglePortRange(dest2.Port),
					},
				},
			}),
		},
	}

	assert.Error(InitializeServerConfig(serverConfig)).IsNil()

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: int(serverPort2),
	})
	assert.Error(err).IsNil()

	payload := "dokodemo request."
	{

		nBytes, err := conn.Write([]byte(payload))
		assert.Error(err).IsNil()
		assert.Int(nBytes).Equals(len(payload))
	}

	{
		response := make([]byte, 1024)
		_, err := conn.Read(response)
		assert.Error(err).IsNotNil()
	}

	assert.Error(conn.Close()).IsNil()

	CloseAllServers()
}
