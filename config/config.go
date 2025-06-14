package config

import (
	"data-enricher-dispatcher/apperrors"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment      string   `env:"ENVIRONMENT,required"`
	GetUsersURL      string   `env:"GET_USERS_URL,required"`
	PostUsersURL     string   `env:"POST_USERS_URL,required"`
	ExcludePostfixes []string `env:"EXCLUDE_POSTFIXES" envSeparator:","`
}

func NewConfig(envFile string) (*Config, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		return nil, apperrors.EnvConfigLoadError.AppendMessage(err)
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return cfg, apperrors.EnvConfigParseError.AppendMessage(err)
	}

	return cfg, nil
}
