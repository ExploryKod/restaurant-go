package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	database "restaurantHTTP/postgressql"
	"restaurantHTTP/web"
)

func main() {
	//conf := mysql.Config{
	//	User:                 "root",
	//	Passwd:               "password",
	//	Addr:                 "database:3306",
	//	DBName:               "restaurantbdd",
	//	Net:                  "tcp",
	//	AllowNativePasswords: true,
	//}

	conf := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "database:3306",
		DBName:               "restaurantbdd",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
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

	err = http.ListenAndServe(":8097", mux)
	if err != nil {
		_ = fmt.Errorf("impossible de lancer le serveur : %w", err)
		return
	}
}
