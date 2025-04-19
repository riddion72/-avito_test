package db

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"pvz-service/internal/config"
	"pvz-service/pkg/logger"
)

func New(cfg config.DBConfig, log logger.Logger) *sqlx.DB {
	const f = "storage.NewStorage"
	var db *sqlx.DB

	// Подключение к PostgreSQL с повторами
	var err error
	db, err = connectWithRetries(cfg, log)
	if err != nil {
		log.Error("Failed to connect to PostgreSQL: ", slog.String("func", f), slog.String("error", err.Error()))
		os.Exit(1)
	}

	return db
}

const (
	maxRetries = 5
	retryDelay = 5 * time.Second
)

func connectWithRetries(cfg config.DBConfig, log logger.Logger) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)
	log.Debug("preparing connect to PostgreSQL", slog.String("dsn", dsn))
	for i := 0; i < maxRetries; i++ {
		db, err = sqlx.Open("postgres", dsn)
		if err == nil {
			if err = db.Ping(); err == nil {
				log.Info("Successfully connected to PostgreSQL",
					slog.String("host", cfg.Host),
					slog.String("db", cfg.Name))
				return db, nil
			}
		}

		log.Warn("Failed to connect to PostgreSQL, retrying...",
			slog.Int("attempt", i+1),
			slog.String("error", err.Error()))

		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("failed to connect after %d attempts: %v", maxRetries, err)
}
