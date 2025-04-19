package app

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"

	"pvz-service/internal/config"
	"pvz-service/internal/repository"
	"pvz-service/pkg/logger"

	"google.golang.org/grpc"
)

type App struct {
	cfg        config.Config
	log        logger.Logger
	HTTPServer *http.Server
	GRPCServer *grpc.Server
	Database   *repository.Storage
}

// NewApp создает новое приложение с заданной конфигурацией и логгером.
func NewApp(cfg config.Config, log logger.Logger) (*App, error) {
	db := repository.New(cfg.DB, log)

	return &App{
		cfg:      cfg,
		log:      log,
		Database: db, // Сохраняем подключение к базе данных
	}, nil
}

// MustRun инициализирует и запускает HTTP-сервер и gRPC-сервер.
func (a *App) MustRun() {
	// Инициализация HTTP-сервера
	a.HTTPServer = &http.Server{
		Addr:    a.cfg.Server.Address, // Используем адрес из конфигурации
		Handler: a.setupRoutes(),      // Настройка маршрутов
	}

	// Инициализация gRPC-сервера
	a.GRPCServer = grpc.NewServer()
	a.setupGRPCServices() // Настройка gRPC-сервисов
}

// Start запускает приложение.
func (a *App) Start(ctx context.Context) error {
	go func() {
		if err := a.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.Error("HTTP server failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	go func() {
		listener, err := net.Listen("tcp", a.cfg.GRPC.Address)
		if err != nil {
			a.log.Error("Failed to create server listener: ", slog.String("error", err.Error()))
			os.Exit(1)
		}
		if err := a.GRPCServer.Serve(listener); err != nil {
			a.log.Error("gRPC server failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	return nil
}

// Stop останавливает приложение.
func (a *App) Stop(ctx context.Context) error {
	a.log.Info("Stopping servers")
	if err := a.HTTPServer.Shutdown(ctx); err != nil {
		return err
	}
	a.GRPCServer.GracefulStop()
	if err := a.Database.Close(); err != nil {
		return err
	}
	return nil
}
