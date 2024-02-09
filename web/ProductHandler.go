package web

import (
	"net/http"
	"restaurantHTTP"
	"restaurantHTTP/entity"

	"github.com/gorilla/sessions"
)

var storeSession = sessions.NewCookieStore([]byte("faux-token-temporaire"))

func (h *Handler) GetProductTypePage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-name")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {

			username := session.Values["username"].(string)
			data := restaurantHTTP.TemplateData{Titre: "Home Page", Content: entity.User{Username: username}, Error: "", Success: ""}

			h.RenderHtml(data, "product/productType.gohtml")(writer, request)
			return
		}

		http.Redirect(writer, request, "product/productType.gohtml", http.StatusSeeOther)
	}
}
