package cmd

import (
	"github.com/nayeem-bd/Todo-App/internal/config"
	"github.com/nayeem-bd/Todo-App/internal/logger"
	worker "github.com/nayeem-bd/Todo-App/internal/queue"
)

func Work() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal("Failed to load config:", err)
	}

	db, err := config.ConnectDatabase(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database:", err)
	}

	cache := config.ConnectRedis(cfg.Redis)

	queue, err := config.SetupRabbitMQConnection(cfg.RabbitMQ)

	if err != nil {
		logger.Fatal("Failed to connect to RabbitMQ:", err)
	}

	worker.Work(db, cache, queue)
}
