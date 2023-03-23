package endpoint

import (
	"github.com/khusainnov/grpc-weather/internal/app/weatherservice/service"
	wapi "github.com/khusainnov/grpc-weather/pkg/weatherapi"
)

type Endpoint struct {
	wapi.UnimplementedWeatherServiceServer
}

func NewEndpoint(srv *service.Service) *Endpoint {
	return &Endpoint{}
}
