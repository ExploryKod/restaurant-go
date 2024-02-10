package database

import (
	"database/sql"
	"errors"
	"restaurantHTTP/entity"

	"github.com/jmoiron/sqlx"
)

// Product Represents a product with restaurant
type ProductStore struct {
	*sqlx.DB
}

// New Product creates and returns a new Product instance
// Parameters:
// - db: Object of Product
// Returns:
// - A pointer to newly created Prodcut instance
func NewProductStore(db *sqlx.DB) *ProductStore {
	return &ProductStore{
		db,
	}
}

// Get Product by restaurant id
// Parameters:
// - resturantId: Id of the restaurant
// Returns:
// - A list of Product
func (t *ProductStore) GetProductByRestaurantId(resturantId string) (*entity.Product, error) {

	product := &entity.Product{}

	err := t.Get(product, "SELECT * FROM Products WHERE restaurant_id = ?", resturantId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return product, nil
}

// Add Product
// Parameters:
// - item: Object Product
// Returns:
// - Id of inserted Product
func (t *ProductStore) AddProduct(item entity.Product) (int, error) {

	res, err := t.DB.Exec("INSERT INTO Products (product_type_id, restaurant_id, name, price, image, description) VALUES (?,?,?,?,?,?)", item.ProductType.ID, item.Restaurant.ID, item.Name, item.Price, item.Image, item.Description)
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

// Get List of allergies
// Parameters:
// Returns:
// - Allergens
func (t *ProductStore) GetAllergiesList() ([]entity.Allergen, error) {
	var allergiesList []entity.Allergen

	rows, err := t.Query("SELECT id, name FROM Allergens")
	if err != nil {
		return []entity.Allergen{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var allergy entity.Allergen
		if err = rows.Scan(&allergy.ID, &allergy.Name); err != nil {
			return []entity.Allergen{}, err
		}
		allergiesList = append(allergiesList, allergy)
	}

	if err = rows.Err(); err != nil {
		return []entity.Allergen{}, err
	}
	return allergiesList, nil
}
