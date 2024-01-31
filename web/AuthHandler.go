package web

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/url"
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
			data := restaurantHTTP.TemplateData{Title: "Home", Content: entity.User{Username: username}}

			h.RenderHtml(writer, data, "pages/home.gohtml")
			return
		}

		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			encodedMessage := request.URL.Query().Get("success")
			decodedMessage, err := url.QueryUnescape(encodedMessage)
			if err != nil {
				log.Println("Error during decoding success msg :", err)
				decodedMessage = ""
			}

			data := restaurantHTTP.TemplateData{Title: "Login"}
			if decodedMessage != "" {
				data.Success = decodedMessage
			}

			h.RenderHtml(writer, data, "auth/login.gohtml")
			return
		}

		username := request.FormValue("username")
		password := request.FormValue("password")

		user, err := h.UserStore.GetUserByUsername(username)
		if err != nil {
			fmt.Println(err)
		}

		if user == nil {
			h.failLogin(writer, request)
			return
		}

		//hash, _ := HashPassword(password)
		//fmt.Println("Hash :", hash)
		match := CheckPasswordHash(password, user.Password)
		fmt.Println("Match :", match)
		fmt.Println("user password :", user.Password)

		if user.Username == username && match {
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
				// HTTPS en dessous
				// Secure: true,
				Name:  "token",
				Value: token,
			})

			http.Redirect(writer, request, "/", http.StatusSeeOther)
		} else {
			h.failLogin(writer, request)
		}
	}
}

func (h *Handler) Signup() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			data := restaurantHTTP.TemplateData{Title: "Signup"}

			h.RenderHtml(writer, data, "auth/signup.gohtml")
			return
		}
		username := request.FormValue("username")
		password := request.FormValue("password")
		name := request.FormValue("name")
		firstname := request.FormValue("firstname")
		mail := request.FormValue("mail")
		phone := request.FormValue("phone")
		//birthday := request.FormValue("birthday")

		response, err := h.UserStore.GetUserByUsername(username)
		if err != nil {
			fmt.Println(response)
			return
		}

		if response != nil {
			h.RenderHtml(writer, restaurantHTTP.TemplateData{Title: "Signup", Error: "Username already taken !"}, "auth/signup.gohtml")
			return
		}

		hashedPassword, _ := HashPassword(password)

		user := &entity.User{
			Username:     username,
			Password:     hashedPassword,
			Name:         name,
			Firstname:    firstname,
			Mail:         mail,
			Phone:        phone,
			IsSuperadmin: false,
			Birthday:     sql.NullTime{},
		}

		var id int
		id, err = h.UserStore.AddUser(user)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("New user id :", id)

		encodedMessage := url.QueryEscape("Account created successfully !")
		http.Redirect(writer, request, "/login?success="+encodedMessage, http.StatusSeeOther)
	}
}

func (h *Handler) Logout() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		session, _ := storeSession.Get(request, "session-name")
		session.Values["authenticated"] = false
		session.Values["username"] = nil

		err := session.Save(request, writer)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) failLogin(writer http.ResponseWriter, request *http.Request) {
	data := restaurantHTTP.TemplateData{Title: "Login", Error: "Nom d'utilisateur ou mot de passe incorrect !"}

	h.RenderHtml(writer, data, "auth/login.gohtml")
}

func (h *Handler) checkUsername() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		username := request.URL.Query().Get("username")

		user, err := h.UserStore.GetUserByUsername(username)
		if err != nil {
			fmt.Println(err)
		}

		if user == nil {
			h.RenderJson(writer, http.StatusOK, struct {
				Exists  bool   `json:"exists"`
				Message string `json:"message"`
			}{
				true,
				"Username available !",
			})
		} else {
			h.RenderJson(writer, http.StatusOK, struct {
				Exists  bool   `json:"exists"`
				Message string `json:"message"`
			}{
				false,
				"Username already taken !",
			})
		}
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
