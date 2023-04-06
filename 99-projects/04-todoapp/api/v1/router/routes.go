package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Mount("/v1", todoRoutes())

	return mux
}
