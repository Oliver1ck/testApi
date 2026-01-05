package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_SSLMODE  string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("ошибка загрузки .env файла %w", err)
	}

	cfg := &Config{
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("ошибка валидации конфига %w", err)
	}
	return cfg, nil
}

func (config *Config) Validate() error {

	switch {
	case config.DB_NAME == "":
		return fmt.Errorf("DB_NAME is empty")
	case config.DB_USERNAME == "":
		return fmt.Errorf("DB_USERNAME is empty")
	case config.DB_PASSWORD == "":
		return fmt.Errorf("DB_PASSWORD is empty")
	case config.DB_HOST == "":
		return fmt.Errorf("DB_HOST is empty")
	case config.DB_PORT == "":
		return fmt.Errorf("DB_PORT is empty")
	case config.DB_SSLMODE == "":
		return fmt.Errorf("DB_SSLMODE is empty")
	default:
		return nil
	}
}
