package service

import (
	"net/http"
	"net/smtp"
	"os"
	"restaurantHTTP"
	"restaurantHTTP/web"
)

type EmailData struct {
	ReceiverEmail string
	SenderEmail   string
	Subject       string
	Message       []byte
}

func MakeMailtrapEmail(h *web.Handler, writer http.ResponseWriter, emailData EmailData, template string) {

	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")

	// Choose auth method
	auth := smtp.PlainAuth("", username, password, smtpHost)
	// Message data
	from := emailData.ReceiverEmail
	to := []string{emailData.SenderEmail}

	// Connect to the server and send message
	smtpUrl := smtpHost + os.Getenv("SMTP_PORT")

	err := smtp.SendMail(smtpUrl, auth, from, to, emailData.Message)
	if err != nil {
		//log.Fatal(err)
		data := restaurantHTTP.TemplateData{Error: "Echec de l'envoi d'email, veuillez réessayer"}
		h.RenderHtml(writer, data, template)
		return
	}
	data := restaurantHTTP.TemplateData{Success: "Vous avez bien envoyé votre email"}
	h.RenderHtml(writer, data, template)
	return
}
