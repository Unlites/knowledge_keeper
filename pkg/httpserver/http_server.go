package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Settings struct {
	Port           string
	Handler        http.Handler
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

type HttpServer struct {
	server *http.Server
}

func NewHttpServer(settings *Settings) *HttpServer {
	return &HttpServer{
		server: &http.Server{
			Addr:           ":" + settings.Port,
			Handler:        settings.Handler,
			ReadTimeout:    settings.ReadTimeout,
			WriteTimeout:   settings.WriteTimeout,
			MaxHeaderBytes: settings.MaxHeaderBytes,
		},
	}
}

func (s *HttpServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *HttpServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
