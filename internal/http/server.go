package http

import (
	"net/http"

	"github.com/khusainnov/grpc-weather/internal/app/weatherservice/endpoint"
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

func (s *Server) Start(e *endpoint.Endpoint) {
	s.srv.Handler = s.setupRoutes(e)

	go func() {
		s.cfg.L.Info("start listening http", zap.String("PORT", s.cfg.HTTPAddr))
		if err := http.ListenAndServe(s.cfg.HTTPAddr, s.srv.Handler); err != nil {
			s.cfg.L.Fatal("error run http & gateway server", zap.Error(err))
		}
	}()
}
