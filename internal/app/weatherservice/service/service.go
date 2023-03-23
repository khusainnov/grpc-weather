package service

import (
	"github.com/khusainnov/grpc-weather/internal/repository"
	wapi "github.com/khusainnov/grpc-weather/pkg/weatherapi"
)

type Weather interface {
	GetWeather(req *wapi.WeatherRequest) (*wapi.WeatherResponse, error)
}

type Service struct {
	Weather
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
