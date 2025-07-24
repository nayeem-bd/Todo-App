package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetupRouter(r *chi.Mux, h *Handler) http.Handler {
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/todos", func(r chi.Router) {
			r.Get("/", h.TodoHandler.GetTodos)
			r.Post("/", h.TodoHandler.CreateTodo)
			r.Get("/{id}", h.TodoHandler.GetTodoByID)
			r.Post("/{id}/complete", h.TodoHandler.CompleteTodo)
		})
	})

	return r
}
