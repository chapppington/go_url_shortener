package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"urlshortener/internal/api"
	"urlshortener/config"
	"urlshortener/internal/infrastructure"
	"urlshortener/internal/logic"
)

func main() {
	cfg := config.LoadFromEnv()

	db, err := infrastructure.NewGormDB(cfg)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(&infrastructure.ShortURLModel{}); err != nil {
		slog.Error("failed to migrate database", "error", err)
		os.Exit(1)
	}
	
	repo := infrastructure.NewRepository(db)
	service := logic.NewService(repo)
	router := api.NewRouter(service)

	slog.Info("starting server", "port", cfg.APIPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.APIPort), router); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
