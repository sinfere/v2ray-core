// Package proxy contains all proxies used by V2Ray.
package proxy

import (
	v2net "v2ray.com/core/common/net"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/transport/internet"
	"v2ray.com/core/transport/ray"
)

type HandlerState int

const (
	HandlerStateStopped = HandlerState(0)
	HandlerStateRunning = HandlerState(1)
)

type SessionInfo struct {
	Source      v2net.Destination
	Destination v2net.Destination
	User        *protocol.User
	Inbound     *InboundHandlerMeta
}

type InboundHandlerMeta struct {
	Tag                    string
	Address                v2net.Address
	Port                   v2net.Port
	AllowPassiveConnection bool
	StreamSettings         *internet.StreamConfig
}

type OutboundHandlerMeta struct {
	Tag            string
	Address        v2net.Address
	StreamSettings *internet.StreamConfig
	ProxySettings  *internet.ProxyConfig
}

func (v *OutboundHandlerMeta) GetDialerOptions() internet.DialerOptions {
	return internet.DialerOptions{
		Stream: v.StreamSettings,
		Proxy:  v.ProxySettings,
	}
}

// An InboundHandler handles inbound network connections to V2Ray.
type InboundHandler interface {
	// Listen starts a InboundHandler.
	Start() error
	// Close stops the handler to accepting anymore inbound connections.
	Close()
	// Port returns the port that the handler is listening on.
	Port() v2net.Port
}

// An OutboundHandler handles outbound network connection for V2Ray.
type OutboundHandler interface {
	// Dispatch sends one or more Packets to its destination.
	Dispatch(destination v2net.Destination, ray ray.OutboundRay)
}
