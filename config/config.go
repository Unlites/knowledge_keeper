package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Port            string        `env:"HTTP_SERVER_PORT" env-default:":8080"`
	ReadTimeout     time.Duration `env:"HTTP_SERVER_READ_TIMEOUT" env-default:"10s"`
	WriteTimeout    time.Duration `env:"HTTP_SERVER_WRITE_TIMEOUT" env-default:"10s"`
	MaxHeaderBytes  int           `env:"HTTP_SERVER_WRITE_MAX_HEADER_BYTES" env-default:"1"`
	ShutdownTimeout time.Duration `env:"HTTP_SERVER_SHUTDOWN_TIMEOUT" env-default:"5s"`
}

type Postgres struct {
	User     string `env:"DB_USERNAME" env-default:"root"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	DbName   string `env:"DB_DATABASE_NAME" env-required:"true"`
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     string `env:"DB_PORT" env-required:"true"`
	SslMode  string `env:"DB_SSL_MODE" env-default:"disable"`
}

type Config struct {
	HttpServer HttpServer
	Postgres   Postgres
}

func Init() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
