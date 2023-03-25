package service

import (
	"github.com/khusainnov/grpc-weather/internal/config"
	"github.com/khusainnov/grpc-weather/internal/repository"
	wapi "github.com/khusainnov/grpc-weather/pkg/weatherapi"
)

type Weather interface {
	GetWeather(cfg *config.Config, req *wapi.WeatherRequest) (*wapi.WeatherResponse, error)
}

type Service struct {
	Weather
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
