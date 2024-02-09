package database

import (
	"restaurantHTTP/entity"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	UserStore        entity.UserStoreInterface
	ProductStore     entity.ProductStoreInterface
	ProductTypeStore entity.ProductTypeStoreInterface
}

func CreateStore(db *sqlx.DB) *Store {
	return &Store{
		UserStore:        NewUserStore(db),
		ProductStore:     NewProductStore(db),
		ProductTypeStore: NewProductTypeStore(db),
	}
}
