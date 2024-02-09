package database

import (
	"github.com/jmoiron/sqlx"
	"restaurantHTTP/entity"
)

type Store struct {
	UserStore            entity.UserStoreInterface
	RestaurantStore      entity.RestaurantStoreInterface
	OrderStore           entity.OrderStoreInterface
	OrderHasProductStore entity.OrderHasProductStoreInterface
}

func CreateStore(db *sqlx.DB) *Store {
	return &Store{
		UserStore:            NewUserStore(db),
		RestaurantStore:      NewRestaurantStore(db),
		OrderStore:           NewOrderStore(db),
		OrderHasProductStore: NewOrderHasProductStore(db),
	}
}
