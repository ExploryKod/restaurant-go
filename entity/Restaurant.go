package entity

import "time"

type Restaurant struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Logo        string    `json:"logo"`
	Image       string    `json:"image"`
	Phone       string    `json:"phone"`
	Mail        string    `json:"mail"`
	IsOpen      bool      `json:"is_open"`
	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`
	Grade       int       `json:"grade"`
	IsValidated bool      `json:"is_validate"`
}
