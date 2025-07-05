package domain

import (
	"context"
	"time"
)

type Todo struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DoneAt      *time.Time `json:"done_at"`
}

func (t *Todo) TableName() string {
	return "todos"
}

type TodoRepository interface {
	GetAll(ctx context.Context) ([]*Todo, error)
	Create(ctx context.Context, todo *Todo) (*Todo, error)
}

type TodoUsecase interface {
	GetAll(ctx context.Context) ([]*Todo, error)
	Create(ctx context.Context, todo *Todo) (*Todo, error)
}

type TodoHandler interface {
	GetAll(ctx context.Context) ([]*Todo, error)
	Create(ctx context.Context, todo *Todo) (*Todo, error)
}
