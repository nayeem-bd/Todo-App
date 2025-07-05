package usecase

import (
	"context"
	"github.com/nayeem-bd/Todo-App/domain"
	"github.com/nayeem-bd/Todo-App/internal/store"
)

type TodoUsecase struct {
	store store.Store
}

func NewTodoUsecase(store store.Store) *TodoUsecase {
	return &TodoUsecase{store: store}
}

func (todoUsecase *TodoUsecase) GetAll(ctx context.Context) ([]*domain.Todo, error) {
	return todoUsecase.store.TodoRepository().GetAll(ctx)
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
