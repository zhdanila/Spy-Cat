package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	HTTPPort string `mapstructure:"HTTP_PORT" validate:"required"`

	// Database configuration
	DBHost     string `mapstructure:"DB_HOST" validate:"required"`
	DBPort     string `mapstructure:"DB_PORT" validate:"required"`
	DBUsername string `mapstructure:"DB_USERNAME" validate:"required"`
	DBName     string `mapstructure:"DB_NAME" validate:"required"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE" validate:"required"`
	DBPassword string `mapstructure:"DB_PASSWORD" validate:"required"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		zap.L().Warn("Error reading .env file, using environment variables", zap.Error(err))
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	zap.L().Info("Config successfully loaded")
	return &cfg, nil
}
