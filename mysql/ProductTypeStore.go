package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"restaurantHTTP/entity"

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

// Get Product Type by restaurant id
// Parameters:
// - resturantId: Id of the restaurant
// Returns:
// - A list of product type
func (t *ProductTypeStore) GetProductTypeByRestaurantId(resturantId int) ([]entity.ProductType, error) {
	var productTypeList []entity.ProductType

	query := "SELECT id, name, icon, restaurant_id FROM Product_type WHERE restaurant_id = ?"
	rows, err := t.Query(query, resturantId)
	if err != nil {
		// Log error and return it
		log.Println("Error fetching product types:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productType entity.ProductType
		if err := rows.Scan(&productType.ID, &productType.Name, &productType.Icon, &productType.RestaurantId); err != nil {
			// Log error and return it
			log.Println("Error scanning product type row:", err)
			return nil, err
		}
		productTypeList = append(productTypeList, productType)
	}

	if err := rows.Err(); err != nil {
		// Log error and return it
		log.Println("Error iterating over product type rows:", err)
		return nil, err
	}
	return productTypeList, nil
}

func (t *ProductTypeStore) GetProductTypeById(id int) (*entity.ProductType, error) {

	product := &entity.ProductType{}

	err := t.Get(product, "SELECT * FROM Product_type WHERE id = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return product, nil
}

// Add Product Type
// Parameters:
// - item: Object Product Type
// Returns:
// - Id of inserted Product Type
func (t *ProductTypeStore) AddProductType(item entity.ProductType) (int, error) {
	fmt.Println(item, item)
	res, err := t.DB.Exec("INSERT INTO Product_type (name, icon, restaurant_id ) VALUES (?, ?, ?)", item.Name, item.Icon, item.RestaurantId)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Delete Product
// Parameters:
// - id: The is of Product Type
// Returns:
// - nil
func (t *ProductTypeStore) DeleteProductType(id int) error {
	_, err := t.DB.Exec("DELETE FROM Product_type WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
