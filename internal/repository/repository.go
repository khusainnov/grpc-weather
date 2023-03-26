package repository

import (
	"github.com/jmoiron/sqlx"
)

type Weather interface {
	UploadWeather(city string) error
}

type Repository struct {
	Weather
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}
