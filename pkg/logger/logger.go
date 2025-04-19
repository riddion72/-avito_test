package logger

import (
	"context"
	"log"
	"log/slog"
	"os"
)

// Logger определяет контракт для логгера
type Logger interface {
	Debug(msg string, attrs ...slog.Attr)
	Info(msg string, attrs ...slog.Attr)
	Warn(msg string, attrs ...slog.Attr)
	Error(msg string, attrs ...slog.Attr)
	With(attrs ...slog.Attr) Logger
	WithContext(ctx context.Context) Logger
}

// Реализация логгера
type slogLogger struct {
	logger *slog.Logger
}

const (
	LevelLocal = "local"
	LevelDev   = "dev"
	LevelProd  = "prod"
)

// MustInit создает новый экземпляр логгера или роняет сервер
func MustInit(level string) Logger {
	var handler slog.Handler

	switch level {
	case LevelLocal:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	case LevelDev:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	case LevelProd:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	default:
		log.Printf("unknown log level: %v", level)
	}

	return &slogLogger{
		logger: slog.New(handler),
	}
}

func (l *slogLogger) log(level slog.Level, msg string, attrs ...slog.Attr) {
	if !l.logger.Enabled(context.Background(), level) {
		return
	}

	args := make([]any, 0, len(attrs))
	for _, attr := range attrs {
		args = append(args, attr)
	}

	switch level {
	case slog.LevelDebug:
		l.logger.Debug(msg, args...)
	case slog.LevelInfo:
		l.logger.Info(msg, args...)
	case slog.LevelWarn:
		l.logger.Warn(msg, args...)
	case slog.LevelError:
		l.logger.Error(msg, args...)
	}
}

func (l *slogLogger) Debug(msg string, attrs ...slog.Attr) {
	l.log(slog.LevelDebug, msg, attrs...)
}

func (l *slogLogger) Info(msg string, attrs ...slog.Attr) {
	l.log(slog.LevelInfo, msg, attrs...)
}

func (l *slogLogger) Warn(msg string, attrs ...slog.Attr) {
	l.log(slog.LevelWarn, msg, attrs...)
}

func (l *slogLogger) Error(msg string, attrs ...slog.Attr) {
	l.log(slog.LevelError, msg, attrs...)
}

func (l *slogLogger) With(fields ...slog.Attr) Logger {
	args := make([]any, len(fields))
	for i, attr := range fields {
		args[i] = attr
	}
	return &slogLogger{
		logger: l.logger.With(args...),
	}
}

func (l *slogLogger) WithContext(ctx context.Context) Logger {
	if ctx == nil {
		return l
	}

	if reqID, ok := ctx.Value("request_id").(string); ok {
		return l.With(slog.String("request_id", reqID))
	}
	return l
}

// Helpers для удобной работы с ошибками
func Err(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func Attr(key string, value any) slog.Attr {
	return slog.Any(key, value)
}
