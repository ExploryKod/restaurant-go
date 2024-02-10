package entity

type OrderHasProduct struct {
	Order    Order     `json:"order" db:"order"`
	Products []Product `json:"products" db:"products"`
}

func NewOrderHasProduct(order Order, products []Product) *OrderHasProduct {
	return &OrderHasProduct{
		Order:    order,
		Products: products,
	}
}

type OrderHasProductStoreInterface interface {
	AddOrderHasProduct(item *OrderHasProduct) (int, error)
	GetAllOrderHasProducts() ([]OrderHasProduct, error)
	GetOrderHasProductByID(id int) *OrderHasProduct
	GetOrderHasProductByOrderID(id int) []OrderHasProduct
	GetOrderHasProductByProductID(id int) []OrderHasProduct
}
