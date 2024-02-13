package database

import (
	"github.com/jmoiron/sqlx"
	"log"
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

		if err = rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Logo,
			&restaurant.Image,
			&restaurant.Phone,
			&restaurant.Mail,
			&restaurant.IsOpen,
			&restaurant.OpeningTime,
			&restaurant.ClosingTime,
			&restaurant.Grade,
			&restaurant.IsValidated); err != nil {
			return []entity.Restaurant{}, err
		}

		// Format the time as HH:mm:ss
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
		log.Println(err)
		return nil
	}
	return restaurant
}

func (s *RestaurantStore) UpdateRestaurant(item entity.Restaurant, restaurantID int) error {

	//_, err := s.DB.Exec("UPDATE Restaurants SET name = ?, logo = ?, image = ?, phone = ?, mail = ?, is_open = ?, opening_time = ?, closing_time = ?, grade = ?, is_validated = ?  WHERE id = ?", item.Name, item.Logo, item.Image, item.Phone, item.Mail, item.IsOpen, item.OpeningTime, item.ClosingTime, item.Grade, item.IsValidated, restaurantID)
	//if err != nil {
	//	return err
	//}

	_, err := s.DB.Exec("UPDATE Restaurants SET name = ?  WHERE id = ?", item.Name, restaurantID)
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
