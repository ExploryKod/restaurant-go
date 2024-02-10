package database

import (
	"github.com/jmoiron/sqlx"
	"restaurantHTTP/entity"
)

type OrderHasProductStore struct {
	*sqlx.DB
}

func NewOrderHasProductStore(db *sqlx.DB) *OrderHasProductStore {
	return &OrderHasProductStore{
		db,
	}
}

func (o OrderHasProductStore) AddOrderHasProduct(item entity.OrderHasProduct) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderHasProductStore) GetAllOrderHasProducts() ([]entity.OrderHasProduct, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderHasProductStore) GetOrderHasProductByID(id int) *entity.OrderHasProduct {
	//TODO implement me
	panic("implement me")
}

func (o OrderHasProductStore) GetOrderHasProductByOrderID(id int) []entity.OrderHasProduct {
	//TODO implement me
	panic("implement me")
}

func (o OrderHasProductStore) GetOrderHasProductByProductID(id int) []entity.OrderHasProduct {
	//TODO implement me
	panic("implement me")
}
