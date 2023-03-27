package endpoint

import (
	"github.com/khusainnov/grpc-weather/internal/app/weatherservice/service"
	"github.com/khusainnov/grpc-weather/internal/config"
	wapi "github.com/khusainnov/grpc-weather/pkg/weatherapi"
)

type Endpoint struct {
	wapi.UnimplementedWeatherServiceServer
	cfg *config.Config
	srv *service.Service
}

func NewEndpoint(srv *service.Service, cfg *config.Config) *Endpoint {
	return &Endpoint{
		cfg: cfg,
		srv: srv,
	}
}
