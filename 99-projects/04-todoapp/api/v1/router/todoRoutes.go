package router

import (
	"net/http"

	"avabodha.in/todoapp/api/v1/handler"
	"github.com/go-chi/chi/v5"
)

func todoRoutes() http.Handler {
	todo := chi.NewRouter()

	todo.Get("/{id}", handler.GetTodo)
	todo.Get("/", handler.GetTodos)
	todo.Post("/", handler.AddTodo)
	todo.Put("/{id}", handler.EditTodo)
	todo.Delete("/{id}", handler.DeleteTodo)

	return todo
}
