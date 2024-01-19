package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	database "restaurantHTTP/postgressql"
)

func NewHandler(store *database.Store) *Handler {
	handler := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)

	handler.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	handler.Post("/", handler.AddUser())
	handler.Delete("/{id}", handler.DeleteUser())
	handler.Patch("/{id}", handler.ToggleIsSuperAdmin())

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}
