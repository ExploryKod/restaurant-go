package web

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"net/http"
	"restaurantHTTP"
	database "restaurantHTTP/mysql"
)

func NewHandler(store *database.Store) *Handler {
	handler := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("src"))
	handler.Handle("/src/*", http.StripPrefix("/src/", fs))

	handler.Get("/", handler.GetHomePage())
	handler.Get("/login", handler.Login())

	handler.Post("/login", handler.Login())

	handler.Get("/signup", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("SIGNUP PAGE !"))
	})
	handler.Post("/signup", handler.Signup())

	handler.Post("/", handler.AddUser())
	handler.Delete("/{id}", handler.DeleteUser())
	handler.Patch("/{id}", handler.ToggleIsSuperadmin())

	// Product
	handler.Get("/productType", handler.GetProductTypePage())

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}

func (h *Handler) jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) RenderHtml(data restaurantHTTP.TemplateData, fileRoute string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tmpl, err := template.ParseFS(restaurantHTTP.EmbedTemplates, "src/templates/layout/layout.gohtml", "src/templates/"+fileRoute)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(writer, "layout", data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
