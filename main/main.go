package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/smtp"
	"os"
	database "restaurantHTTP/mysql"
	"restaurantHTTP/web"
)

func main() {
	errdot := godotenv.Load()
	if errdot != nil {
		log.Fatal("Error loading .env file:", errdot)
	}

	conf := mysql.Config{
		User:                 os.Getenv("BDD_USER"),
		Passwd:               os.Getenv("BDD_PASSWORD"),
		Addr:                 os.Getenv("BDD_PORT"),
		DBName:               os.Getenv("BDD_NAME"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	store := database.CreateStore(db)
	mux := web.NewHandler(store)

	err = http.ListenAndServe(os.Getenv("MUX_PORT"), mux)

	if err != nil {
		log.Fatal(err)

		return
	}

}

func sendMail() {
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

	//if err := d.DialAndSend(m); err != nil {
	//
	//	panic(err)
	//
	//}

	// Mailtrap account config

	username := "953143d5103e03"

	password := "143f5c5914b162"

	smtpHost := "sandbox.smtp.mailtrap.io"

	// Prod:
	//username := "api"
	//
	//password := "<secret_token>"
	//
	//smtpHost := "live.smtp.mailtrap.io"

	// Choose auth method and set it up

	auth := smtp.PlainAuth("", username, password, smtpHost)

	// Message data

	from := "amaury.fra@restaurantgo.dev"

	to := []string{"a_franssen@hetic.eu"}

	message := []byte("To: kate.doe@example.com\r\n" +
		"From: john.doe@your.domain\r\n" +
		"\r\n" +
		"Subject: Why aren't you using Mailtrap yet?\r\n" +
		"\r\n" +
		"Here's the space for your great sales pitch\r\n")

	// Connect to the server and send message

	smtpUrl := smtpHost + ":25"

	err := smtp.SendMail(smtpUrl, auth, from, to, message)
	if err != nil {

		log.Fatal(err)

	}
}
