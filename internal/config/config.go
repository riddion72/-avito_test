package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB         DBConfig
	Server     ServerConfig
	GRPC       GRPCConfig
	Prometheus PrometheusConfig
	LogLevel   Logger
}

type DBConfig struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

type ServerConfig struct {
	Address         string `envconfig:"HTTP_ADDRESS" default:":8080"`
	ShutdownTimeout int    `envconfig:"SHUTDOWN_TIMEOUT" default:"10"`
}

type GRPCConfig struct {
	Address string `envconfig:"GRPC_ADDRESS" default:":3000"`
}

type PrometheusConfig struct {
	Address string `envconfig:"PROMETHEUS_ADDRESS" default:":9000"`
}

type Logger struct {
	Level string `envconfig:"LEVEL" default:":local"`
}

func NewConfig() (*Config, error) {
	// Загружаем переменные из .env файла (если существует)
	_ = godotenv.Load()

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
