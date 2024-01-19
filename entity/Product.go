package entity

type Product struct {
	ID          int         `json:"id"`
	ProductType ProductType `json:"product_type"`
	Restaurant  Restaurant  `json:"restaurant"`
	Name        string      `json:"name"`
	Price       float64     `json:"price"`
	Image       string      `json:"image"`
	Description string      `json:"description"`
}
