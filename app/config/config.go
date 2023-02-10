package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB DB
}

type DB struct {
	Driver   string `envconfig:"DB_DRIVER" default:"postgres"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"postgres"`
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	Name     string `envconfig:"DB_NAME" default:"develop"`
	Schema   string `envconfig:"DB_SCHEMA" default:"public"`
}

func New() (*Config, error) {
	var db DB
	if err := envconfig.Process("DB", &db); err != nil {
		return nil, fmt.Errorf("failed to get db env [%s]", err)
	}
	return &Config{DB: db}, nil
}
