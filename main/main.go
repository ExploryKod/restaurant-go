package main

import (
	"log"
	"net/http"
	"os"
	database "restaurantHTTP/mysql"
	"restaurantHTTP/web"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	// errdot := godotenv.Load()
	// if errdot != nil {
	// 	log.Error("Error loading .env file:", errdot)
	// }

	conf := mysql.Config{
		// User:                 os.Getenv("BDD_USER"),
		// Passwd:               os.Getenv("BDD_PASSWORD"),
		// Addr:                 os.Getenv("BDD_PORT"),
		// DBName:               os.Getenv("BDD_NAME"),
		User:                 "ueill1e8a2djeyha",
		Passwd:               "qbpkI0vidAOA5qha2Q6X",
		Addr:                 "bg34o0geswbybq906ljp-mysql.services.clever-cloud.com:3306"
		DBName:               "bg34o0geswbybq906ljp",
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

	store := database.CreateStore(db)
	mux := web.NewHandler(store)

	err = http.ListenAndServe(":9999", mux)

	if err != nil {
		log.Fatal(err)

		return
	}

}
