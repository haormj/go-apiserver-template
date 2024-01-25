package provider

import (
	"context"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/haormj/go-apiserver-template/internal/option"
	"github.com/haormj/go-apiserver-template/internal/service"
	"github.com/haormj/go-apiserver-template/pkg/jsonrpc"
)

type HTTP struct {
	quit            chan struct{}
	opt             *option.HTTPProvider
	svc             service.Service
	userHandlerFunc http.HandlerFunc
}

func NewHTTP(opt *option.HTTPProvider, svc service.Service) (*HTTP, error) {
	userHandlerFunc, err := jsonrpc.NewHandleFunc(context.Background(), svc.User())
	if err != nil {
		return nil, err
	}

	return &HTTP{
		quit:            make(chan struct{}),
		opt:             opt,
		svc:             svc,
		userHandlerFunc: userHandlerFunc,
	}, nil
}

func (h *HTTP) Run(ctx context.Context) error {
	server := http.Server{
		Addr: h.opt.Addr,
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/user", h.userHandlerFunc)
		server.Handler = mux
		return server.ListenAndServe()
	})

	g.Go(func() error {
		<-ctx.Done()
		return server.Shutdown(ctx)
	})

	return g.Wait()
}
