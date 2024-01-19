package database

import (
	"database/sql"
	"demoHTTP"
)

func CreateStore(db *sql.DB) *Store {
	return &Store{
		NewUserStore(db),
	}
}

type Store struct {
	demoHTTP.UserStoreInterface
}
