package web

import (
	"encoding/json"
	"html/template"
	"net/http"
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

		r.Route("/restaurantUser", func(r chi.Router) {
			r.Get("/list/{restaurantId}", handler.RestaurantUserList())
			r.Get("/create/{restaurantId}", handler.RestaurantUserCreate())
			r.Post("/create/{restaurantId}", handler.RestaurantUserCreate())
			r.Get("/update/{id}/{restaurantId}", handler.RestaurantUserUpdate())
			r.Post("/update/{id}/{restaurantId}", handler.RestaurantUserUpdate())
			r.Get("/delete/{id}/{restaurantId}", handler.RestaurantUserDelete())
		})
		r.Route("/user", func(r chi.Router) {
			//r.Get("/getAll", handler.GetAllUsers())
			//r.Get("/{id}/get", handler.GetUser())
			r.Post("/add", handler.AddUser())
			r.Delete("/delete/{id}", handler.DeleteUser())
			r.Patch("/modify/{id}", handler.ToggleIsSuperadmin())
		})

		r.Route("/restaurant", func(r chi.Router) {
			r.Get("/", handler.ShowRestaurantsPage())
			r.Get("/restaurant/menu/{id}", handler.CreateOrder())
			r.Post("/restaurant/orders/create", handler.CreateOrder())
			r.Get("/get/all", handler.GetAllRestaurants())
			r.Get("/{id}", handler.ShowRestaurantProfile())
			r.Get("/{id}/menu", handler.CreateOrder())

			r.Post("/{id}/create-order", handler.CreateOrder())
			r.Get("/order/get/{restaurantId}", handler.GetAllOrdersByRestaurantId())
			r.Get("/order/validate/{id}", handler.ValidateOrderById())
			r.Get("/order/done/{id}", handler.CompleteOrderById())
			r.Get("/order/ready/{id}", handler.ReadyOrderById())
			r.Get("/restaurator/{id}", handler.ShowRestaurantProfile())
			r.Get("/manage/{restaurantId}", handler.ManageRestaurant())
		})

		r.Route("/order", func(r chi.Router) {
			r.Get("/", handler.ShowOrdersPage())
			r.Get("/get/all", handler.GetAllOrders())
			r.Get("/{id}", handler.GetOrder())

			//r.Post("/add", handler.AddOrder())
			//r.Delete("/delete/{id}", handler.DeleteOrder())
		})

		r.Route("/product", func(r chi.Router) {
			r.Get("/list/{restaurantId}", handler.ListProducts())
			r.Get("/list/delete/{id}/{restaurantId}", handler.DeleteProducts())
			r.Get("/type/create/{restaurantId}", handler.AddProductType())
			r.Post("/type/create/{restaurantId}", handler.AddProductType())

			r.Get("/create/{restaurantId}", handler.AddProduct())
			r.Post("/create/{restaurantId}", handler.AddProduct())
		})
		r.Route("/restaurants", func(r chi.Router) {
			r.Get("/", handler.ShowRestaurantsPage())
			r.Get("/menu/{id}", handler.ShowMenuByRestaurant())
			r.Get("/get", handler.GetAllRestaurants())
			r.Get("/{id}", handler.ShowRestaurantProfile())
			r.Get("/restaurants/menu/{id}", handler.ShowMenuByRestaurant())
			r.Get("/restaurants/get", handler.GetAllRestaurants())
			r.Get("/restaurator/{id}", handler.ShowRestaurantProfile())
		})
		r.Route("/admin", func(r chi.Router) {
			r.Get("/register-restaurant", handler.ShowAddRestaurantAdminPage())
		})

		/*pusherClient := pusher.Client{
			AppID:   os.Getenv("PUSHER_APP_ID"),
			Key:     os.Getenv("PUSHER_KEY"),
			Secret:  os.Getenv("PUSHER_SECRET"),
			Cluster: os.Getenv("PUSHER_CLUSTER"),
			Secure:  true,
		}*/
		r.Route("/api", func(r chi.Router) {
			/*r.Get("/pusher/auth", func(w http.ResponseWriter, r *http.Request) {
				params, _ := r.URL.Query()["channel_name"]
				channelName := params[0]
				userID := r.Context().Value("jwt").(*jwtauth.Token).Claims["id"].(float64)
				userIDString := fmt.Sprintf("%v", userID)
				response, err := pusherClient.AuthenticatePrivateChannel(channelName, userIDString)
				if err != nil {
					http.Error(w, "Forbidden", http.StatusForbidden)
					return
				}
				handler.RenderJson(w, http.StatusOK, response)
			})*/
			r.Post("/order/{id}/create", handler.CreateOrder())

		})

	})

	return handler
}

type Handler struct {
	*chi.Mux
	*database.Store
}

func (h *Handler) RenderJson(writer http.ResponseWriter, status int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
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
