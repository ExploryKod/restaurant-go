package database

import (
	"database/sql"
	"errors"
	"restaurantHTTP/entity"

	"github.com/jmoiron/sqlx"
)

// ProductStore Product Represents a product with restaurant
type ProductStore struct {
	*sqlx.DB
}

// NewProductStore New Product creates and returns a new Product instance
// Parameters:
// - db: Object of Product
// Returns:
// - A pointer to newly created Prodcut instance
func NewProductStore(db *sqlx.DB) *ProductStore {
	return &ProductStore{
		db,
	}
}

// GetProductByRestaurantId Get Product by restaurant id
// Parameters:
// - resturantId: Id of the restaurant
// Returns:
// - A list of Product
func (t *ProductStore) GetProductByRestaurantId(resturantId string) (*entity.Product, error) {

	product := &entity.Product{}

	err := t.Get(product, "SELECT * FROM Product WHERE restaurantId = ?", resturantId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return product, nil
}

// AddProduct Add Product
// Parameters:
// - item: Object Product
// Returns:
// - Id of inserted Product
func (t *ProductStore) AddProduct(item entity.Product) (int, error) {
	//res, err := t.DB.Exec("INSERT INTO Product ( , ) VALUES ( , )", item, item)
	//if err != nil {
	//	return 0, err
	//}
	//
	//id, err := res.LastInsertId()
	//if err != nil {
	//	return 0, err
	//}

	return int(10), nil
}

// DeleteProduct Delete Product
// Parameters:
// - id: The is of Product
// Returns:
// - nil
func (t *ProductStore) DeleteProduct(id int) error {
	_, err := t.DB.Exec("DELETE FROM Products WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
