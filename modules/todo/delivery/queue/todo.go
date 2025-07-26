package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nayeem-bd/Todo-App/domain"
	"github.com/nayeem-bd/Todo-App/domain/dto"
	amqp "github.com/rabbitmq/amqp091-go"
)

type TodoWorker struct {
	todoUsecase domain.TodoUsecase
}

func NewTodoWorker(todoUsecase domain.TodoUsecase) *TodoWorker {
	return &TodoWorker{
		todoUsecase: todoUsecase,
	}
}

func (w *TodoWorker) ProcessMessage(ctx context.Context, message amqp.Delivery) error {
	var event dto.Event
	if err := json.Unmarshal(message.Body, &event); err != nil {
		return fmt.Errorf("failed to decode message: %w", err)
	}

	switch event.Event {
	case "todo_completed":
		if event.TodoID == nil {
			return fmt.Errorf("todo ID is required for todo_completed event")
		}
		err := w.todoUsecase.CompleteTodo(ctx, *event.TodoID)
		if err != nil {
			return fmt.Errorf("failed to complete todo: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown event type: %s", event.Event)
	}
}
