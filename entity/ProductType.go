package entity

type ProductType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type ProductTypeStoreInterface interface {
	GetProductTypeByRestaurantId(resturantId string) (*ProductType, error)
	AddProduct(item ProductType) (int, error)
	DeleteProductType(id int) error
}
