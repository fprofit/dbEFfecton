package entry

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Host      string `env:"DB_HOST"`
	User      string `env:"DB_USER"`
	Password  string `env:"DB_PASSWORD"`
	DBName    string `env:"DB_NAME"`
	Port      string `env:"DB_PORT"`
	DBSslmode string `env:"DB_SSLMODE"`
}

func InitializeEnv() (config EnvConfig, err error) {
	if err := godotenv.Load(); err != nil {
		return EnvConfig{}, fmt.Errorf("Failed to load .env file: %w", err)
	}

	if err := env.Parse(&config); err != nil {
		return EnvConfig{}, fmt.Errorf("Failed to parse env from environment variables: %w", err)
	}
	if err := validateEnvConfig(config); err != nil {
		return EnvConfig{}, err
	}

	return config, nil
}

func validateEnvConfig(config EnvConfig) error {
	if config.Host == "" {
		return fmt.Errorf("DB_HOST is a required environment variable")
	}

	if config.User == "" {
		return fmt.Errorf("DB_USER is a required environment variable")
	}
	if config.Password == "" {
		return fmt.Errorf("DB_PASSWORD is a required environment variable")
	}
	if config.Port == "" {
		return fmt.Errorf("DB_PORT is a required environment variable")
	}

	if config.DBName == "" {
		return fmt.Errorf("DB_NAME is a required environment variable")
	}

	if config.DBSslmode == "" {
		return fmt.Errorf("DB_SSLMODE is a required environment variable")
	}
	return nil
}
