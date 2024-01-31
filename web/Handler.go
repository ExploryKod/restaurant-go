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

	handler.Route("/", func(r chi.Router) {
		r.Get("/", handler.GetHomePage())

		r.Get("/login", handler.Login())
		r.Post("/login", handler.Login())

		r.Get("/signup", handler.Signup())
		r.Post("/signup", handler.Signup())

		r.Get("/checkEmailAndUsername", handler.checkEmailAndUsername())

		r.Get("/logout", handler.Logout())
	})

	handler.Route("/user", func(r chi.Router) {
		//r.Get("/getAll", handler.GetAllUsers())
		//r.Get("/get/{id}", handler.GetUser())
		r.Post("/add", handler.AddUser())
		r.Delete("/delete/{id}", handler.DeleteUser())
		r.Patch("/modify/{id}", handler.ToggleIsSuperadmin())
	})

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}

func (h *Handler) RenderJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) RenderHtml(writer http.ResponseWriter, data restaurantHTTP.TemplateData, fileRoute string) {
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
