package main

import (
	"countenum404/itmo-golang-computation/internal/app"
	"countenum404/itmo-golang-computation/internal/handlers"
	"countenum404/itmo-golang-computation/internal/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(zap.NewExample),
		service.Module,
		handlers.Module,
		app.Module,
	).Run()
}
