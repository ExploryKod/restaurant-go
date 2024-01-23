package database

import (
	"github.com/jmoiron/sqlx"
	"restaurantHTTP/entity"
)

type Store struct {
	UserStore entity.UserStoreInterface
}

func CreateStore(db *sqlx.DB) *Store {
	return &Store{
		UserStore: NewUserStore(db),
	}
}
