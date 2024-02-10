package web

import (
	"net/http"
	"restaurantHTTP"
	"restaurantHTTP/entity"
)

func (h *Handler) GetProductTypePage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-name")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {

			username := session.Values["username"].(string)
			data := restaurantHTTP.TemplateData{Titre: "Choix catégorie produit", Content: entity.User{Username: username}, Error: "", Success: ""}
			h.RenderHtml(writer, data, "product/productType.gohtml")
			return
		}

		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}
