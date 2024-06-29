package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type (
	Config struct {
		HTTP     HTTPConfig
		Logger   LoggerConf
		Postgres PostgresConf
	}

	HTTPConfig struct {
		Host                string `env:"SERVICE_HOST" env-default:"127.0.0.1"`
		Port                string `env:"SERVICE_PORT" env-default:"8080"`
		ReadTimeoutSeconds  int    `env:"SERVICE_READ_TIMEOUT_SECONDS" env-default:"10"`
		WriteTimeoutSeconds int    `env:"SERVICE_WRITE_TIMEOUT_SECONDS" env-default:"10"`
		MaxHeaderBytes      int    `env:"SERVICE_MAX_HEADER_BYTES" env-default:"1"`
	}

	LoggerConf struct {
		LogLevel         string         `env:"LOG_LEVEL" env-default:"INFO"`
		Encoding         string         `env:"LOG_ENCODING" env-default:"json"`
		OutputPaths      EnvStringSlice `env:"LOG_OUTPUT_PATHS" env-default:"stdout"`
		ErrorOutputPaths EnvStringSlice `env:"LOG_ERROR_OUTPUT_PATHS" env-default:"stderr"`
	}

	EnvStringSlice []string

	PostgresConf struct {
		Host     string `env:"POSTGRES_HOST" env-default:"127.0.0.1"`
		Port     string `env:"POSTGRES_PORT" env-default:"8080"`
		Username string `env:"POSTGRES_USERNAME" env-default:"admin"`
		Password string `env:"POSTGRES_PASSWORD" env-default:"123"`
		DBName   string `env:"POSTGRES_DBNAME" env-default:"default"`
		SSLMode  string `env:"POSTGRES_SSL" env-default:"disable"`
	}
)

func New() *Config {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envPath := filepath.Join(cwd, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Println("Error loading .env file, using default environment variables")
	}

	var c Config
	err = cleanenv.ReadEnv(&c)
	if err != nil {
		panic(err)
	}
	return &c
}
