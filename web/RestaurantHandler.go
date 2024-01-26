package web

import (
	"html/template"
	"net/http"
	"restaurantHTTP"
)

func (h *Handler) GetRestaurantList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := store.Get(request, "session-name")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Titre: "Restaurant Page", Content: "", Error: "", Success: ""}
			tmpl, err := template.ParseFS(restaurantHTTP.EmbedTemplates, "src/templates/layout.gohtml", "src/templates/components/restaurantList.gohtml")
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

// TODO: separer la logique métier des handler de page
func (h *Handler) AddRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		return
	}
}
