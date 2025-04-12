package app

import (
	"context"
	"countenum404/itmo-golang-computation/internal/handlers"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type HttpServer struct {
	Handlers   handlers.Handlers
	ServeMux   *http.ServeMux
	HttpServer *http.Server
	Logger     *zap.Logger
	Message    string
}

func (s *HttpServer) Start() {
	log.Fatal(s.HttpServer.ListenAndServe())
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
			go s.Start()
			s.Logger.Log(zap.InfoLevel, "Server started")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.Logger.Log(zap.InfoLevel, "Server stopped")
			return nil
		},
	})
	return s
}
