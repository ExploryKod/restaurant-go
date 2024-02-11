package database

import (
	"github.com/jmoiron/sqlx"
	"restaurantHTTP/entity"
	"time"
)

type RestaurantStore struct {
	*sqlx.DB
}

func NewRestaurantStore(db *sqlx.DB) *RestaurantStore {
	return &RestaurantStore{
		db,
	}
}

func (s *RestaurantStore) AddRestaurant(item entity.Restaurant) (int, error) {
	res, err := s.DB.Exec("INSERT INTO Restaurants (name, logo, image, phone, mail, is_open, opening_time, closing_time, grade, is_validated) VALUES ( ? , ? , ?, ?, ?, ?, ?, ?, ?, ?)",
		item.Name, item.Logo, item.Image, item.Phone, item.Mail, item.IsOpen, item.OpeningTime, item.ClosingTime, item.Grade, item.IsValidated)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *RestaurantStore) AddTagToRestaurant(item entity.Tag) (int, error) {
	res, err := s.DB.Exec("INSERT INTO Tags (name) VALUES ( ? )", item.Name)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *RestaurantStore) AddTagToRestaurantHasTag(tagID int, restaurantID int) (int, error) {
	res, err := s.DB.Exec("INSERT INTO RestaurantHasTags (restaurant_id, tag_id) VALUES ( ? , ?)", tagID, restaurantID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *RestaurantStore) GetAllRestaurants() ([]entity.Restaurant, error) {
	var restaurantList []entity.Restaurant

	rows, err := s.Query("SELECT id, name, logo, image, phone, mail, is_open, opening_time, closing_time, grade, is_validated  FROM Restaurants")
	if err != nil {
		return []entity.Restaurant{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var restaurant entity.Restaurant
		var openingTimeStr, closingTimeStr string // Temporary variables for string representation

		if err = rows.Scan(&restaurant.ID,
			&restaurant.Name,
			&restaurant.Logo,
			&restaurant.Image,
			&restaurant.Phone,
			&restaurant.Mail,
			&restaurant.IsOpen,
			&openingTimeStr,
			&closingTimeStr,
			&restaurant.Grade,
			&restaurant.IsValidated); err != nil {
			return []entity.Restaurant{}, err
		}

		// Convert opening_time and closing_time strings to time.Time
		restaurant.OpeningTime, err = time.Parse("15:04:05", openingTimeStr)
		if err != nil {
			return []entity.Restaurant{}, err
		}

		restaurant.ClosingTime, err = time.Parse("15:04:05", closingTimeStr)
		if err != nil {
			return []entity.Restaurant{}, err
		}

		restaurantList = append(restaurantList, restaurant)
	}

	if err = rows.Err(); err != nil {
		return []entity.Restaurant{}, err
	}
	return restaurantList, nil
}

func (s *RestaurantStore) GetRestaurantByID(id int) *entity.Restaurant {
	restaurant := &entity.Restaurant{}
	err := s.Get(restaurant, "SELECT name, logo, image, phone, mail, is_open, opening_time, closing_time, grade, is_validated FROM Restaurants WHERE id = ?", id)
	if err != nil {
		return nil
	}
	return restaurant
}

func (s *RestaurantStore) UpdateRestaurant(item entity.Restaurant, restaurantID int) error {

	_, err := s.DB.Exec("UPDATE Restaurant SET name = ?, logo = ?, image = ?, phone = ?, mail = ?, is_open = ?, opening_time = ?, closing_time = ?, grade = ?, is_validated = ?  WHERE id = ?", item.Name, item.Logo, item.Image, item.Phone, item.IsOpen, item.Mail, item.IsOpen, item.OpeningTime, item.ClosingTime, item.Grade, item.IsValidated, restaurantID)
	if err != nil {
		return err
	}

	return nil

}

func (s *RestaurantStore) DeleteRestaurantById(restaurantID int) error {
	_, err := s.DB.Exec("DELETE FROM Restaurants WHERE id = ?", restaurantID)
	if err != nil {
		return err
	}

	return nil
}

//func (t *RestaurantStore) GetRestaurantByID(id int) (*entity.Restaurant, error) {
//	var restaurant *entity.Restaurant
//
//	rows, err := t.Query("SELECT id, name, logo, mail, is_validated  FROM Restaurants WHERE id = ?", id)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//		if err = rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Logo, &restaurant.Mail, &restaurant.IsValidated); err != nil {
//			return nil, err
//		}
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return restaurant, nil
//}
