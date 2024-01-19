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
