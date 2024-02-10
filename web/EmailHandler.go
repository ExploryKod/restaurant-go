package web

import (
	"net/http"
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
		template := "pages/restaurants.subscribe.gohtml"

		message := []byte("" +
			"To: a_franssen@hetic.eu\r\n" +
			"From:" + restaurantEmail +
			"\r\n" +
			"Sujet: " + restaurantSubject + "\r\n" +
			"\r\n" +
			"Nom du restaurant:\n" + restaurantName + "\r\n" + restaurantMessage + "\r\n")

		emailData := EmailData{
			ReceiverEmail:  restaurantEmail,
			SenderEmail:    "restaurago@goemail.com",
			Subject:        restaurantSubject,
			errorMessage:   "Votre message n'est pas parvenu, veuillez réessayer",
			successMessage: "Votre email a bien été envoyé, nous vous contactons dans les meilleurs délais",
			Message:        message,
		}

		MakeMailtrapEmail(h, writer, emailData, template)
	}
}

//func (h *Handler) AskToAddRestaurantByEmailWithHTML() http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		if request.Method != "POST" {
//			http.Error(writer, "cette route n'est disponible qu'en POST", http.StatusBadRequest)
//			return
//		}
//
//		// Traiter tous les cas d'erreur et renvoyer ces message au front
//		// Pourquoi utiliser renderHTML et non redirect - gestion des status d'erreur
//		// Interdire de s'inscrire 2 fois
//		// Envoyer un email de confirmation d'inscription au restaurateur
//		// Attention > rectifier le uint[] dans login birthday
//
//		restaurantName := request.FormValue("restaurant-name")
//		restaurantEmail := request.FormValue("restaurant-email")
//		restaurantSubject := request.FormValue("restaurant-subject")
//		restaurantMessage := request.FormValue("restaurant-message")
//
//		// Mailtrap account config
//		username := "953143d5103e03"
//		password := "143f5c5914b162"
//		smtpHost := "sandbox.smtp.mailtrap.io"
//
//		// Choose auth method
//		auth := smtp.PlainAuth("", username, password, smtpHost)
//
//		// Create a new email message
//		m := mail.NewMessage()
//		m.SetHeader("From", restaurantEmail)
//		m.SetHeader("To", "a_franssen@hetic.eu")
//		m.SetHeader("Subject", restaurantSubject)
//
//		// Set the email body as HTML
//		m.SetBody("text/html", "<strong>"+restaurantName+"</strong><br/>"+restaurantMessage)
//
//		// Setup the dialer with the SMTP server details
//		d := mail.NewDialer(smtpHost, 25, username, password)
//
//		// Send the email
//		if err := d.DialAndSend(m); err != nil {
//			data := restaurantHTTP.TemplateData{Error: "Echec de l'envoi d'email, veuillez réessayer"}
//			h.RenderHtml(writer, data, "pages/home.gohtml")
//			return
//		}
//
//		data := restaurantHTTP.TemplateData{Success: "Vous avez bien envoyé votre email"}
//		h.RenderHtml(writer, data, "pages/home.gohtml")
//		return
//	}
//}
