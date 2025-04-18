package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"pvz-service/internal/config"
	"pvz-service/internal/handlers"
	"pvz-service/internal/metrics"
	"pvz-service/internal/repositories"
	"pvz-service/internal/services"
	"pvz-service/pkg/db"
	"pvz-service/pkg/logger"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,                     // Функция для создания конфигурации
			db.NewPostgresDB,                     // Функция для создания подключения к БД
			logger.New,                           // Функция для создания логгера
			handlers.NewPVZHandler,               // Функция для создания обработчика ПВЗ
			handlers.NewAcceptanceHandler,        // Функция для создания обработчика приёмок
			handlers.NewProductHandler,           // Функция для создания обработчика товаров
			repositories.NewPVZRepository,        // Функция для создания репозитория ПВЗ
			repositories.NewAcceptanceRepository, // Функция для создания репозитория приёмок
			repositories.NewProductRepository,    // Функция для создания репозитория товаров
			services.NewPVZService,               // Функция для создания сервиса ПВЗ
			services.NewAcceptanceService,        // Функция для создания сервиса приёмок
			services.NewProductService,           // Функция для создания сервиса товаров
			metrics.NewMetrics,                   // Функция для настройки метрик
		),
		fx.Invoke(
			startHTTPServer, // Функция для запуска HTTP сервера
			startGRPCServer, // Функция для запуска gRPC сервера
		),
	)

	// Запуск приложения
	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("failed to start app: %v", err)
	}

	// Обработка сигналов завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Остановка приложения
	if err := app.Stop(context.Background()); err != nil {
		log.Fatalf("failed to stop app: %v", err)
	}
}

// startHTTPServer запускает HTTP сервер
func startHTTPServer(e *echo.Echo, logger logger.Logger) {
	// Настройка маршрутов
	// e.POST("/pvz", pvzHandler.CreatePVZ)
	// e.POST("/acceptance", acceptanceHandler.CreateAcceptance)
	// e.POST("/product", productHandler.AddProduct)
	// и т.д.

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			logger.Warn(fmt.Sprintf("HTTP server failed: %v", err))
		}
	}()
}

// startGRPCServer запускает gRPC сервер
func startGRPCServer() {
	// Настройка gRPC сервера
	// gRPC-сервер будет запущен на порту 3000
}
