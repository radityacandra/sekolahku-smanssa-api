package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	PostgresUri        string `env:"POSTGRES_URI"`
	ServerPort         string `env:"PORT" env-default:"8080"`
	NewRelicLisenceKey string `env:"NR_LISENCE_KEY"`
}

func LoadGlobalConfig(logger *zap.Logger) (*Config, error) {
	var config Config
	if err := godotenv.Load(); err != nil {
		logger.Warn("no .env found, using default envvar...", zap.Error(err))
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
