package app

import (
	"countenum404/itmo-golang-computation/internal/handlers"
	_ "countenum404/itmo-golang-computation/internal/handlers/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
)

type HttpServer struct {
	Handlers   handlers.Handlers
	ServeMux   *http.ServeMux
	HttpServer *http.Server
	Logger     *zap.Logger
}

func (s *HttpServer) Start() error {
	return s.HttpServer.ListenAndServe()
}

func (s *HttpServer) RegisterHandlers() {
	s.ServeMux.HandleFunc("POST /{$}", s.Handlers.HandleCalc())

	s.ServeMux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Send a POST request to process your query"))
	})

	s.ServeMux.HandleFunc("/swagger", httpSwagger.WrapHandler)

	s.ServeMux.HandleFunc("/", http.NotFound)
}

func NewHttpServer(httpServer *http.Server, logger *zap.Logger, handlers handlers.Handlers) *HttpServer {
	s := &HttpServer{HttpServer: httpServer, Logger: logger, Handlers: handlers, ServeMux: http.DefaultServeMux}
	s.HttpServer.Addr = ":8080"
	s.RegisterHandlers()
	return s
}
