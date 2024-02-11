package database

import (
	"restaurantHTTP/entity"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	UserStore            entity.UserStoreInterface
	RestaurantStore      entity.RestaurantStoreInterface
	OrderStore           entity.OrderStoreInterface
	OrderHasProductStore entity.OrderHasProductStoreInterface
	ProductStore         entity.ProductStoreInterface
	ProductTypeStore     entity.ProductTypeStoreInterface
	RestaurantUserStore  entity.RestaurantUserStoreInterface
}

func CreateStore(db *sqlx.DB) *Store {
	return &Store{
		UserStore:            NewUserStore(db),
		RestaurantStore:      NewRestaurantStore(db),
		OrderStore:           NewOrderStore(db),
		OrderHasProductStore: NewOrderHasProductStore(db),
		ProductStore:         NewProductStore(db),
		ProductTypeStore:     NewProductTypeStore(db),
		RestaurantUserStore:  NewRestaurantUserStore(db),
	}
}
