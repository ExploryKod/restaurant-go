package database

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type ProductTypeStore struct {
	*sqlx.DB
}

func NewProductTypeStore(db *sqlx.DB) *ProductTypeStore {
	return &ProductTypeStore{
		db,
	}
}
// Get Product type by restaurant id
// Parameters:
// - resturantId: Id of the restaurant
// Returns:
// - A list of product type
func (t *ProductTypeStore) GetProductTypeByRestaurantId(resturantId string) (*entity.ProductType, error) {

	product := &entity.ProductType{}

	err := t.Get(product, "SELECT * FROM Product_type WHERE restaurantId = ?", resturantId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return product, nil
}


func (t *ProductTypeStore) AddProduct(item entity.ProductType) (int, error) {
	res, err := t.DB.Exec("INSERT INTO Product_type ( , ) VALUES ( , )", item, item)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *ProductTypeStore) DeleteProductType(id int) error {
	_, err := t.DB.Exec("DELETE FROM Product_type WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
