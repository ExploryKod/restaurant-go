package entity

type ProductType struct {
	ID         int        `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	Icon       string     `json:"icon" db:"icon"`
	Restaurant Restaurant `json:"restaurant" db:"restaurant_id"`
}

type ProductTypeStoreInterface interface {
	GetProductTypeByRestaurantId(restaurantId string) (*ProductType, error)
	AddProduct(item ProductType) (int, error)
	DeleteProductType(id int) error
}
