package http

import (
	"net/http"

	"github.com/khusainnov/grpc-weather/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	cfg *config.Config
	srv *http.Server
}

func New(cfg *config.Config) *Server {
	srv := &http.Server{
		Addr: cfg.HTTPAddr,
	}

	return &Server{
		cfg: cfg,
		srv: srv,
	}
}

func (s *Server) Start() {
	s.srv.Handler = s.setupRoutes()

	go func() {
		s.cfg.L.Info("start listening http", zap.String("PORT", s.cfg.HTTPAddr))
		if err := s.srv.ListenAndServe(); err != nil {
			s.cfg.L.Fatal("error run http", zap.Error(err))
		}
	}()
}
