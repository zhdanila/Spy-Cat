package server

import (
	"context"
	"net/http"
	"sca/internal/config"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    time.Duration(cfg.HTTP.ReadTimeoutSeconds) * time.Second,
			WriteTimeout:   time.Duration(cfg.HTTP.WriteTimeoutSeconds) * time.Second,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderBytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
