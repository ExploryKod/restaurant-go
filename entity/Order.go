package entity

import "time"

type Order struct {
	ID          int        `json:"id"`
	User        User       `json:"user"`
	Restaurant  Restaurant `json:"restaurant"`
	Status      string     `json:"status"`
	TotalPrice  float64    `json:"total_price"`
	CreatedDate time.Time  `json:"created_date"`
	ClosedDate  time.Time  `json:"closed_date"`
}

func NewOrder(user User, restaurant Restaurant, status string, totalPrice float64, createdDate time.Time, closedDate time.Time) *Order {
	return &Order{
		User:        user,
		Restaurant:  restaurant,
		Status:      status,
		TotalPrice:  totalPrice,
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
