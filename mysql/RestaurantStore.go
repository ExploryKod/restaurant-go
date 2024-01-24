package database

import "restaurantHTTP/entity"

func (t *UserStore) AddRestaurant(item entity.Restaurant) (int, error) {
	res, err := t.DB.Exec("INSERT INTO restaurants (name, mail, is_validated) VALUES ( ? , ? , ?)", item, item, item)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
