package main

import (
	"database/sql"
	database "demoHTTP/mysql"
	"demoHTTP/web"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	conf := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "todo",
		AllowNativePasswords: true, // Il faut le préciser
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
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