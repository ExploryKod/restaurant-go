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
	res, err := t.DB.Exec("INSERT INTO Restaurants (name) VALUES (?)", item.Name)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *RestaurantStore) GetRestaurantByID(id int) *entity.Restaurant {
	restaurant := &entity.Restaurant{}
	err := s.Get(restaurant, "SELECT id, name, logo, mail, is_validated FROM Restaurants WHERE id = ?", id)
	if err != nil {
		return nil
	}
	return restaurant
}

func (s *RestaurantStore) GetAllRestaurants() ([]entity.Restaurant, error) {
	var restaurantList []entity.Restaurant

	rows, err := s.Query("SELECT id, name, logo, mail, is_validated  FROM Restaurants")
	if err != nil {
		return []entity.Restaurant{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var restaurant entity.Restaurant
		if err = rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Logo, &restaurant.Mail, &restaurant.IsValidated); err != nil {
			return []entity.Restaurant{}, err
		}
		restaurantList = append(restaurantList, restaurant)
	}

	if err = rows.Err(); err != nil {
		return []entity.Restaurant{}, err
	}
	return restaurantList, nil
}
