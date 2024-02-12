package web

import (
	"fmt"
	"net/http"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) ShowRestaurantsPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-basic")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		restaurants, err := h.RestaurantStore.GetAllRestaurants()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Title: "Restaurant Page", Content: restaurants, Error: "Nous n'avons pas compris votre requête", Success: "Bienvenue"}
			h.RenderHtml(writer, data, "pages/restaurants.client.gohtml")
		}
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) ShowMenuByRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-basic")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		restaurants, err := h.RestaurantStore.GetAllRestaurants()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Content: restaurants, Error: "Nous n'avons pas compris votre requête", Success: "Bienvenue"}
			h.RenderHtml(writer, data, "pages/restaurants.menu.gohtml")
		}
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) ShowAddRestaurantAdminPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-basic")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		restaurants, err := h.RestaurantStore.GetRestaurant()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: limiter à un accés admin seulement (il créé le restaurant à la suite d'un email > formulaire de contact restaurateur > admin)
		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Title: "Inscription d'un nouveau restaurant", Content: restaurants}
			h.RenderHtml(writer, data, "pages/restaurants.create.gohtml")
		}
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) ShowRestaurantProfile() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-name")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		restaurants, err := h.RestaurantStore.GetRestaurant()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: doit se faire en fonction de l'id du restaurant
		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Title: "Fiche restaurant", Content: restaurants}
			h.RenderHtml(writer, data, "pages/restaurants.profile.gohtml")
		}
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}
func (h *Handler) ManageRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		RestaurantId, _ := strconv.Atoi(chi.URLParam(request, "restaurantId"))
		h.RenderHtml(writer, restaurantHTTP.TemplateData{Content: RestaurantId}, "pages/restaurants.manage.gohtml")

	}
}
func (h *Handler) GetAllRestaurants() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restaurants, err := h.RestaurantStore.GetAllRestaurants()
		if err != nil {
			// Handle database error
			h.RenderJson(w, http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
			})
			return
		}

		h.RenderJson(w, http.StatusOK, restaurants)
	}
}

func (h *Handler) RegisterRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			http.Error(writer, "cette route n'est disponible qu'en POST", http.StatusBadRequest)
			return
		}

		// Traiter tous les cas d'erreur et renvoyer ces message au front
		// Pourquoi utiliser renderHTML et non redirect - gestion des status d'erreur
		// Interdire de s'inscrire 2 fois
		// Envoyer un email de confirmation d'inscription au restaurateur
		// Attention > rectifier le uint[] dans login birthday

		restaurantName := request.FormValue("restaurant-name")
		restaurantEmail := request.FormValue("restaurant-email")
		restaurantTel := request.FormValue("restaurant-tel")
		restaurantGrade := request.FormValue("restaurant-grade")
		restaurantGradeInt, err := strconv.Atoi(restaurantGrade)
		if err != nil {
			fmt.Println("type parsing for grade failed")
			return
		}
		_, err = h.RestaurantStore.AddRestaurant(entity.Restaurant{Name: restaurantName, Phone: restaurantTel, Mail: restaurantEmail, Grade: restaurantGradeInt, IsValidated: true})
		if err != nil {
			data := restaurantHTTP.TemplateData{Error: "Echec de l'inscription du restaurant"}
			h.RenderHtml(writer, data, "pages/restaurants.create.gohtml")
			return
		}
		data := restaurantHTTP.TemplateData{Title: "Inscription d'un nouveau restaurant", Success: restaurantName + " est inscris dans le FoodCourt."}
		h.RenderHtml(writer, data, "pages/restaurants.create.gohtml")
		return
	}
}
