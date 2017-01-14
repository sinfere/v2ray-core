package core

import (
	"context"

	"v2ray.com/core/app"
	"v2ray.com/core/common/dice"
	"v2ray.com/core/common/log"
	"v2ray.com/core/common/retry"
	"v2ray.com/core/proxy"
)

// InboundDetourHandlerAlways is a handler for inbound detour connections.
type InboundDetourHandlerAlways struct {
	space  app.Space
	config *InboundConnectionConfig
	ich    []proxy.InboundHandler
}

func NewInboundDetourHandlerAlways(ctx context.Context, config *InboundConnectionConfig) (*InboundDetourHandlerAlways, error) {
	space := app.SpaceFromContext(ctx)
	handler := &InboundDetourHandlerAlways{
		space:  space,
		config: config,
	}
	ports := config.PortRange
	handler.ich = make([]proxy.InboundHandler, 0, ports.To-ports.From+1)
	for i := ports.FromPort(); i <= ports.ToPort(); i++ {
		ichConfig, err := config.GetTypedSettings()
		if err != nil {
			return nil, err
		}
		ich, err := proxy.CreateInboundHandler(proxy.ContextWithInboundMeta(ctx, &proxy.InboundHandlerMeta{
			Address:                config.GetListenOnValue(),
			Port:                   i,
			Tag:                    config.Tag,
			StreamSettings:         config.StreamSettings,
			AllowPassiveConnection: config.AllowPassiveConnection,
		}), ichConfig)

		if err != nil {
			log.Error("Failed to create inbound connection handler: ", err)
			return nil, err
		}
		handler.ich = append(handler.ich, ich)
	}
	return handler, nil
}

func (v *InboundDetourHandlerAlways) GetConnectionHandler() (proxy.InboundHandler, int) {
	ich := v.ich[dice.Roll(len(v.ich))]
	return ich, int(v.config.GetAllocationStrategyValue().GetRefreshValue())
}

func (v *InboundDetourHandlerAlways) Close() {
	for _, ich := range v.ich {
		ich.Close()
	}
}

// Start starts the inbound connection handler.
func (v *InboundDetourHandlerAlways) Start() error {
	for _, ich := range v.ich {
		err := retry.ExponentialBackoff(10 /* times */, 200 /* ms */).On(func() error {
			err := ich.Start()
			if err != nil {
				log.Error("Failed to start inbound detour:", err)
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
