package database

import (
	"fmt"
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
func (t *ProductStore) GetProductByRestaurantId(resturantId int) ([]entity.Product, error) {
	//
	var productList []entity.Product

	query := `WITH ranked_products AS (
		SELECT 
			p.id,
			p.name,
			p.price,
			p.image,
			p.description,
			r.name AS restaurant_name,
			r.logo AS restaurant_logo,
			r.image AS restaurant_image,
			r.phone AS restaurant_phone,
			r.mail AS restaurant_mail,
			r.is_open AS restaurant_is_open,
			r.grade AS restaurant_grade,
			pt.name AS productType_name,
			pt.icon AS productType_icon,
			ROW_NUMBER() OVER (PARTITION BY p.name ORDER BY p.id) AS row_num
		FROM 
			Products p 
		JOIN 
			Restaurants r ON p.restaurant_id = r.id
		JOIN 
			Product_type pt ON p.restaurant_id = pt.restaurant_id
		WHERE 
			p.restaurant_id = ?
	)
	SELECT 
		id,
		name,
		price,
		image,
		description,
		restaurant_name,
		restaurant_logo,
		restaurant_image,
		restaurant_phone,
		restaurant_mail,
		restaurant_is_open,
		restaurant_grade,
		productType_name,
		productType_icon
	FROM 
		ranked_products
	WHERE 
		row_num = 1;
	`
	rows, err := t.Query(query, resturantId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		if err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Image, &product.Description,
			&product.Restaurant.Name, &product.Restaurant.Logo, &product.Restaurant.Image, &product.Restaurant.Phone,
			&product.Restaurant.Mail, &product.Restaurant.IsOpen,
			&product.Restaurant.Grade, &product.ProductType.Name, &product.ProductType.Icon); err != nil {
			fmt.Println(err)
			return nil, err
		}
		productList = append(productList, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return productList, nil
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
