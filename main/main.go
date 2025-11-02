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
	// Détecter si on est sur Render.com
	// Render.com définit automatiquement RENDER_EXTERNAL_HOSTNAME pour tous les services web
	// Cette variable est automatiquement injectée par Render.com (pas besoin de la définir manuellement)
	isRender := os.Getenv("RENDER_EXTERNAL_HOSTNAME") != ""

	// En production Render.com, NE JAMAIS charger .env
	// En développement local, charger .env seulement si les variables ne sont pas définies
	if !isRender {
		bdUser := os.Getenv("BDD_USER")
		bdPort := os.Getenv("BDD_PORT")
		bdName := os.Getenv("BDD_NAME")

		// Si aucune variable n'est définie, charger .env (développement local)
		if bdUser == "" && bdPort == "" && bdName == "" {
			if err := godotenv.Load(); err != nil {
				log.Println("Warning: .env file not found, using environment variables only")
			}
		}
	}

	// Récupérer les variables d'environnement (priorité aux variables système)
	bdUser := os.Getenv("BDD_USER")
	bdPassword := os.Getenv("BDD_PASSWORD")
	bdPort := os.Getenv("BDD_PORT")
	bdName := os.Getenv("BDD_NAME")

	// Debug: Afficher toutes les variables d'environnement pour diagnostic
	log.Printf("=== Database Configuration Debug ===")
	log.Printf("Is Render.com environment: %v", isRender)
	if isRender {
		log.Printf("RENDER_EXTERNAL_HOSTNAME: %s", os.Getenv("RENDER_EXTERNAL_HOSTNAME"))
	}
	log.Printf("BDD_USER: %s", bdUser)
	log.Printf("BDD_PORT: %s", bdPort)
	log.Printf("BDD_NAME: %s", bdName)
	log.Printf("BDD_PASSWORD: %s", map[bool]string{true: "***SET***", false: "NOT SET"}[bdPassword != ""])
	log.Printf("Connecting to database: user=%s, addr=%s, dbname=%s", bdUser, bdPort, bdName)

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
