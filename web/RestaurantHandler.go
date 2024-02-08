package web

import (
	"html/template"
	"net/http"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"time"
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
			tmpl, err := template.ParseFS(restaurantHTTP.EmbedTemplates, "src/templates/layout/layout.gohtml", "src/templates/pages/restaurants.client.gohtml")
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
			data := restaurantHTTP.TemplateData{Title: "Restaurant Page", Content: restaurants, Error: "Nous n'avons pas compris votre requête", Success: "Bienvenue"}
			h.RenderHtml(writer, data, "pages/restaurants.menu.gohtml")
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

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Title: "Restaurant Profile", Content: restaurants, Error: "Nous n'avons pas compris votre requête", Success: "Bienvenue"}
			h.RenderHtml(writer, data, "pages/restaurants.create.gohtml")
		}
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) GetRestaurants() http.HandlerFunc {
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

func (h *Handler) getRestaurantById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		return
	}
}

// TODO: separer la logique métier des handler de page
func (h *Handler) AddRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		_, err := h.RestaurantStore.AddRestaurant(entity.Restaurant{1, "name", "/logo", "img/", "0367", "mail.com", true, time.Now(), time.Now(), 2, true})
		if err != nil {
			// Handle database error
			h.RenderJson(writer, http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
			})
			return
		}
		return
	}
}
