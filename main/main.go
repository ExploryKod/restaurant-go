package main

import (
	"github.com/go-mail/mail"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	database "restaurantHTTP/mysql"
	"restaurantHTTP/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	m := mail.NewMessage()

	m.SetHeader("From", "a_franssen@hetic.eu")

	m.SetHeader("To", "amauryfranssen@gmail.com", "a_franssen@hetic.eu")

	m.SetAddressHeader("Cc", "", "")

	m.SetHeader("Subject", "Hello!")

	m.SetBody("text/html", "Hello <b>Kate</b> and <i>Noah</i>!")

	m.Attach("")

	d := mail.NewDialer("smtp.gmail.com", 587, "amauryfranssen@gmail.com", "")

	if err := d.DialAndSend(m); err != nil {

		panic(err)

	}

	conf := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Addr:                 os.Getenv("BDD_PORT"),
		DBName:               "restaurantbdd",
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
