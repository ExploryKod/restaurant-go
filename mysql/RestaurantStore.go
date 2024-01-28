package database

import (
	"github.com/jmoiron/sqlx"
	"restaurantHTTP/entity"
)

type RestaurantStore struct {
	*sqlx.DB
}

func NewRestaurantStore(db *sqlx.DB) *RestaurantStore {
	return &RestaurantStore{
		db,
	}
}

func (t *RestaurantStore) AddRestaurant(item entity.Restaurant) (int, error) {
	res, err := t.DB.Exec("INSERT INTO Restaurants (name, logo, mail, is_validated) VALUES ( ? , ? , ?, ?)", item, item, item, item)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *RestaurantStore) GetRestaurants() ([]entity.Restaurant, error) {
	var restaurantList []entity.Restaurant

	rows, err := t.Query("SELECT id, name, description FROM Restaurants")
	if err != nil {
		return []entity.Restaurant{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var restaurant entity.Restaurant
		if err = rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Mail, &restaurant.IsValidated); err != nil {
			return []entity.Restaurant{}, err
		}
		restaurantList = append(restaurantList, restaurant)
	}

	if err = rows.Err(); err != nil {
		return []entity.Restaurant{}, err
	}

	return restaurantList, nil
}
