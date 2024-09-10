package server

import (
	"context"
	"github.com/fanfaronDo/test_avito/internal/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.HttpServer, httpHandler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           cfg.Address,
		Handler:        httpHandler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
