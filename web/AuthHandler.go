package web

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"strconv"
	"time"
)

var storeSession = sessions.NewCookieStore([]byte("faux-token-temporaire"))

func (h *Handler) GetHomePage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-name")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {

			username := session.Values["username"].(string)
			data := restaurantHTTP.TemplateData{Titre: "Home Page", Content: entity.User{Username: username}, Error: "", Success: ""}

			h.RenderHtml(data, "pages/home.gohtml")(writer, request)
			return
		}

		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			data := restaurantHTTP.TemplateData{Titre: "Login Page"}

			h.RenderHtml(data, "auth/login.gohtml")(writer, request)
			return
		}

		username := request.FormValue("username")
		password := request.FormValue("password")

		user, err := h.UserStore.GetUserByUsername(username)
		if err != nil {
			fmt.Println(err)
		}

		if user == nil {
			h.failLogin()(writer, request)
			return
		}

		if user.Username == username && user.Password == password {

			session, _ := storeSession.Get(request, "session-name")
			session.Values["authenticated"] = true
			session.Values["username"] = user.Username

			err := session.Save(request, writer)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			token := strconv.Itoa(user.ID) + "." + user.Username + "." + user.Password
			http.SetCookie(writer, &http.Cookie{
				HttpOnly: true,
				Expires:  time.Now().Add(7 * 24 * time.Hour),
				SameSite: http.SameSiteLaxMode,
				// Uncomment below for HTTPS:
				// Secure: true,
				Name:  "token", // Must be named "jwt" or else the token cannot be searched for by jwtauth.Verifier.
				Value: token,
			})

			http.Redirect(writer, request, "/", http.StatusSeeOther)
		} else {
			h.failLogin()(writer, request)
		}
	}
}

func (h *Handler) Signup() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			data := restaurantHTTP.TemplateData{Titre: "Signup Page"}

			h.RenderHtml(data, "auth/signup.gohtml")(writer, request)
			return
		}
		username := request.FormValue("username")
		password := request.FormValue("password")

		fmt.Println(username, password)
	}
}

func (h *Handler) failLogin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := restaurantHTTP.TemplateData{Titre: "Login Page", Content: nil, Error: "Nom d'utilisateur ou mot de passe incorrect !", Success: ""}

		h.RenderHtml(data, "auth/login.gohtml")(writer, request)
	}
}
