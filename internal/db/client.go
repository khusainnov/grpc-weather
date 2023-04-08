package db

import (
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/grpc-weather/internal/config"
	_ "github.com/lib/pq"
	pg "gitlab.com/khusainnov/driver/postgres"
	"go.uber.org/zap"
)

func NewClient(cfg *config.Config) (*ClientImpl, error) {
	db, err := pg.NewPostgresDB(
		pg.ConfigPG{
			Host:         cfg.PgHost,
			Port:         cfg.PgPort,
			User:         cfg.PgUser,
			Password:     cfg.PgPassword,
			DBName:       cfg.PgName,
			SSLMode:      cfg.PgSSLMode,
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

	pgURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.PgUser, cfg.PgPassword, cfg.PgHost, cfg.PgPort, cfg.PgName, cfg.PgSSLMode)
	if err = runDBMigration(cfg.MigrationPath, pgURL); err != nil {
		return nil, err
	}

	return &ClientImpl{db: db}, nil
}

type ClientImpl struct {
	db *sqlx.DB
}

func (db *ClientImpl) GetDB() *sqlx.DB {
	return db.db
}

func runDBMigration(sourceURL, databaseURL string) error {
	migration, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		return fmt.Errorf("cannot create migration, %w", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("cannot up migration, %w", err)
	}

	return nil
}
