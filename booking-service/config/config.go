package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func LoadConfig() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "localhost"
	}

	port := os.Getenv("Port")
	if port == "" {
		port = "8080"
	}

	return &Config{
		DatabaseURL: dbURL,
		Port:        port,
	}
}
