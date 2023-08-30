package framework

import (
	"go-fiber/src/env"
	"go-fiber/src/server"
	"go.uber.org/fx"
)

type fm struct {
	*env.Env
}

func (f *fm) GetPort() int {
	return f.Port
}

var Module fx.Option = fx.Module(
	"framework",
	fx.Provide(func(ev *env.Env) *fm {
		return &fm{Env: ev}
	}),
	server.NewServerFx[*fm](),
)
