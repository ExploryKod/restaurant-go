package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pusher/pusher-http-go/v5"
	"log"
	"net/http"
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
		ParseTime:            true,
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

	pusherClient := &pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  true,
	}

	store := database.CreateStore(db)
	mux := web.NewHandler(store, pusherClient)

	err = http.ListenAndServe(os.Getenv("MUX_PORT"), mux)

	if err != nil {
		log.Fatal(err)

		return
	}

}
