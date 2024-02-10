package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
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

func (o OrderHasProductStore) GetOrderHasProductByOrderID(id int) *entity.OrderHasProduct {
	orderHasProduct := &entity.OrderHasProduct{}

	query := `
			SELECT id, name, price, image, description FROM Products p JOIN Order_has_products op ON p.id = op.product_id WHERE op.order_id = ?
	`

	err := o.Select(&orderHasProduct.Products, query, id)
	if err != nil {
		log.Println("orderhasproduct", err)
		return nil
	}
	fmt.Printf("%#v\n", orderHasProduct.Products)

	return orderHasProduct
}

func (o OrderHasProductStore) GetOrderHasProductByProductID(id int) *entity.OrderHasProduct {
	//TODO implement me
	panic("implement me")
}
