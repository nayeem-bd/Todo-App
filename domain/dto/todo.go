package dto

import "github.com/nayeem-bd/Todo-App/domain"

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=150"`
	Description string `json:"description" validate:"required,min=5,max=500"`
	Category    string `json:"category" validate:"omitempty,max=50"`
}

func (req *CreateTodoRequest) ToDomain() *domain.Todo {
	return &domain.Todo{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
	}
}
