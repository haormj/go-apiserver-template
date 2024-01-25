package provider

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/haormj/go-apiserver-template/internal/option"
	"github.com/haormj/go-apiserver-template/internal/service"
)

type Provider struct {
	HTTP *HTTP
	svc  service.Service
}

func New(opt *option.Provider, svc service.Service) (*Provider, error) {
	httpProvider, err := NewHTTP(&opt.HTTP, svc)
	if err != nil {
		return nil, err
	}

	p := &Provider{
		HTTP: httpProvider,
		svc:  svc,
	}

	return p, nil
}

func (p *Provider) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return p.HTTP.Run(ctx)
	})

	return g.Wait()
}
