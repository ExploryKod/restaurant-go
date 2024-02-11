package database

import (
	"github.com/jmoiron/sqlx"
	"restaurantHTTP/entity"
)

type RestaurantUserStore struct {
	*sqlx.DB
}

func NewRestaurantUserStore(db *sqlx.DB) *ProductStore {
	return &ProductStore{
		db,
	}
}
func (t *ProductStore) AddRestaurantUser(item entity.RestaurantHasUsers) (int, error) {

	res, err := t.DB.Exec("INSERT INTO Restaurant_has_users (restaurant_id, user_id, is_admin, role) VALUES (?,?,?,?)", item.Restaurant.ID, item.User.ID, item.IsAdmin, item.Role)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *ProductStore) UpdateRestaurantUser(item entity.RestaurantHasUsers) error {
	_, err := t.DB.Exec("UPDATE Restaurant_has_users SET is_admin = ?, role = ? WHERE restaurant_id = ? AND user_id = ?", item.IsAdmin, item.Role, item.Restaurant.ID, item.User.ID)
	if err != nil {
		return err
	}
	return nil
}

// fun delete
func (t *ProductStore) DeleteRestaurantUser(userId int) error {
	_, err := t.DB.Exec("DELETE FROM Restaurant_has_users WHERE user_id = ?", userId)
	if err != nil {
		return err
	}
	return nil
}

// func GetRestaurantUsers only with restaurant id dont use user
func (t *ProductStore) GetRestaurantUsers(restaurantId int) ([]entity.RestaurantHasUsers, error) {
	var restaurantHasUsers []entity.RestaurantHasUsers

	rows, err := t.Query("SELECT r.*, u.username, u.mail FROM Restaurant_has_users r JOIN Users u ON r.user_id = u.id WHERE restaurant_id = ?", restaurantId)
	if err != nil {
		return []entity.RestaurantHasUsers{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var restaurantHasUser entity.RestaurantHasUsers
		if err = rows.Scan(&restaurantHasUser.Restaurant.ID, &restaurantHasUser.User.ID, &restaurantHasUser.IsAdmin, &restaurantHasUser.Role, &restaurantHasUser.User.Username, &restaurantHasUser.User.Mail); err != nil {
			return []entity.RestaurantHasUsers{}, err
		}
		restaurantHasUsers = append(restaurantHasUsers, restaurantHasUser)
	}

	if err = rows.Err(); err != nil {
		return []entity.RestaurantHasUsers{}, err
	}

	return restaurantHasUsers, nil
}
