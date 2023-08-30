package main

import (
	"go-fiber/src/env"
	"go-fiber/src/framework"
	"go.uber.org/fx"
)

func main() {
	// Create an instance of the fx application.
	app := fx.New(
		fx.Provide(
			func() (*env.Env, error) {
				return env.LoadEnv()
			},
		),
		framework.Module,
	)
	app.Run()
}
