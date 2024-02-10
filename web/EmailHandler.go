package web

import (
	"net/http"
	"net/smtp"
	"restaurantHTTP"
)

func (h *Handler) AskToAddRestaurantByEmail() http.HandlerFunc {
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
		restaurantSubject := request.FormValue("restaurant-subject")
		restaurantMessage := request.FormValue("restaurant-message")

		//_, err = h.RestaurantStore.AddRestaurant(entity.Restaurant{Name: restaurantName, Phone: restaurantTel, Mail: restaurantEmail})
		// Mailtrap account config

		username := "953143d5103e03"

		password := "143f5c5914b162"

		smtpHost := "sandbox.smtp.mailtrap.io"

		// Use this to send real email (Prod) with MailTrap api account:
		//username := "api"
		//
		//password := "<secret_token>"
		//
		//smtpHost := "live.smtp.mailtrap.io"

		// Choose auth method
		auth := smtp.PlainAuth("", username, password, smtpHost)
		// Message data
		from := restaurantEmail
		to := []string{"a_franssen@hetic.eu"}
		message := []byte("To: a_franssen@hetic.eu\r\n" +
			"From: amaury.fra@restaurantgo.dev\r\n" +
			"\r\n" +
			"Subject: " + restaurantSubject + "\r\n" +
			"\r\n" + restaurantName + restaurantMessage + "\r\n")

		// Connect to the server and send message
		smtpUrl := smtpHost + ":25"

		err := smtp.SendMail(smtpUrl, auth, from, to, message)
		if err != nil {
			//log.Fatal(err)
			data := restaurantHTTP.TemplateData{Error: "Echec de l'envoi d'email, veuillez réessayer"}
			h.RenderHtml(writer, data, "pages/home.gohtml")
			return
		}
		data := restaurantHTTP.TemplateData{Success: "Vous avez bien envoyé votre email"}
		h.RenderHtml(writer, data, "pages/home.gohtml")
		return
	}
}
