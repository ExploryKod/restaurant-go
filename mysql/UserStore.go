package database

import (
	"database/sql"
	"restaurantHTTP"
)

type UserStore struct {
	*sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

func (t *UserStore) GetUsers() ([]restaurantHTTP.User, error) {
	var todos []restaurantHTTP.User

	rows, err := t.Query("SELECT id, title, completed FROM users")
	if err != nil {
		return []restaurantHTTP.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var user restaurantHTTP.User
		if err = rows.Scan(&user.ID /* donnée à ajouter*/); err != nil {
			return []restaurantHTTP.User{}, err
		}
		todos = append(todos, user)
	}

	if err = rows.Err(); err != nil {
		return []restaurantHTTP.User{}, err
	}

	return todos, nil
}

func (t *UserStore) AddUser(item restaurantHTTP.User) (int, error) {
	res, err := t.DB.Exec("INSERT INTO users ( , ) VALUES ( , )", item, item)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *UserStore) DeleteUser(id int) error {
	_, err := t.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (t *UserStore) ToggleIsAdmin(id int) error {
	_, err := t.DB.Exec("UPDATE users SET `is_superadmin` = IF (`is_superadmin`, 0, 1) WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
