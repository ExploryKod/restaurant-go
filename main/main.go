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
	// En production (Render.com), les variables d'environnement sont définies directement
	// Ne charger .env QUE si les variables d'environnement ne sont PAS déjà définies
	// Cela évite que le .env (s'il est présent dans l'image) écrase les variables de production
	bdUser := os.Getenv("BDD_USER")
	bdPort := os.Getenv("BDD_PORT")
	bdName := os.Getenv("BDD_NAME")

	// Si aucune variable n'est définie, on est probablement en développement local
	// Charger le .env seulement dans ce cas
	if bdUser == "" && bdPort == "" && bdName == "" {
		// Mode développement local : charger .env
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found, using environment variables only")
		}
		// Recharger les variables après le chargement du .env
		bdUser = os.Getenv("BDD_USER")
		bdPort = os.Getenv("BDD_PORT")
		bdName = os.Getenv("BDD_NAME")
	}

	// Récupérer les variables d'environnement (priorité aux variables système)
	bdPassword := os.Getenv("BDD_PASSWORD")

	// Debug: afficher la configuration de connexion (sans le mot de passe)
	log.Printf("Connecting to database: user=%s, addr=%s, dbname=%s",
		bdUser, bdPort, bdName)

	// Vérifier que les variables essentielles sont définies
	if bdUser == "" || bdPort == "" || bdName == "" {
		log.Fatal("Missing required database environment variables: BDD_USER, BDD_PORT, BDD_NAME must be set")
		return
	}

	conf := mysql.Config{
		User:                 bdUser,
		Passwd:               bdPassword,
		Addr:                 bdPort,
		DBName:               bdName,
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
