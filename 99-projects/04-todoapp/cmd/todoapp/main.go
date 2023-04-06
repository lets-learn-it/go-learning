package main

import (
	"net/http"

	v1routes "avabodha.in/todoapp/api/v1/router"
	"avabodha.in/todoapp/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := routes()

	_ = app.NewApp(&r)

	// set up the application config

	http.ListenAndServe(":3000", r)
}

func routes() http.Handler {
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	// mount v1 version
	mux.Mount("/todo", v1routes.Routes())

	return mux
}
