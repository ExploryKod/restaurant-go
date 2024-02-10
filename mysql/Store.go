package database

import (
	"restaurantHTTP/entity"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	UserStore        entity.UserStoreInterface
	RestaurantStore  entity.RestaurantStoreInterface
	ProductStore     entity.ProductStoreInterface
	ProductTypeStore entity.ProductTypeStoreInterface
}

func CreateStore(db *sqlx.DB) *Store {
	return &Store{
		UserStore:        NewUserStore(db),
		RestaurantStore:  NewRestaurantStore(db),
		ProductStore:     NewProductStore(db),
		ProductTypeStore: NewProductTypeStore(db),
	}
}
