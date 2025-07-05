package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	appHttp "github.com/nayeem-bd/Todo-App/http"
	"github.com/nayeem-bd/Todo-App/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("Failed to load config:", err)
		os.Exit(1)
	}

	db, err := config.ConnectDatabase(cfg.Database)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}

	addr := fmt.Sprintf(":%s", cfg.Server.Port)

	r := chi.NewRouter()

	handler := appHttp.RegisterHandlers(db)
	appHttp.SetupRouter(r, handler)

	srv := &http.Server{Addr: addr, Handler: r, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 120 * time.Second}

	go shutdownServer(srv)

	fmt.Printf("Server is running on http://localhost%s\n", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}

func shutdownServer(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	fmt.Println("\nShutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Error shutting down server:", err)
	}
	fmt.Println("Server gracefully stopped")
}
