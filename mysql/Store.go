package database

import (
	"github.com/jmoiron/sqlx"
	"restaurantHTTP/entity"
)

type Store struct {
	UserStore entity.UserStoreInterface
	ProductStore entity.ProductStoreInterface
	ProductType entity.ProductTypeStoreInterface
}

func CreateStore(db *sqlx.DB) *Store {
	return &Store{
		UserStore: NewUserStore(db),
	}
}
