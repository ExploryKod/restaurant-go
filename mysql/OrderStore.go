package database

import (
	"fmt"
	"log"
	"restaurantHTTP/entity"

	"github.com/jmoiron/sqlx"
)

type OrderStore struct {
	*sqlx.DB
}

func NewOrderStore(db *sqlx.DB) *OrderStore {
	return &OrderStore{
		db,
	}
}

func (o *OrderStore) AddOrder(order entity.Order) (int, error) {

	var currentOrderNumber int
	currentDate := order.CreatedDate.Format("2006-01-02")

	err := o.QueryRow("SELECT MAX(number) FROM Orders WHERE restaurant_id = ? AND DATE(created_date) = ? AND (SELECT COUNT(id) FROM Orders WHERE restaurant_id = ? AND DATE(created_date) = ?) > 0", order.Restaurant.ID, currentDate, order.Restaurant.ID, currentDate).Scan(&currentOrderNumber)
	if err != nil {
		currentOrderNumber = 0
		log.Println(err)
	}
	nextOrderNumber := currentOrderNumber + 1

	res, err := o.DB.Exec("INSERT INTO Orders (user_id, restaurant_id, status, total_price, number, created_date, closed_date) VALUES ( ? , ? , ?, ?, ?, ?, ?)", order.User.ID, order.Restaurant.ID, order.Status, order.TotalPrice, nextOrderNumber, order.CreatedDate, order.ClosedDate)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println()
	return int(id), nil
}

func (o *OrderStore) GetAllOrders() ([]entity.Order, error) {
	var orderList []entity.Order

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
			u.birthday AS user_birthday,
			o.restaurant_id, 
			r.name AS restaurant_name,
			r.logo AS restaurant_logo,
			r.image AS restaurant_image,
			r.phone AS restaurant_phone,
			r.mail AS restaurant_mail,
			r.is_open AS restaurant_is_open,
			r.opening_time AS restaurant_opening_time,
			r.closing_time AS restaurant_closing_time,
			r.grade AS restaurant_grade,
			r.is_validated AS restaurant_is_validated
		FROM
			Orders o 
		JOIN 
			Users u ON o.user_id = u.id 
		JOIN 
			Restaurants r ON o.restaurant_id = r.id`

	rows, err := o.Queryx(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		order := entity.Order{}
		err = rows.Scan(&order.ID, &order.Status, &order.TotalPrice, &order.Number, &order.CreatedDate, &order.ClosedDate, &order.User.ID, &order.User.Username, &order.User.Password, &order.User.Name, &order.User.Firstname, &order.User.Mail, &order.User.Phone, &order.User.IsSuperadmin, &order.User.Birthday, &order.Restaurant.ID, &order.Restaurant.Name, &order.Restaurant.Logo, &order.Restaurant.Image, &order.Restaurant.Phone, &order.Restaurant.Mail, &order.Restaurant.IsOpen, &order.Restaurant.OpeningTime, &order.Restaurant.ClosingTime, &order.Restaurant.Grade, &order.Restaurant.IsValidated)
		if err != nil {
			log.Fatalln(err)
		}
		orderList = append(orderList, order)
	}
	return orderList, nil
}

func (o *OrderStore) GetOrderByID(id int) *entity.Order {
	order := &entity.Order{}

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
            u.birthday AS user_birthday,
            o.restaurant_id, 
            r.name AS restaurant_name,
            r.logo AS restaurant_logo,
            r.image AS restaurant_image,
            r.phone AS restaurant_phone,
			r.mail AS restaurant_mail,
			r.is_open AS restaurant_is_open,
			r.opening_time AS restaurant_opening_time,
			r.closing_time AS restaurant_closing_time,
			r.grade AS restaurant_grade,
			r.is_validated AS restaurant_is_validated
        FROM
            Orders o 
        JOIN 
            Users u ON o.user_id = u.id 
        JOIN 
            Restaurants r ON o.restaurant_id = r.id 
        WHERE 
            o.id = ?`

	rows, err := o.Queryx(query, id)
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		err = rows.Scan(&order.ID, &order.Status, &order.TotalPrice, &order.Number, &order.CreatedDate, &order.ClosedDate, &order.User.ID, &order.User.Username, &order.User.Password, &order.User.Name, &order.User.Firstname, &order.User.Mail, &order.User.Phone, &order.User.IsSuperadmin, &order.User.Birthday, &order.Restaurant.ID, &order.Restaurant.Name, &order.Restaurant.Logo, &order.Restaurant.Image, &order.Restaurant.Phone, &order.Restaurant.Mail, &order.Restaurant.IsOpen, &order.Restaurant.OpeningTime, &order.Restaurant.ClosingTime, &order.Restaurant.Grade, &order.Restaurant.IsValidated)
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Printf("%#v\n", order)
	}
	return order
}

