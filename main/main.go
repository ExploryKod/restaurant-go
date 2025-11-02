package main

import (
	"log"
	"net/http"
	"os"
	database "restaurantHTTP/mysql"
	"restaurantHTTP/web"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	// Log IMMÉDIATEMENT pour capturer toutes les variables avant toute autre opération
	log.Printf("=== EARLY START - All environment variables ===")
	allEnv := os.Environ()
	for _, env := range allEnv {
		// Afficher toutes les variables qui pourraient être liées à la DB
		if strings.Contains(env, "BDD_") ||
			strings.Contains(env, "DATABASE") ||
			strings.Contains(env, "MYSQL") ||
			strings.Contains(env, "RENDER") {
			log.Printf("ENV: %s", env)
		}
	}
	log.Printf("Total env vars: %d", len(allEnv))

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
		log.Printf("RENDER_SERVICE_ID: %s", os.Getenv("RENDER_SERVICE_ID"))
		log.Printf("RENDER_SERVICE_NAME: %s", os.Getenv("RENDER_SERVICE_NAME"))
		log.Printf("RENDER_SERVICE_TYPE: %s", os.Getenv("RENDER_SERVICE_TYPE"))
	}

	// Afficher toutes les variables d'environnement qui commencent par BDD
	log.Printf("--- All BDD_* environment variables ---")
	for _, env := range os.Environ() {
		if len(env) >= 4 && env[0:4] == "BDD_" {
			parts := strings.SplitN(env, "=", 2)
			if len(parts) == 2 {
				key := parts[0]
				value := parts[1]
				if key == "BDD_PASSWORD" {
					log.Printf("%s: %s", key, map[bool]string{true: "***SET***", false: "NOT SET"}[value != ""])
				} else {
					log.Printf("%s: %s", key, value)
				}
			}
		}
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
		log.Printf("ERROR: Failed to open database connection: %v", err)
		log.Fatal(err)
		return
	}

	defer db.Close()

	log.Printf("Attempting to ping database...")
	if err = db.Ping(); err != nil {
		log.Printf("ERROR: Failed to ping database: %v", err)
		log.Printf("DSN used: %s", conf.FormatDSN())
		log.Fatal(err)
	}

	log.Printf("✓ Database connection successful!")

	store := database.CreateStore(db)
	mux := web.NewHandler(store)

	// Render.com définit la variable PORT, l'utiliser si disponible
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999" // Port par défaut pour le développement local
	}

	log.Printf("Starting server on port %s", port)
	err = http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatal(err)
		return
	}

}
