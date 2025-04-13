package app

import (
	"context"
	"countenum404/itmo-golang-computation/internal/handlers"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type HttpServer struct {
	Handlers   handlers.Handlers
	ServeMux   *http.ServeMux
	HttpServer *http.Server
	Logger     *zap.Logger
	Message    string
}

func (s *HttpServer) Start() error {
	return s.HttpServer.ListenAndServe()
}

func (s *HttpServer) RegisterHandlers() {
	s.ServeMux.HandleFunc("POST /{$}", s.Handlers.HandleCalc())

	s.ServeMux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Send a POST request to process your query"))
	})

	s.ServeMux.HandleFunc("/", http.NotFound)
}

func NewHttpServer(lc fx.Lifecycle, httpServer *http.Server, logger *zap.Logger, cfg *Config, handlers handlers.Handlers) *HttpServer {
	s := &HttpServer{HttpServer: httpServer, Logger: logger, Message: cfg.Message, Handlers: handlers, ServeMux: http.DefaultServeMux}
	s.HttpServer.Addr = ":8080"
	s.RegisterHandlers()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := s.Start()
				if err != nil {
					s.Logger.Fatal("Failed to start server", zap.Error(err))
				}
			}()
			s.Logger.Info("HTTP Server started")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.Logger.Info("HTTP Server stopped")
			return nil
		},
	})
	return s
}