/*func (o *OrderStore) GetOrderProductsByOrderID(id int) *entity.OrderHasProduct {
	orderHasProduct := &entity.OrderHasProduct{}

	query := `
			SELECT id, name, price, image, description FROM Products p JOIN Order_has_products op ON p.id = op.product_id WHERE op.order_id = ?
	`

	err := o.Select(&orderHasProduct.Products, query, id)
	if err != nil {
		log.Println("order", err)
		return nil
	}
	fmt.Printf("%#v\n", orderHasProduct.Products)

	return orderHasProduct
}*/

func (o *OrderStore) GetOrderByUserID(id int) []entity.Order {
	//TODO implement me
	panic("implement me")
}
func (o *OrderStore) ValidateOrder(id int) (any, error) {
	query := `UPDATE Orders
	SET status = "preparation"
	WHERE id = ?`
	rows, err := o.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return true, nil
}
func (o *OrderStore) ReadyOrder(id int) (bool, error) {
	query := `UPDATE Orders
	SET status = "ready"
	WHERE id = ?`
	rows, err := o.Queryx(query, id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}

func (o *OrderStore) CompleteOrder(id int) (bool, error) {
	query := `UPDATE Orders
	SET closed_date = CURRENT_DATE,
	status = "delivered"
	WHERE id = ?`
	rows, err := o.Queryx(query, id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}

func (o *OrderStore) GetOrderByRestaurantIDIncoming(id int) ([]entity.Order, error) {
	var orderList []entity.Order
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
	u.birthday AS user_birthday,
	o.restaurant_id, 
	r.name AS restaurant_name,
	r.logo AS restaurant_logo,
	r.image AS restaurant_image,
	r.phone AS restaurant_phone,
	r.mail AS restaurant_mail,
	r.is_open AS restaurant_is_open,
	r.grade AS restaurant_grade,
	r.is_validated AS restaurant_is_validated
    FROM
        Orders o 
    JOIN 
        Users u ON o.user_id = u.id 
    JOIN 
        Restaurants r ON o.restaurant_id = r.id 
    WHERE o.restaurant_id = ?
	AND o.closed_date IS NULL`

	rows, err := o.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order entity.Order
		if err = rows.Scan(&order.ID, &order.Status, &order.TotalPrice, &order.Number, &order.CreatedDate, &order.ClosedDate, &order.User.ID, &order.User.Username, &order.User.Password, &order.User.Name, &order.User.Firstname, &order.User.Mail, &order.User.Phone, &order.User.IsSuperadmin, &order.User.Birthday, &order.Restaurant.ID, &order.Restaurant.Name, &order.Restaurant.Logo, &order.Restaurant.Image, &order.Restaurant.Phone, &order.Restaurant.Mail, &order.Restaurant.IsOpen, &order.Restaurant.Grade, &order.Restaurant.IsValidated); err != nil {
			return nil, err
		}
		orderList = append(orderList, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orderList, nil
}
