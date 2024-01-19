package database

import (
	"database/sql"
)

type Store struct {
	UserStore *UserStore
}

func CreateStore(db *sql.DB) *Store {
	return &Store{
		UserStore: NewUserStore(db),
	}
}
