package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"sca/internal/config"
	"time"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cnf *config.Config) (*sqlx.DB, error) {
	dbConfig := DBConfig{
		Host:     cnf.DBHost,
		Port:     cnf.DBPort,
		Username: cnf.DBUsername,
		DBName:   cnf.DBName,
		SSLMode:  cnf.DBSSLMode,
		Password: cnf.DBPassword,
	}

	var db *sqlx.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.DBName, dbConfig.Password, dbConfig.SSLMode))
		if err == nil {
			err = db.Ping()
			if err == nil {
				zap.L().Info("DB connected")
				return db, nil
			}
		}

		zap.L().Error("Database not ready, retrying in 5 seconds...", zap.Error(err))
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("unable to connect to database after 10 attempts: %w", err)
}
