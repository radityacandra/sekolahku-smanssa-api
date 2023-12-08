package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUri string `env:"POSTGRES_URI"`
	ServerPort  string `env:"PORT" env-default:"8080"`
}

func LoadGlobalConfig() (*Config, error) {
	var config Config
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		return nil, err
	}

	return &config, nil
}