package entity

import (
	"database/sql"
	"time"
)

type Order struct {
	ID          int          `json:"id" db:"id"`
	User        User         `json:"user" db:"user_id"`
	Restaurant  Restaurant   `json:"restaurant" db:"restaurant_id"`
	Status      string       `json:"status" db:"status"`
	TotalPrice  float64      `json:"total_price" db:"total_price"`
	Number      int          `json:"number" db:"number"`
	CreatedDate time.Time    `json:"created_date" db:"created_date"`
	ClosedDate  sql.NullTime `json:"closed_date" db:"closed_date"`
}

func NewOrder(user User, restaurant Restaurant, status string, totalPrice float64, number int, createdDate time.Time, closedDate sql.NullTime) *Order {
	return &Order{
		User:        user,
		Restaurant:  restaurant,
		Status:      status,
		TotalPrice:  totalPrice,
		Number:      number,
		CreatedDate: createdDate,
		ClosedDate:  closedDate,
	}
}

type OrderStoreInterface interface {
	AddOrder(item Order) (int, error)
	GetAllOrders() ([]Order, error)
	GetOrderByID(id int) *Order
	GetOrderByUserID(id int) []Order
	GetOrderByRestaurantID(id int) []Order
}

//func NewOrder(user User, restaurant Restaurant, status string, totalPrice float64, createdDate time.Time, closedDate time.Time) *Order {
//	return &Order{
//		User:        user,
//		Restaurant:  restaurant,
//		Status:      status,
//		TotalPrice:  totalPrice,
//		CreatedDate: createdDate,
//		ClosedDate:  closedDate,
//	}
//}
//
//type OrderStoreInterface interface {
//	AddOrder(item Order) (int, error)
//	GetAllOrders() ([]Order, error)
//	GetOrderByID(id int) *Order
//	GetOrderByUserID(id int) []Order
//	GetOrderByRestaurantID(id int) []Order
//}
