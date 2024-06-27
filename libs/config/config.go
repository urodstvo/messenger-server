package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	EnvMode     string `envconfig:"ENVIRONMENT_MODE" default:"development"`
	DatabaseUrl string `envconfig:"DATABASE_URL"`
}

func New() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func NewFx() Config {
	config, err := New()
	if err != nil {
		panic(err)
	}

	return *config
}
