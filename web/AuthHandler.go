package web

import (
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"restaurantHTTP"
	"strconv"
	"time"
)

var store = sessions.NewCookieStore([]byte("faux-token-temporaire"))

func (h *Handler) GetHomePage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := store.Get(request, "session-name")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Titre: "Home Page", Content: nil, Error: "", Success: ""}
			tmpl, err := template.ParseFS(restaurantHTTP.EmbedTemplates, "src/templates/layout.gohtml", "src/templates/home.gohtml")
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

func (h *Handler) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			data := restaurantHTTP.TemplateData{Titre: "Login Page"}
			tmpl, err := template.ParseFS(restaurantHTTP.EmbedTemplates, "src/templates/layout.gohtml", "src/templates/login.gohtml")
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

			session, _ := store.Get(request, "session-name")
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

			data := restaurantHTTP.TemplateData{Titre: "Home Page", Content: user, Error: "", Success: "Connexion réussie !"}
			tmpl, err := template.ParseFS(restaurantHTTP.EmbedTemplates, "src/templates/layout.gohtml", "src/templates/home.gohtml")
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(writer, "layout", data)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		} else {
			h.failLogin()(writer, request)
		}
	}
}

func (h *Handler) Signup() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// handler signup from post request
		username := request.FormValue("username")
		password := request.FormValue("password")

		fmt.Println(username, password)

	}
}

func (h *Handler) failLogin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := restaurantHTTP.TemplateData{Titre: "Login Page", Content: nil, Error: "Nom d'utilisateur ou mot de passe incorrect !", Success: ""}
		tmpl, err := template.ParseFS(restaurantHTTP.EmbedTemplates, "src/templates/layout.gohtml", "src/templates/login.gohtml")
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
}
