package handler

import (
	"encoding/json"
	"github.com/nayeem-bd/Todo-App/domain"
	"github.com/nayeem-bd/Todo-App/domain/dto"
	"github.com/nayeem-bd/Todo-App/internal/utils"
	"net/http"
)

type TodoHandler struct {
	todoUsecase domain.TodoUsecase
	validator   *utils.Validator
}

func NewTodoHandler(todoUsecase domain.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
		validator:   utils.NewValidator(),
	}
}

func (todoHandler *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := todoHandler.todoUsecase.GetAll(r.Context())

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to fetch todos", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "Todos retrieved successfully", todos)
}

func (todoHandler *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Validate the request
	if validationErrors := todoHandler.validator.Validate(&req); len(validationErrors) > 0 {
		utils.WriteError(w, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	todo := req.ToDomain()

	createdTodo, err := todoHandler.todoUsecase.Create(r.Context(), todo)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to create todo", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusCreated, "Todo created successfully", createdTodo)
}
