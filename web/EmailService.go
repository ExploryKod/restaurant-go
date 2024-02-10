package web

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"restaurantHTTP"
)

type EmailData struct {
	ReceiverEmail  string
	SenderEmail    string
	Subject        string
	errorMessage   string
	successMessage string
	Message        []byte
}

func MakeMailtrapEmail(h *Handler, writer http.ResponseWriter, emailData EmailData, template string) {

	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")

	// Choose auth method
	auth := smtp.PlainAuth("", username, password, smtpHost)
	// Message data
	from := emailData.ReceiverEmail
	to := []string{emailData.SenderEmail}

	// Connect to the server and send message
	smtpUrl := smtpHost + ":" + os.Getenv("SMTP_PORT")

	err := smtp.SendMail(smtpUrl, auth, from, to, emailData.Message)
	if err != nil {
		fmt.Println(err)
		data := restaurantHTTP.TemplateData{Error: emailData.errorMessage}
		h.RenderHtml(writer, data, template)
		return
	}
	data := restaurantHTTP.TemplateData{Success: emailData.successMessage}
	h.RenderHtml(writer, data, template)
	return
}

// with net/mail package
//m := mail.NewMessage()
//
//m.SetHeader("From", "a_franssen@hetic.eu")
//
//m.SetHeader("To", "amauryfranssen@gmail.com", "a_franssen@hetic.eu")
//
//m.SetAddressHeader("Cc", "", "")
//
//m.SetHeader("Subject", "Hello!")
//
//m.SetBody("text/html", "Hello <b>Kate</b> and <i>Noah</i>!")
//
//m.Attach("")
//
//d := mail.NewDialer("smtp.gmail.com", 587, "amauryfranssen@gmail.com", "")
//
//if err := d.DialAndSend(m); err != nil {
//
//panic(err)
//
//}
