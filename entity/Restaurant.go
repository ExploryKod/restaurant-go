package entity

import "time"

type Restaurant struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Logo        string    `json:"logo" db:"logo"`
	Image       string    `json:"image" db:"image"`
	Phone       string    `json:"phone" db:"phone"`
	Mail        string    `json:"mail" db:"mail"`
	IsOpen      bool      `json:"is_open" db:"is_open"`
	OpeningTime time.Time `json:"opening_time" db:"opening_time"`
	ClosingTime time.Time `json:"closing_time" db:"closing_time"`
	Grade       int       `json:"grade" db:"grade"`
	IsValidated bool      `json:"is_validate" db:"is_validated"`
}

type RestaurantStoreInterface interface {
	AddRestaurant(item Restaurant) (int, error)
	GetRestaurant() ([]Restaurant, error)
}
