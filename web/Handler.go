package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	database "restaurantHTTP/mysql"
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

	handler.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("HOME PAGE !"))
	})
	handler.Get("/login", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("LOGIN PAGE !"))
	})
	handler.Post("/login", handler.Login())

	// signup
	handler.Get("/signup", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("SIGNUP PAGE !"))
	})
	handler.Post("/signup", handler.Signup())

	handler.Post("/", handler.AddUser())
	handler.Delete("/{id}", handler.DeleteUser())
	handler.Patch("/{id}", handler.ToggleIsSuperAdmin())

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}
