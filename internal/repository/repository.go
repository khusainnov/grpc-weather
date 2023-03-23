package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/grpc-weather/internal/model"
)

type Weather interface {
	UploadWeather(city string) (*model.Weather, error)
}

type Repository struct {
	Weather
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
