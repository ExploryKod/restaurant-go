package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"strings"
	"restaurantHTTP"
	database "restaurantHTTP/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("restaurantGo"), nil)
}

func makeToken(id int, username string, mail string, isSuperadmin bool) string {
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"id": id, "username": username, "mail": mail, "isSuperadmin": isSuperadmin})
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

	// Servir les fichiers statiques avec les bons MIME types
	fs := http.FileServer(http.Dir("src"))
	// Wrapper pour forcer les bons Content-Type
	fileServer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Définir les MIME types pour les extensions courantes
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		} else if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		} else if strings.HasSuffix(r.URL.Path, ".png") {
			w.Header().Set("Content-Type", "image/png")
		} else if strings.HasSuffix(r.URL.Path, ".jpg") || strings.HasSuffix(r.URL.Path, ".jpeg") {
			w.Header().Set("Content-Type", "image/jpeg")
		} else if strings.HasSuffix(r.URL.Path, ".svg") {
			w.Header().Set("Content-Type", "image/svg+xml")
		}
		fs.ServeHTTP(w, r)
	})
	handler.Handle("/src/*", http.StripPrefix("/src/", fileServer))

	handler.Route("/", func(r chi.Router) {

		r.Get("/login", handler.Login())
		r.Post("/login", handler.Login())

		r.Get("/signup", handler.Signup())
		r.Post("/signup", handler.Signup())

		r.Get("/checkEmailAndUsername", handler.checkEmailAndUsername())

		// Endpoint pour obtenir une image Unsplash par mot-clé
		r.Get("/api/unsplash/image", handler.GetUnsplashImage())
	})

	handler.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		
		// Middleware personnalisé pour rediriger vers /login si pas de token
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				token, _, err := jwtauth.FromContext(r.Context())
				if err != nil || token == nil {
					// Rediriger vers /login si pas de token
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
				next.ServeHTTP(w, r)
			})
		})

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

			r.Get("/delete/{id}", handler.DeleteRestaurantHandler())

			r.Get("/become-restaurant", handler.ShowBecomeRestaurantPage())

			r.Post("/restaurant/register", handler.RegisterRestaurant())

			r.Get("/manage-restaurants", handler.ShowAddRestaurantAdminPage())

			r.Post("/update", handler.UpdateRestaurantHandler())

			r.Get("/show/restaurant-update/{id}", handler.ShowRestaurantUpdatePage())

			r.Get("/admin/{id}", handler.ShowAdminRestaurantPage())
			// TODO: ajout de {id} ici pour isoler l'id restaurant et l'adjoindre dans le handler pour populer la rable restaurantHasTag
			r.Post("/tag/add", handler.AddTagToRestaurant())

		})

		r.Route("/order", func(r chi.Router) {
			r.Get("/", handler.ShowOrdersPage())
			r.Get("/get/all", handler.GetAllOrders())
			r.Get("/{id}", handler.GetOrder())
			//r.Post("/add", handler.AddOrder())
			//r.Delete("/delete/{id}", handler.DeleteOrder())
		})

		r.Route("/email", func(r chi.Router) {
			r.Post("/create-restaurant", handler.AskToAddRestaurantByEmail())
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
	var tmpl *template.Template
	var err error

	// Fonctions helper pour les templates
	funcMap := template.FuncMap{
		"fixImagePath": func(path string) string {
			if path == "" {
				return ""
			}
			// Si le chemin commence déjà par /src/, le retourner tel quel
			if len(path) >= 5 && path[:5] == "/src/" {
				return path
			}
			// Si le chemin commence par image/, supprimer le préfixe et utiliser assets/
			if len(path) >= 6 && path[:6] == "image/" {
				return "/src/assets/" + path[6:]
			}
			// Si le chemin commence par logo/, supprimer le préfixe et utiliser assets/
			if len(path) >= 5 && path[:5] == "logo/" {
				return "/src/assets/" + path[5:]
			}
			// Sinon, ajouter /src/assets/ devant
			return "/src/assets/" + path
		},
	}

	// En développement, charger depuis le système de fichiers pour le hot reload
	// En production, utiliser les templates embeddés
	if os.Getenv("DEV_MODE") != "" || os.Getenv("AIR") != "" {
		// Mode développement : charger depuis le système de fichiers
		tmpl, err = template.New("layout").Funcs(funcMap).ParseFiles("src/templates/layout/layout.gohtml", "src/templates/"+fileRoute)
	} else {
		// Mode production : utiliser les templates embeddés
		tmpl, err = template.New("layout").Funcs(funcMap).ParseFS(restaurantHTTP.EmbedTemplates, "src/templates/layout/layout.gohtml", "src/templates/"+fileRoute)
	}

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

// UnsplashImageResponse représente la réponse de l'API Unsplash pour une recherche
type UnsplashImageResponse struct {
	Results []struct {
		ID   string `json:"id"`
		URLs struct {
			Small string `json:"small"`
		} `json:"urls"`
	} `json:"results"`
}

// GetUnsplashImage génère une URL d'image Unsplash basée sur un mot-clé
func (h *Handler) GetUnsplashImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		if query == "" {
			http.Error(w, "Query parameter is required", http.StatusBadRequest)
			return
		}

		unsplashAccessKey := os.Getenv("UNSPLASH_ACCESS_KEY")
		if unsplashAccessKey == "" {
			// Si pas de clé API, utiliser une URL par défaut (fallback)
			// Note: Cette URL peut ne pas fonctionner car source.unsplash.com est déprécié
			imageURL := fmt.Sprintf("https://source.unsplash.com/300x150/?%s", url.QueryEscape(query))
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"url": imageURL})
			return
		}

		// Appeler l'API Unsplash pour rechercher une photo
		apiURL := fmt.Sprintf("https://api.unsplash.com/search/photos?query=%s&per_page=1&client_id=%s", url.QueryEscape(query), unsplashAccessKey)

		resp, err := http.Get(apiURL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error calling Unsplash API: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var unsplashResp UnsplashImageResponse
		if err := json.NewDecoder(resp.Body).Decode(&unsplashResp); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding Unsplash response: %v", err), http.StatusInternalServerError)
			return
		}

		imageURL := ""
		if len(unsplashResp.Results) > 0 {
			imageURL = unsplashResp.Results[0].URLs.Small
		} else {
			// Fallback si aucune image trouvée
			imageURL = fmt.Sprintf("https://source.unsplash.com/300x150/?%s", url.QueryEscape(query))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"url": imageURL})
	}
}
