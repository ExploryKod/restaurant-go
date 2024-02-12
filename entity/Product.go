package entity

type Product struct {
	ID          int         `json:"id" db:"id"`
	ProductType ProductType `json:"product_type" db:"product_type_id"`
	Restaurant  Restaurant  `json:"restaurant" db:"restaurant_id"`
	Name        string      `json:"name" db:"name"`
	Price       float64     `json:"price" db:"price"`
	Image       string      `json:"image" db:"image"`
	Description string      `json:"description" db:"description"`
}

type ProductStoreInterface interface {
	GetProductByRestaurantId(restaurantId int) ([]Product, error)
	AddProduct(item Product) (int, error)
	DeleteProduct(id int) (bool, error)
}
