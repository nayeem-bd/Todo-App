package store

import (
	"github.com/nayeem-bd/Todo-App/domain"
	todoRepo "github.com/nayeem-bd/Todo-App/modules/todo/repository"
	"gorm.io/gorm"
)

type Store interface {
	TodoRepository() domain.TodoRepository
}

type DataStore struct {
	db       *gorm.DB
	TodoRepo domain.TodoRepository
}

func New(db *gorm.DB) Store {
	return &DataStore{
		db:       db,
		TodoRepo: todoRepo.NewTodoRepository(db),
	}
}

func (d DataStore) TodoRepository() domain.TodoRepository {
	return d.TodoRepo
}
