package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	StoragePath string
	HTTPServer  HTTPServer
}

type HTTPServer struct {
	Address     string
	Timeout     time.Duration
	IdleTimeout time.Duration
}

func LoadFromEnv() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found, using environment variables")
	}

	storagePath := getEnv("STORAGE_PATH", "")
	if storagePath == "" {
		log.Fatal("required environment variable STORAGE_PATH is not set")
	}

	config := &Config{
		Env:         getEnv("ENV", "development"),
		StoragePath: storagePath,
		HTTPServer: HTTPServer{
			Address:     getEnv("HTTP_ADDRESS", "localhost:8080"),
			Timeout:     parseDuration(getEnv("HTTP_TIMEOUT", "4s")),
			IdleTimeout: parseDuration(getEnv("HTTP_IDLE_TIMEOUT", "4s")),
		},
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseDuration(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		panic(fmt.Sprintf("invalid duration format: %s", s))
	}
	return duration
}