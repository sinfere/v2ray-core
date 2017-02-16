package impl

import (
	"context"

	"v2ray.com/core/app"
	"v2ray.com/core/app/dispatcher"
	"v2ray.com/core/app/log"
	"v2ray.com/core/app/proxyman"
	"v2ray.com/core/app/router"
	"v2ray.com/core/common"
	"v2ray.com/core/common/errors"
	"v2ray.com/core/common/net"
	"v2ray.com/core/proxy"
	"v2ray.com/core/transport/ray"
)

type DefaultDispatcher struct {
	ohm    proxyman.OutboundHandlerManager
	router *router.Router
}

func NewDefaultDispatcher(ctx context.Context, config *dispatcher.Config) (*DefaultDispatcher, error) {
	space := app.SpaceFromContext(ctx)
	if space == nil {
		return nil, errors.New("DefaultDispatcher: No space in context.")
	}
	d := &DefaultDispatcher{}
	space.OnInitialize(func() error {
		d.ohm = proxyman.OutboundHandlerManagerFromSpace(space)
		if d.ohm == nil {
			return errors.New("DefaultDispatcher: OutboundHandlerManager is not found in the space.")
		}
		d.router = router.FromSpace(space)
		return nil
	})
	return d, nil
}

func (DefaultDispatcher) Start() error {
	return nil
}

func (DefaultDispatcher) Close() {}

func (DefaultDispatcher) Interface() interface{} {
	return (*dispatcher.Interface)(nil)
}

func (v *DefaultDispatcher) Dispatch(ctx context.Context, destination net.Destination) (ray.InboundRay, error) {
	dispatcher := v.ohm.GetDefaultHandler()
	if !destination.IsValid() {
		panic("Dispatcher: Invalid destination.")
	}

	ctx = proxy.ContextWithTarget(ctx, destination)

	if v.router != nil {
		if tag, err := v.router.TakeDetour(ctx); err == nil {
			if handler := v.ohm.GetHandler(tag); handler != nil {
				log.Info("DefaultDispatcher: Taking detour [", tag, "] for [", destination, "].")
				dispatcher = handler
			} else {
				log.Warning("DefaultDispatcher: Nonexisting tag: ", tag)
			}
		} else {
			log.Info("DefaultDispatcher: Default route for ", destination)
		}
	}

	direct := ray.NewRay(ctx)
	go dispatcher.Dispatch(ctx, direct)

	return direct, nil
}

func init() {
	common.Must(common.RegisterConfig((*dispatcher.Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return NewDefaultDispatcher(ctx, config.(*dispatcher.Config))
	}))
}
