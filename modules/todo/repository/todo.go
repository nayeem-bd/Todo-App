package repository

import (
	"context"
	"errors"
	"github.com/nayeem-bd/Todo-App/domain"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetAll(ctx context.Context) ([]*domain.Todo, error) {
	var todos []*domain.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) Create(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	if err := r.db.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) GetByID(ctx context.Context, id int) (*domain.Todo, error) {
	var todo domain.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Update(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	if err := r.db.Save(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}
