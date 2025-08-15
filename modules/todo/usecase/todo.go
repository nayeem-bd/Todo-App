package usecase

import (
	"context"
	"encoding/json"
	"github.com/nayeem-bd/Todo-App/domain"
	"github.com/nayeem-bd/Todo-App/internal/config"
	"github.com/nayeem-bd/Todo-App/internal/logger"
	"github.com/nayeem-bd/Todo-App/internal/store"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type TodoUsecase struct {
	store  store.Store
	cacher *config.Cache
	queue  *config.Queue
}

func NewTodoUsecase(store store.Store, cacher *config.Cache, queue *config.Queue) *TodoUsecase {
	return &TodoUsecase{store: store, cacher: cacher, queue: queue}
}

func (todoUsecase *TodoUsecase) GetAll(ctx context.Context) ([]*domain.Todo, error) {
	todoStr, err := todoUsecase.cacher.Get(ctx, "todos")
	var todos []*domain.Todo
	if err == nil {
		if err := json.Unmarshal([]byte(todoStr), &todos); err == nil {
			return todos, err
		}
	}

	todos, err = todoUsecase.store.TodoRepository().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Cache the todos
	todoBytes, err := json.Marshal(todos)
	if err != nil {
		return nil, err
	}
	_ = todoUsecase.cacher.Set(ctx, "todos", string(todoBytes), 30*time.Second)

	return todos, nil
}

func (todoUsecase *TodoUsecase) Create(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	if todo.Category == "" {
		todo.Category = "default"
	}

	createdTodo, err := todoUsecase.store.TodoRepository().Create(ctx, todo)
	if err != nil {
		return nil, err
	}
	return createdTodo, nil
}

func (todoUsecase *TodoUsecase) GetByID(ctx context.Context, id int) (*domain.Todo, error) {
	return todoUsecase.store.TodoRepository().GetByID(ctx, id)
}

func (todoUsecase *TodoUsecase) Complete(ctx context.Context, id int) error {
	todo, err := todoUsecase.GetByID(ctx, id)
	if err != nil {
		return err
	}

	ch, err := todoUsecase.queue.Conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	message := map[string]interface{}{
		"todo_id": todo.ID,
		"event":   "todo_completed",
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = ch.Publish(
		todoUsecase.queue.ExchangeName,
		todoUsecase.queue.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBytes,
		},
	)

	return err
}

func (todoUsecase *TodoUsecase) CompleteTodo(ctx context.Context, id int) error {
	todo, err := todoUsecase.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if todo.DoneAt != nil {
		logger.Info("Todo already completed ", "todo_id: ", todo.ID)
		return nil
	}
	now := time.Now()
	todo.DoneAt = &now

	_, err = todoUsecase.store.TodoRepository().Update(ctx, todo)

	return err
}
