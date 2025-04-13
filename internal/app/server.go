package app

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("server",
	fx.Provide(zap.NewExample),
	fx.Invoke(NewGrpcServer),
)

type Config struct {
	Message string
}

type Server interface {
	Start() error
}
