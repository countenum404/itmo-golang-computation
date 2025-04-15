package app

import (
	"context"
	"countenum404/itmo-golang-computation/internal/handlers"
	"countenum404/itmo-golang-computation/internal/service"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"os"
)

var Module = fx.Module("server",
	fx.Provide(NewConfig),
	fx.Provide(func(cfg *Config) *http.Server {
		if cfg.ServerType == ServerTypeHTTP {
			return &http.Server{}
		}
		return nil
	}),
	fx.Provide(func(cfg *Config) *grpc.Server {
		if cfg.ServerType == ServerTypeGRPC {
			return grpc.NewServer()
		}
		return nil
	}),
	fx.Provide(NewServer),
	fx.Invoke(StartServer),
)

const (
	ServerTypeHTTP = "http"
	ServerTypeGRPC = "grpc"
)

type Config struct {
	ServerType string
}

func NewConfig() *Config {
	return &Config{
		ServerType: os.Getenv("SERVER_TYPE"),
	}
}

type Server interface {
	Start() error
}

func NewServer(httpServer *http.Server, grpcServer *grpc.Server, logger *zap.Logger, cfg *Config, handlers handlers.Handlers, solverService service.SolverService) Server {
	if cfg.ServerType == "grpc" {
		return NewGrpcServer(grpcServer, logger, solverService)
	}
	return NewHttpServer(httpServer, logger, handlers)
}

func StartServer(lc fx.Lifecycle, server Server, cfg *Config, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := server.Start()
				if err != nil {
					logger.Fatal("Failed to start server", zap.Error(err))
				}
			}()
			logger.Info(fmt.Sprintf("%s Server started", cfg.ServerType))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info(fmt.Sprintf("%s Server stopped", cfg.ServerType))
			return nil
		},
	})
}
