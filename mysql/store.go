package database

import (
	"database/sql"
)

func CreateStore(db *sql.DB) *Store {
	return &Store{
		NewUserStore(db),
	}
}

type Store struct {
	restaurantHTTP.UserStoreInterface
}
