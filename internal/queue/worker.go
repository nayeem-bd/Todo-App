package queue

import (
	"context"
	"github.com/nayeem-bd/Todo-App/internal/config"
	"github.com/nayeem-bd/Todo-App/internal/logger"
	"github.com/nayeem-bd/Todo-App/internal/store"
	queue2 "github.com/nayeem-bd/Todo-App/modules/todo/delivery/queue"
	"github.com/nayeem-bd/Todo-App/modules/todo/usecase"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

func Work(db *gorm.DB, cache *config.Cache, queue *config.Queue) {
	todoUsecase := usecase.NewTodoUsecase(store.New(db), cache, queue)
	todoWorker := queue2.NewTodoWorker(todoUsecase)

	ch, err := queue.Conn.Channel()
	if err != nil {
		panic("Failed to open a channel: " + err.Error())
	}

	defer ch.Close()

	msgs, err := ch.Consume(queue.QueueName, "", false, false, false, false, nil)
	if err != nil {
		panic("Failed to register a consumer: " + err.Error())
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		for msg := range msgs {
			if err := todoWorker.ProcessMessage(ctx, msg); err != nil {
				logger.Error("Failed to process message: " + err.Error())
				if err := msg.Nack(false, true); err != nil {
					logger.Error("Failed to nack message: " + err.Error())
				}
				continue
			}
			if err := msg.Ack(false); err != nil {
				logger.Error("Failed to ack message: " + err.Error())
			}
		}
	}()
	logger.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-ctx.Done()
	logger.Info("Shutting down worker gracefully...")
}
