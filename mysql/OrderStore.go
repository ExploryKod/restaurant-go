package database

import (
	"github.com/jmoiron/sqlx"
	"restaurantHTTP/entity"
)

type OrderStore struct {
	*sqlx.DB
}

func NewOrderStore(db *sqlx.DB) *OrderStore {
	return &OrderStore{
		db,
	}
}

func (o OrderStore) AddOrder(item entity.Order) (int, error) {

	res, err := o.DB.Exec("INSERT INTO Orders (user_id, restaurant_id, date, total_price, is_delivered) VALUES ( ? , ? , ?, ?, ?)", item.User.ID, item.Restaurant.ID, item.TotalPrice, item.CreatedDate, item.ClosedDate)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err

	}
	return int(id), nil
}

func (o OrderStore) GetAllOrders() ([]entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderStore) GetOrderByID(id int) *entity.Order {
	//TODO implement me
	panic("implement me")
}

func (o OrderStore) GetOrderByUserID(id int) []entity.Order {
	//TODO implement me
	panic("implement me")
}

func (o OrderStore) GetOrderByRestaurantID(id int) []entity.Order {
	//TODO implement me
	panic("implement me")
}
