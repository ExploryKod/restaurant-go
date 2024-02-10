package web

import (
	"fmt"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/url"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"time"
)

var storeSession = sessions.NewCookieStore([]byte("faux-token-temporaire"))

func (h *Handler) GetHomePage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-basic")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {

			username := session.Values["username"].(string)
			token := session.Values["token"].(string)
			data := restaurantHTTP.TemplateData{Titre: "Accueil", Content: entity.User{Username: username}, Token: token}

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

			data := restaurantHTTP.TemplateData{Titre: "Login"}
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

		match := CheckPasswordHash(password, user.Password)

		if user.Username == username && match {
			token := makeToken(user.ID, user.Username, user.Mail)

			session, _ := storeSession.Get(request, "session-basic")
			session.Values["token"] = token
			session.Values["authenticated"] = true
			session.Values["username"] = user.Username

			err := session.Save(request, writer)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			http.SetCookie(writer, &http.Cookie{
				HttpOnly: true,
				Expires:  time.Now().Add(7 * 24 * time.Hour),
				SameSite: http.SameSiteLaxMode,
				// Secure: true, // si on souhaite activer le HTTPS
				Name:  "jwt",
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
			data := restaurantHTTP.TemplateData{Titre: "Signup"}

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
			h.RenderHtml(writer, restaurantHTTP.TemplateData{Titre: "Signup", Error: "Username already taken !"}, "auth/signup.gohtml")
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
		session, _ := storeSession.Get(request, "session-basic")
		session.Values["authenticated"] = false
		session.Values["username"] = nil

		err := session.Save(request, writer)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(writer, &http.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Now(),
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			//Secure:   true,
		})

		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) failLogin(writer http.ResponseWriter, request *http.Request) {
	data := restaurantHTTP.TemplateData{Titre: "Login", Error: "Nom d'utilisateur ou mot de passe incorrect !"}

	h.RenderHtml(writer, data, "auth/login.gohtml")
}

func (h *Handler) checkEmailAndUsername() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		username := request.URL.Query().Get("username")
		email := request.URL.Query().Get("email")

		var user *entity.User
		var err error

		EmailAndUsernameCheck := struct {
			Email struct {
				Exists  *bool  `json:"exists"`
				Message string `json:"message"`
			} `json:"email"`
			Username struct {
				Exists  *bool  `json:"exists"`
				Message string `json:"message"`
			} `json:"username"`
		}{
			Email: struct {
				Exists  *bool  `json:"exists"`
				Message string `json:"message"`
			}{
				nil,
				"",
			},
			Username: struct {
				Exists  *bool  `json:"exists"`
				Message string `json:"message"`
			}{
				nil,
				"",
			},
		}

		if username != "" {
			user, err = h.UserStore.GetUserByUsername(username)
			if err != nil {
				fmt.Println(err)
			}
			if user == nil {
				EmailAndUsernameCheck.Username.Exists = new(bool)
				*EmailAndUsernameCheck.Username.Exists = false
				EmailAndUsernameCheck.Username.Message = "Username available !"
			} else {
				EmailAndUsernameCheck.Username.Exists = new(bool)
				*EmailAndUsernameCheck.Username.Exists = true
				EmailAndUsernameCheck.Username.Message = "Username already taken !"
			}
		}

		if email != "" {
			user, err = h.UserStore.GetUserByMail(email)
			if err != nil {
				fmt.Println(err)
			}

			if user == nil {
				EmailAndUsernameCheck.Email.Exists = new(bool)
				*EmailAndUsernameCheck.Email.Exists = false
				EmailAndUsernameCheck.Email.Message = "Email available !"
			} else {
				EmailAndUsernameCheck.Email.Exists = new(bool)
				*EmailAndUsernameCheck.Email.Exists = true
				EmailAndUsernameCheck.Email.Message = "Email already taken !"
			}
		}

		h.RenderJson(writer, http.StatusOK, EmailAndUsernameCheck)
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
