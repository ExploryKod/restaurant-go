package web

import (
	"html/template"
	"net/http"
	"restaurantHTTP"
)

func (h *Handler) ShowRestaurantsPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := store.Get(request, "session-name")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		restaurants, err := h.RestaurantStore.GetRestaurant()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Titre: "Restaurant Page", Content: restaurants, Error: "Nous n'avons pas compris votre requête", Success: "Bienvenue"}
			tmpl, err := template.ParseFS(restaurantHTTP.EmbedTemplates, "src/templates/layout.gohtml", "src/templates/restaurants.gohtml")
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(writer, "layout", data)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) GetRestaurants() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restaurants, err := h.RestaurantStore.GetRestaurant()
		if err != nil {
			// Handle database error
			h.jsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
			})
			return
		}

		h.jsonResponse(w, http.StatusOK, restaurants)
	}
}

// TODO: separer la logique métier des handler de page
func (h *Handler) AddRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		return
	}
}
