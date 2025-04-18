package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
	"go.uber.org/fx"
)

// Config содержит параметры для подключения к базе данных
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// NewPostgresDB создает новое подключение к базе данных PostgreSQL
func NewPostgresDB(lc fx.Lifecycle, cfg *Config) (*sql.DB, error) {
	// Формируем строку подключения
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	// Открываем подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Проверяем подключение к базе данных
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Добавляем обработчик для закрытия подключения при остановке приложения
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	log.Println("Connected to PostgreSQL database")
	return db, nil
}
