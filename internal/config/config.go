package config

import (
	"time"

	"go.uber.org/zap"
)

type AppMode string

const (
	LocalAppMode AppMode = "local"
)

type Config struct {
	L              *zap.Logger
	AppMode        AppMode       `env:"APP_MODE" envDefault:"local"`
	GRPCAddr       string        `env:"GRPC_ADDR" envDefault:":9000"`
	HTTPAddr       string        `env:"HTTP_ADDR" envDefault:":8082"`
	PgPort         string        `env:"PG_PORT" envDefault:"5434"`
	PgHost         string        `env:"PG_HOST" envDefault:"localhost"`
	PgName         string        `env:"PG_NAME" envDefault:"db"`
	PgUser         string        `env:"PG_USER" envDefault:"db"`
	PgPassword     string        `env:"PG_PASSWORD" envDefault:"qwerty"`
	PgPingEnabled  bool          `env:"PG_PING_ENABLED" envDefault:"true"`
	PgPingInterval time.Duration `env:"PG_PING_INTERVAL" envDefault:"40m"`
	PgMaxOpenConn  int           `env:"PG_MAX_OPEN_CONN" envDefault:"10"`
	PgIdleConn     int           `env:"PG_MAX_IDLE_CONN" envDefault:"10"`

	// Weather API Token
	WeatherToken string `env:"WEATHER_API_TOKEN" envDefault:"c46b15514212497bb9a162423220904"`
}
