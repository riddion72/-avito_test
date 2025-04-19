package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"pvz-service/internal/app"
	"pvz-service/internal/config"
	"pvz-service/pkg/logger"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	logg := logger.MustInit(cfg.LogLevel.Level)
	logg.Info("Starting application", slog.String("version", "1.0.0"))

	app, err := app.NewApp(*cfg, logg)
	if err != nil {
		logg.Error("failed to create app: %v", slog.String("error", err.Error()))
		os.Exit(1)
	}

	app.Server.MustRun()

	// Запуск приложения
	if err := app.Start(context.Background()); err != nil {
		logg.Error("failed to start app: %v", slog.String("error", err.Error()))
		os.Exit(1)
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

// // startHTTPServer запускает HTTP сервер
// func startHTTPServer(e *echo.Echo, logger logger.Logger) {
// 	// Настройка маршрутов
// 	// e.POST("/pvz", pvzHandler.CreatePVZ)
// 	// e.POST("/acceptance", acceptanceHandler.CreateAcceptance)
// 	// e.POST("/product", productHandler.AddProduct)
// 	// и т.д.

// 	go func() {
// 		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
// 			logger.Warn(fmt.Sprintf("HTTP server failed: %v", err))
// 		}
// 	}()
// }

// // startGRPCServer запускает gRPC сервер
// func startGRPCServer() {
// 	// Настройка gRPC сервера
// 	// gRPC-сервер будет запущен на порту 3000
// }
