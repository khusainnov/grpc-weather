package http

import (
	"net"
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
		lis, err := net.Listen("tcp", s.cfg.HTTPAddr)
		if err != nil {
			s.cfg.L.Fatal("error listen http & gateway server", zap.Any("PORT", s.cfg.HTTPAddr), zap.Error(err))
		}

		s.cfg.L.Info("start listening http", zap.String("PORT", s.cfg.HTTPAddr))
		if err := http.Serve(lis, s.srv.Handler); err != nil {
			s.cfg.L.Fatal("error run http & gateway server", zap.Error(err))
		}
	}()
}
