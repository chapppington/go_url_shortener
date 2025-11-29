package main

import (
	"fmt"
	"log/slog"
	"urlshortener/internal/config"
)

func main() {
	cfg := config.LoadFromEnv()

	fmt.Println(cfg)

	slog.Info("starting urlshortener", "env", cfg.Env)
}