package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	appHttp "github.com/nayeem-bd/Todo-App/http"
	"github.com/nayeem-bd/Todo-App/internal/config"
	"github.com/nayeem-bd/Todo-App/internal/logger"
	loggerMiddleware "github.com/nayeem-bd/Todo-App/internal/middleware"
	"github.com/nayeem-bd/Todo-App/internal/migrations"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal("Failed to load config:", err)
	}

	db, err := config.ConnectDatabase(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database:", err)
	}

	// migrations
	migrations.Migrate(db)

	cache := config.ConnectRedis(cfg.Redis)

	defer cache.Close()

	addr := fmt.Sprintf(":%s", cfg.Server.Port)

	r := chi.NewRouter()

	//middlewares
	r.Use(middleware.RequestID)
	r.Use(loggerMiddleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))
	r.Handle("/metrics", promhttp.Handler())

	handler := appHttp.RegisterHandlers(db, cache)
	appHttp.SetupRouter(r, handler)

	srv := &http.Server{Addr: addr, Handler: r, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 120 * time.Second}

	go shutdownServer(srv)

	//fmt.Printf("Server is running on http://localhost%s\n", addr)
	logger.Info("Starting server on", addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("Failed to start server:", err)
	}
}

func shutdownServer(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Error shutting down server:", err)
	}
	logger.Info("Server gracefully stopped")
}
