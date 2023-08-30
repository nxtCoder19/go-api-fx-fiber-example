package server

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type ServerOptions interface {
	GetPort() int
}

func NewFiberApp() *fiber.App {
	app := fiber.New()
	return app
}

// NewHelloHandler returns a handler for the /hello route.
func NewHelloHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}
}

func NewServerFx[T ServerOptions]() fx.Option {
	return fx.Module("server",
		fx.Provide(
			NewFiberApp,
			NewHelloHandler,
		),
		fx.Invoke(func(lifecycle fx.Lifecycle, app *fiber.App, helloHandler fiber.Handler, env T) {
			app.Get("/hello", helloHandler)
			lifecycle.Append(fx.Hook{
				OnStart: func(context.Context) error {
					go func() {
						// port := ":3000"
						port := env.GetPort()
						log.Printf("Starting server on port :%d", port)
						if err := app.Listen(":" + strconv.Itoa(port)); err != nil {
							log.Fatalf("Failed to start server: %v", err)
						}
					}()
					return nil
				},
				OnStop: func(context.Context) error {
					log.Println("Stopping server gracefully")
					return app.Shutdown()
				},
			})
		}),
	)
}
