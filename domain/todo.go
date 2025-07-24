package domain

import (
	"context"
	"time"
)

type Todo struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"type:varchar(100);not null"`
	Description string     `json:"description" gorm:"type:varchar(255);not null"`
	Category    string     `json:"category" gorm:"type:varchar(50);default:'default'"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DoneAt      *time.Time `json:"done_at" gorm:"type:timestamp;default:null"`
}

func (t *Todo) TableName() string {
	return "todos"
}

type TodoRepository interface {
	GetAll(ctx context.Context) ([]*Todo, error)
	Create(ctx context.Context, todo *Todo) (*Todo, error)
	GetByID(ctx context.Context, id int) (*Todo, error)
}

type TodoUsecase interface {
	GetAll(ctx context.Context) ([]*Todo, error)
	Create(ctx context.Context, todo *Todo) (*Todo, error)
	GetByID(ctx context.Context, id int) (*Todo, error)
	Complete(ctx context.Context, id int) error
}
