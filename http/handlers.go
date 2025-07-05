package http

import (
	"github.com/nayeem-bd/Todo-App/internal/store"
	handler "github.com/nayeem-bd/Todo-App/modules/todo/delivery/http"
	"github.com/nayeem-bd/Todo-App/modules/todo/usecase"
	"gorm.io/gorm"
)

type Handler struct {
	TodoHandler *handler.TodoHandler
}

func RegisterHandlers(db *gorm.DB) *Handler {
	s := store.New(db)

	todoUsecase := usecase.NewTodoUsecase(s)

	return &Handler{
		TodoHandler: handler.NewTodoHandler(todoUsecase),
	}
}
