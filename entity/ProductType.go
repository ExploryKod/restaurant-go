package entity

type ProductType struct {
	ID         int        `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	Icon       string     `json:"icon" db:"icon"`
	Restaurant Restaurant `json:"restaurant" db:"restaurant_id"`
}
