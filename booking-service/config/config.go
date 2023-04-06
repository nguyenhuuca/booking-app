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
		//log.Fatal("DATABASE_URL environment variable is not set")
		dbURL = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		DatabaseURL: dbURL,
		Port:        port,
	}
}
