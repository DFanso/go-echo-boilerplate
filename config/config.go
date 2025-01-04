package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
	}
	MongoDB struct {
		URI string
	}
}

func Load() *Config {
	// Load .env file
	godotenv.Load()

	cfg := &Config{}

	// Server configuration
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")

	// MongoDB configuration
	cfg.MongoDB.URI = getEnv("MONGODB_URI", "mongodb://localhost:27017")

	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
