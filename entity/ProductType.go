package entity

type ProductType struct {
	ID   int    `json:"id" db:"id"``
	Name string `json:"name" db:"name"``
	Icon string `json:"icon" db:"icon"``
	RestaurantId int `json:"restaurantId" db:"restaurantId"``
}
