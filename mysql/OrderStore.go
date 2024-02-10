package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
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

func (o OrderStore) AddOrder(order entity.Order) (int, error) {

	var currentOrderNumber int
	currentDate := order.CreatedDate.Format("2006-01-02")

	err := o.QueryRow("SELECT MAX(number) FROM Orders WHERE restaurant_id = ? AND DATE(created_date) = ? AND (SELECT COUNT(id) FROM Orders WHERE restaurant_id = ? AND DATE(created_date) = ?) > 0", order.Restaurant.ID, currentDate, order.Restaurant.ID, currentDate).Scan(&currentOrderNumber)
	if err != nil {
		currentOrderNumber = 0
		log.Println(err)
	}
	nextOrderNumber := currentOrderNumber + 1

	res, err := o.DB.Exec("INSERT INTO Orders (user_id, restaurant_id, status, total_price, number, created_date, closed_date) VALUES ( ? , ? , ?, ?, ?, ?, ?)", order.User.ID, order.Restaurant.ID, order.Status, order.TotalPrice, nextOrderNumber, order.CreatedDate, order.ClosedDate)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println()
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
