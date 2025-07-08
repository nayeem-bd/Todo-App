package usecase

import (
	"context"
	"encoding/json"
	"github.com/nayeem-bd/Todo-App/domain"
	"github.com/nayeem-bd/Todo-App/internal/config"
	"github.com/nayeem-bd/Todo-App/internal/store"
	"time"
)

type TodoUsecase struct {
	store  store.Store
	cacher *config.Cache
}

func NewTodoUsecase(store store.Store, cacher *config.Cache) *TodoUsecase {
	return &TodoUsecase{store: store, cacher: cacher}
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
