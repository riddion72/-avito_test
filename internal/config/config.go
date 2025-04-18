package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB         DBConfig
	Server     ServerConfig
	GRPC       GRPCConfig
	Prometheus PrometheusConfig
}

type DBConfig struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     int    `envconfig:"DB_PORT" default:"5432"`
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

func NewConfig() *Config {
	// Загружаем переменные из .env файла (если существует)
	_ = godotenv.Load()

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return &cfg
}
