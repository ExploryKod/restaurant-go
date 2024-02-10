package entity

type Product struct {
	ID          int         `json:"id"`
	ProductType ProductType `json:"-"`
	Restaurant  Restaurant  `json:"-"`
	Name        string      `json:"name"`
	Price       int         `json:"price"`
	Image       string      `json:"image"`
	Description string      `json:"description"`
}

type ProductStoreInterface interface {
	GetProductByRestaurantId(restaurantId string) (*Product, error)
	AddProduct(item Product) (int, error)
	DeleteProduct(id int) error
}
