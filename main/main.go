package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	database "restaurantHTTP/postgressql"
	"restaurantHTTP/web"
)

func main() {
	connectionStr := "user=postgres password=password dbname=postgres port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	store := database.CreateStore(db)
	mux := web.NewHandler(store)

	err = http.ListenAndServe(":8097", mux)
	if err != nil {
		_ = fmt.Errorf("impossible de lancer le serveur : %w", err)
		return
	}
}
