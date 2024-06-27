package config

import (
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret   string `envconfig:"JWT_SECRET"`
	EnvMode     string `envconfig:"ENVIRONMENT_MODE" default:"development"`
	DatabaseUrl string `envconfig:"DATABASE_URL"`
	RedisUrl    string `envconfig:"REDIS_URL"`

	EmailFrom     string `envconfig:"EMAIL_FROM"`
	EmailPassword string `envconfig:"EMAIL_PASSWORD"`
	EmailHost     string `envconfig:"EMAIL_HOST"`
	EmailPort     string `envconfig:"EMAIL_PORT"`
}

func NewWithEnvPath(envPath string) (*Config, error) {
	var cfg Config
	_ = godotenv.Load(envPath)

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func New() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	wd = filepath.Join(wd, "..", "..")

	envPath := filepath.Join(wd, ".env")

	return NewWithEnvPath(envPath)
}

func NewFx() Config {
	config, err := New()
	if err != nil {
		panic(err)
	}

	return *config
}
