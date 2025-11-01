package main

import (
	"log"
	"net/http"
	"os"
	database "restaurantHTTP/mysql"
	"restaurantHTTP/web"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	// Charger .env seulement si les variables d'environnement ne sont pas déjà définies
	// (comme dans Docker Compose où elles sont définies via environment)
	// Cela permet aux variables Docker de prendre la priorité sur le fichier .env
	if os.Getenv("BDD_PORT") == "" {
		godotenv.Load() // Charge .env seulement si BDD_PORT n'est pas déjà défini
	}

	// Debug: afficher la configuration de connexion (sans le mot de passe)
	log.Printf("Connecting to database: user=%s, addr=%s, dbname=%s",
		os.Getenv("BDD_USER"), os.Getenv("BDD_PORT"), os.Getenv("BDD_NAME"))

	conf := mysql.Config{
		User:   os.Getenv("BDD_USER"),
		Passwd: os.Getenv("BDD_PASSWORD"),
		Addr:   os.Getenv("BDD_PORT"),
		DBName: os.Getenv("BDD_NAME"),
		//User:                 "ueill1e8a2djeyha",
		//Passwd:               "",
		//Addr:                 "bg34o0geswbybq906ljp-mysql.services.clever-cloud.com:3306",
		//DBName:               "bg34o0geswbybq906ljp",
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
