package web

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"html/template"
	"net/http"
	"restaurantHTTP"
	database "restaurantHTTP/mysql"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("restaurantGo"), nil)
}

func makeToken(name string) string {
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"username": name})
	return tokenString
}

func NewHandler(store *database.Store) *Handler {
	handler := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)

	handler.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fs := http.FileServer(http.Dir("src"))
	handler.Handle("/src/*", http.StripPrefix("/src/", fs))

	handler.Route("/", func(r chi.Router) {

		r.Get("/login", handler.Login())
		r.Post("/login", handler.Login())

		r.Get("/signup", handler.Signup())
		r.Post("/signup", handler.Signup())

		r.Get("/checkEmailAndUsername", handler.checkEmailAndUsername())
	})

	handler.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))

		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/", handler.GetHomePage())
		r.Get("/logout", handler.Logout())

		r.Route("/user", func(r chi.Router) {
			//r.Get("/getAll", handler.GetAllUsers())
			//r.Get("/get/{id}", handler.GetUser())
			r.Post("/add", handler.AddUser())
			r.Delete("/delete/{id}", handler.DeleteUser())
			r.Patch("/modify/{id}", handler.ToggleIsSuperadmin())
		})

		r.Route("/restaurants", func(r chi.Router) {
			r.Get("/", handler.ShowRestaurantsPage())
			r.Get("/menu/{id}", handler.ShowMenuByRestaurant())
			r.Get("/get", handler.GetRestaurants())
			r.Get("/restaurator/{id}", handler.ShowRestaurantProfile())
		})

		r.Route("/admin", func(r chi.Router) {
			r.Get("/register-restaurant", handler.ShowAddRestaurantAdminPage())
		})

		r.Route("/api", func(r chi.Router) {
			r.Post("/restaurant/register", handler.RegisterRestaurant())
		})

	})

	handler.Group(func(r chi.Router) {
		handler.Get("/restaurants", handler.ShowRestaurantsPage())
		handler.Get("/restaurants/menu/{id}", handler.ShowMenuByRestaurant())
		handler.Get("/restaurants/get", handler.GetRestaurants())
		handler.Get("/restaurant/add", handler.AddRestaurant())

		r.Get("/restaurant/menu/{id}", handler.CreateOrder())
		r.Post("/restaurant/orders/create", handler.CreateOrder())
		handler.Post("/restaurant/add", handler.AddRestaurant())
		handler.Get("/restaurator/{id}", handler.ShowRestaurantProfile())
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
