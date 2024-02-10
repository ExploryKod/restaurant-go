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

func makeToken(id int, username string, mail string) string {
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"id": id, "username": username, "mail": mail})
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
			//r.Get("/{id}/get", handler.GetUser())
			r.Post("/add", handler.AddUser())
			r.Delete("/delete/{id}", handler.DeleteUser())
			r.Patch("/modify/{id}", handler.ToggleIsSuperadmin())
		})

		r.Route("/restaurant", func(r chi.Router) {
			r.Get("/", handler.ShowRestaurantsPage())

			r.Get("/get/all", handler.GetAllRestaurants())

			r.Get("/{id}/menu", handler.CreateOrder())

			r.Post("/{id}/create-order", handler.CreateOrder())

			r.Get("/become-restaurant", handler.ShowBecomeRestaurantPage())

			r.Get("/admin/{id}", handler.ShowAdminRestaurantPage())
			// TODO: ajout de {id} ici pour isoler l'id restaurant et l'adjoindre dans le handler pour populer la rable restaurantHasTag
			r.Post("/tag/add", handler.AddTagToRestaurant())

		})

		r.Route("/order", func(r chi.Router) {
			//r.Get("/get/all", handler.GetAllOrders())
			//r.Get("/{id}", handler.GetOrder())
			//r.Post("/add", handler.AddOrder())
			//r.Delete("/delete/{id}", handler.DeleteOrder())
		})

		r.Route("/admin", func(r chi.Router) {
			r.Get("/register-restaurant", handler.ShowAddRestaurantAdminPage())
		})

		r.Route("/email", func(r chi.Router) {
			r.Post("/create-restaurant", handler.AskToAddRestaurantByEmail())
		})

		r.Route("/api", func(r chi.Router) {
			r.Post("/restaurant/register", handler.RegisterRestaurant())
		})

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
