package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nayeem-bd/Todo-App/domain"
)

// MockTodoRepository is a mock implementation of TodoRepository for testing
type MockTodoRepository struct {
	todos       []*domain.Todo
	err         error
	getByIDFunc func(ctx context.Context, id int) (*domain.Todo, error)
	createFunc  func(ctx context.Context, todo *domain.Todo) (*domain.Todo, error)
}

func (m *MockTodoRepository) GetAll(ctx context.Context) ([]*domain.Todo, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.todos, nil
}

func (m *MockTodoRepository) Create(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	if m.createFunc != nil {
		return m.createFunc(ctx, todo)
	}
	if m.err != nil {
		return nil, m.err
	}
	// Simulate creating a todo with an ID
	newTodo := *todo
	newTodo.ID = len(m.todos) + 1
	newTodo.CreatedAt = time.Now()
	newTodo.UpdatedAt = time.Now()
	m.todos = append(m.todos, &newTodo)
	return &newTodo, nil
}

func (m *MockTodoRepository) GetByID(ctx context.Context, id int) (*domain.Todo, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, id)
	}
	if m.err != nil {
		return nil, m.err
	}
	for _, todo := range m.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, errors.New("todo not found")
}

// MockStore is a mock implementation of Store for testing
type MockStore struct {
	todoRepo domain.TodoRepository
}

func (m *MockStore) TodoRepository() domain.TodoRepository {
	return m.todoRepo
}

func TestTodoUsecase_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		todos   []*domain.Todo
		err     error
		wantErr bool
		wantLen int
	}{
		{
			name: "successful get all todos",
			todos: []*domain.Todo{
				{ID: 1, Title: "Test Todo 1", Description: "Description 1", Category: "work"},
				{ID: 2, Title: "Test Todo 2", Description: "Description 2", Category: "personal"},
			},
			err:     nil,
			wantErr: false,
			wantLen: 2,
		},
		{
			name:    "empty todos list",
			todos:   []*domain.Todo{},
			err:     nil,
			wantErr: false,
			wantLen: 0,
		},
		{
			name:    "repository error",
			todos:   nil,
			err:     errors.New("database error"),
			wantErr: true,
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockTodoRepository{
				todos: tt.todos,
				err:   tt.err,
			}
			mockStore := &MockStore{todoRepo: mockRepo}
			usecase := NewTodoUsecase(mockStore)

			ctx := context.Background()
			result, err := usecase.GetAll(ctx)

			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(result) != tt.wantLen {
				t.Errorf("TodoUsecase.GetAll() returned %d todos, want %d", len(result), tt.wantLen)
			}

			if !tt.wantErr && len(result) > 0 {
				for i, todo := range result {
					if todo.ID != tt.todos[i].ID {
						t.Errorf("TodoUsecase.GetAll() todo[%d].ID = %d, want %d", i, todo.ID, tt.todos[i].ID)
					}
				}
			}
		})
	}
}

func TestTodoUsecase_Create(t *testing.T) {
	tests := []struct {
		name         string
		input        *domain.Todo
		err          error
		wantErr      bool
		wantCategory string
	}{
		{
			name: "successful create with category",
			input: &domain.Todo{
				Title:       "Test Todo",
				Description: "Test Description",
				Category:    "work",
			},
			err:          nil,
			wantErr:      false,
			wantCategory: "work",
		},
		{
			name: "successful create with default category",
			input: &domain.Todo{
				Title:       "Test Todo",
				Description: "Test Description",
				Category:    "",
			},
			err:          nil,
			wantErr:      false,
			wantCategory: "default",
		},
		{
			name: "repository error",
			input: &domain.Todo{
				Title:       "Test Todo",
				Description: "Test Description",
				Category:    "work",
			},
			err:          errors.New("database error"),
			wantErr:      true,
			wantCategory: "work",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockTodoRepository{
				err: tt.err,
			}
			mockStore := &MockStore{todoRepo: mockRepo}
			usecase := NewTodoUsecase(mockStore)

			ctx := context.Background()
			result, err := usecase.Create(ctx, tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if result == nil {
					t.Error("TodoUsecase.Create() returned nil result")
					return
				}

				if result.Category != tt.wantCategory {
					t.Errorf("TodoUsecase.Create() category = %v, want %v", result.Category, tt.wantCategory)
				}

				if result.Title != tt.input.Title {
					t.Errorf("TodoUsecase.Create() title = %v, want %v", result.Title, tt.input.Title)
				}

				if result.ID == 0 {
					t.Error("TodoUsecase.Create() should set an ID")
				}
			}
		})
	}
}

func TestTodoUsecase_GetByID(t *testing.T) {
	existingTodo := &domain.Todo{
		ID:          1,
		Title:       "Existing Todo",
		Description: "Existing Description",
		Category:    "work",
	}

	tests := []struct {
		name     string
		id       int
		mockFunc func(ctx context.Context, id int) (*domain.Todo, error)
		wantErr  bool
		wantTodo *domain.Todo
	}{
		{
			name: "successful get by id",
			id:   1,
			mockFunc: func(ctx context.Context, id int) (*domain.Todo, error) {
				if id == 1 {
					return existingTodo, nil
				}
				return nil, errors.New("todo not found")
			},
			wantErr:  false,
			wantTodo: existingTodo,
		},
		{
			name: "todo not found",
			id:   999,
			mockFunc: func(ctx context.Context, id int) (*domain.Todo, error) {
				return nil, errors.New("todo not found")
			},
			wantErr:  true,
			wantTodo: nil,
		},
		{
			name: "repository error",
			id:   1,
			mockFunc: func(ctx context.Context, id int) (*domain.Todo, error) {
				return nil, errors.New("database error")
			},
			wantErr:  true,
			wantTodo: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockTodoRepository{
				getByIDFunc: tt.mockFunc,
			}
			mockStore := &MockStore{todoRepo: mockRepo}
			usecase := NewTodoUsecase(mockStore)

			ctx := context.Background()
			result, err := usecase.GetByID(ctx, tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if result == nil {
					t.Error("TodoUsecase.GetByID() returned nil result")
					return
				}

				if result.ID != tt.wantTodo.ID {
					t.Errorf("TodoUsecase.GetByID() ID = %v, want %v", result.ID, tt.wantTodo.ID)
				}

				if result.Title != tt.wantTodo.Title {
					t.Errorf("TodoUsecase.GetByID() Title = %v, want %v", result.Title, tt.wantTodo.Title)
				}
			}
		})
	}
}
