package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle) *http.Server {
	srv := &http.Server{Addr: ":8080"}
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		ln, err := net.Listen("tcp", srv.Addr)
		if err != nil {
			return err
		}
		fmt.Println("starting server at", srv.Addr)
		go srv.Serve(ln)
		return nil
	},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func main() {
	fx.New(
		fx.Provide(NewHttpServer),
		fx.Invoke(func(server *http.Server) {}),
	).Run()
}
