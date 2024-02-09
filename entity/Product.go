package entity

import (
	"database/sql"
)

type Product struct {
	ID          int         `json:"id" db:"id"`
	ProductType ProductType `json:"product_type" db:"product_type"`
	Restaurant  Restaurant  `json:"restaurant" db:"restaurant"`
	Name        string      `json:"name" db:"name"`
	Price       float64     `json:"price" db:"price"`
	Image       string      `json:"image" db:"image"`
	Description string      `json:"description" db:"description"`
}

type ProductStoreInterface interface {
	GetProductByRestaurantId() ([]Product, error)
	AddProduct(item Product) (int, error)
	DeleteProduct(id Product) error
}
