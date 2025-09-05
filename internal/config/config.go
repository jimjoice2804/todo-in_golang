package config

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTP struct {
	Address         string        `yaml:"http_address" env:"HTTP_ADDRESS" env-default:"localhost:8080"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout" env:"HTTP_SHUTDOWN_TIMEOUT" env-default:"10s"`
}

type CORS struct {
	AllowedOrigins []string `yaml:"allowed_origins" env:"CORS_ALLOWED_ORIGINS" env-separator:","`
}

type Config struct {
	Env  string `yaml:"env" env:"APP_ENV" env-default:"development"`
	HTTP HTTP   `yaml:"HTTP"`
	CORS CORS   `yaml:"CORS"`
}

func MustLoad() *Config {
	var cfg Config

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "local.yaml"
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config from %s: %v", configPath, err)
	}

	if len(cfg.CORS.AllowedOrigins) == 0 {
		log.Fatalf("CORS.AllowedOrigins cannot be empty")
	}

	if cfg.HTTP.ShutdownTimeout <= 0 {
		log.Fatalf("TimeOut must be > 0 %v", cfg.HTTP.ShutdownTimeout)
	}

	if _, err := net.ResolveTCPAddr("tcp", cfg.HTTP.Address); err != nil {
		log.Fatalf("Failed to resolve HTTP address %q: %v", cfg.HTTP.Address, err)
	}

	return &cfg
}
