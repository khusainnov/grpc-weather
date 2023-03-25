package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/grpc-weather/internal/config"
	_ "github.com/lib/pq"
	"gitlab.com/khusainnov/driver/postgres"
	"go.uber.org/zap"
)

func NewClient(cfg *config.Config) (*ClientImpl, error) {
	db, err := postgres.NewPostgresDB(
		postgres.ConfigPG{
			Host:         cfg.PgHost,
			Port:         cfg.PgPort,
			User:         cfg.PgUser,
			Password:     cfg.PgPassword,
			DBName:       cfg.PgName,
			SSLMode:      "disable",
			MaxOpenConns: cfg.PgMaxOpenConn,
			MaxIdleConns: cfg.PgIdleConn,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("cannot init db, %w", err)
	}

	go func() {
		t := time.NewTicker(cfg.PgPingInterval)

		for range t.C {
			if err := db.Ping(); err != nil {
				cfg.L.Error("error ping", zap.Error(err))
			}
		}
	}()

	return &ClientImpl{db: db}, nil
}

type ClientImpl struct {
	db *sqlx.DB
}

func (db *ClientImpl) GetDB() *sqlx.DB {
	return db.db
}
