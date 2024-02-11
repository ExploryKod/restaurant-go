package entity

type ProductType struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Icon         string `json:"icon" db:"icon"`
	RestaurantId int    `json:"restaurant_id" db:"restaurant_id"`
}

type ProductTypeStoreInterface interface {
	GetProductTypeByRestaurantId(resturantId int) ([]ProductType, error)
	AddProductType(item ProductType) (int, error)
	DeleteProductType(id int) error
	GetProductTypeById(id int) (*ProductType, error)
}
