package config

import (
	"tree/logger"

	"github.com/kelseyhightower/envconfig"
)

type PostgreConfig struct {
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	PostgresDataBase string `envconfig:"POSTGRES_DATABASE"`
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresPort     string `envconfig:"POSTGRES_PORT"`
}

func IniatilizePostgreConfig() *PostgreConfig {
	var p PostgreConfig
	err := envconfig.Process("", &p)
	logger.Error(err)
	return &p
}
