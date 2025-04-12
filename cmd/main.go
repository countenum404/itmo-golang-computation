package main

import (
	"countenum404/itmo-golang-computation/internal/app"
	"countenum404/itmo-golang-computation/internal/handlers"
	"countenum404/itmo-golang-computation/internal/service"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	config := &app.Config{
		Message: "APP_MESSAGE",
	}

	fx.New(
		fx.Provide(func() *app.Config {
			return config
		}),
		fx.Provide(func() *http.Server {
			return &http.Server{}
		}),
		service.Module,
		handlers.Module,
		app.Module,
	).Run()
}
