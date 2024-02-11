package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"restaurantHTTP/entity"
)

type OrderHasProductStore struct {
	*sqlx.DB
}

func NewOrderHasProductStore(db *sqlx.DB) *OrderHasProductStore {
	return &OrderHasProductStore{
		db,
	}
}

func (o OrderHasProductStore) AddOrderHasProduct(item *entity.OrderHasProduct) (int, error) {
	for _, product := range item.Products {
		_, err := o.Exec("INSERT INTO Order_has_products (order_id, product_id) VALUES ( ? , ? )", item.Order.ID, product.ID)
		if err != nil {
			return 0, err
		}
	}
	return 0, nil
}

func (o OrderHasProductStore) GetAllOrderHasProducts() ([]entity.OrderHasProduct, error) {
	//var orderHasProductList []entity.OrderHasProduct
	var orderMap = make(map[int]*entity.OrderHasProduct)

	query := `
		SELECT
			o.id,
			o.status,
			o.total_price,
			o.number,	
			o.created_date,
			o.closed_date,
			o.user_id,
			u.username AS user_username,
			u.password AS user_password,
			u.name AS user_name,
			u.firstname AS user_firstname,
			u.mail AS user_mail,
			u.phone AS user_phone,
			u.is_superadmin AS user_is_superadmin,
			r.id AS restaurant_id,
			r.name AS restaurant_name,
			r.logo AS restaurant_logo,
			r.mail AS restaurant_mail,
			r.is_validated AS restaurant_is_validated,
			p.id AS product_id,
			p.name AS product_name,
			p.price AS product_price,
			p.image AS product_image,
			p.description AS product_description
		FROM Orders o
		JOIN Users u ON o.user_id = u.id
		JOIN Restaurants r ON o.restaurant_id = r.id
		JOIN Order_has_products op ON o.id = op.order_id
		JOIN Products p ON p.id = op.product_id
	`

	rows, err := o.Queryx(query)
	if err != nil {
		return []entity.OrderHasProduct{}, err
	}

	defer rows.Close()

	/*for rows.Next() {
		var orderHasProduct entity.OrderHasProduct
		var user entity.User
		var restaurant entity.Restaurant
		var product entity.Product
		if err = rows.Scan(
			&orderHasProduct.Order.ID,
			&orderHasProduct.Order.Status,
			&orderHasProduct.Order.TotalPrice,
			&orderHasProduct.Order.Number,
			&orderHasProduct.Order.CreatedDate,
			&orderHasProduct.Order.ClosedDate,
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Name,
			&user.Firstname,
			&user.Mail,
			&user.Phone,
			&user.IsSuperadmin,
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Logo,
			&restaurant.Mail,
			&restaurant.IsValidated,
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Image,
			&product.Description,
		); err != nil {
			return []entity.OrderHasProduct{}, err
		}
		orderHasProduct.Order.User = user
		orderHasProduct.Order.Restaurant = restaurant
		orderHasProduct.Products = append(orderHasProduct.Products, product)
		orderHasProductList = append(orderHasProductList, orderHasProduct)
	}

	if err = rows.Err(); err != nil {
		return []entity.OrderHasProduct{}, err
	}*/

	for rows.Next() {
		var order entity.OrderHasProduct
		var user entity.User
		var restaurant entity.Restaurant
		var product entity.Product
		err := rows.Scan(
			&order.Order.ID,
			&order.Order.Status,
			&order.Order.TotalPrice,
			&order.Order.Number,
			&order.Order.CreatedDate,
			&order.Order.ClosedDate,
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Name,
			&user.Firstname,
			&user.Mail,
			&user.Phone,
			&user.IsSuperadmin,
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Logo,
			&restaurant.Mail,
			&restaurant.IsValidated,
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Image,
			&product.Description,
		)
		if err != nil {
			return nil, err
		}

		// Si la commande n'est pas déjà dans la map, l'ajouter
		if _, ok := orderMap[order.Order.ID]; !ok {
			order.Order.User = user
			order.Order.Restaurant = restaurant
			orderMap[order.Order.ID] = &order
		}

		// Ajouter le produit à la commande
		orderMap[order.Order.ID].Products = append(orderMap[order.Order.ID].Products, product)
	}

	// Convertir la map en slice
	var orderHasProductList []entity.OrderHasProduct
	for _, order := range orderMap {
		orderHasProductList = append(orderHasProductList, *order)
	}

	return orderHasProductList, nil
}

func (o OrderHasProductStore) GetOrderHasProductByOrderID(id int) *entity.OrderHasProduct {
	orderHasProduct := &entity.OrderHasProduct{}

	query := `
			SELECT id, name, price, image, description FROM Products p JOIN Order_has_products op ON p.id = op.product_id WHERE op.order_id = ?
	`

	err := o.Select(&orderHasProduct.Products, query, id)
	if err != nil {
		log.Println("orderhasproduct", err)
		return nil
	}
	fmt.Printf("%#v\n", orderHasProduct.Products)

	return orderHasProduct
}

func (o OrderHasProductStore) GetOrderHasProductByProductID(id int) *entity.OrderHasProduct {
	//TODO implement me
	panic("implement me")
}
