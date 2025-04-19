package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"pvz-service/internal/config"
	"pvz-service/pkg/db"
	"pvz-service/pkg/logger"
)

type Storage struct {
	DB *sqlx.DB
}

func New(cfg config.DBConfig, log logger.Logger) *Storage {
	database := db.New(cfg, log)

	return &Storage{
		DB: database, // Сохраняем соединение с базой данных
	}
}

// Close закрывает соединение с базой данных.
func (s *Storage) Close() error {
	return s.DB.Close() // Закрываем соединение с базой данных
}
