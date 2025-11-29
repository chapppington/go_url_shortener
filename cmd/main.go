package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"urlshortener/internal/api"
	"urlshortener/internal/config"
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

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.APIPort),
		Handler: router,
	}

	go func() {
		slog.Info("starting server", "port", cfg.APIPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server forced to shutdown", "error", err)
	}

	slog.Info("server exited")
}
