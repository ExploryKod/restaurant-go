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

func (o OrderHasProductStore) AddOrderHasProduct(item *entity.OrderHasProduct) (int, error) {
	for _, product := range item.Products {
		_, err := o.Exec("INSERT INTO Order_has_products (order_id, product_id) VALUES ( ? , ? )", item.Order.ID, product.ID)
		if err != nil {
			return 0, err
		}
	}
	return 0, nil
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
