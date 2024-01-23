package database

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"restaurantHTTP/entity"
)

type UserStore struct {
	*sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		db,
	}
}

func (t *UserStore) GetUserByUsername(username string) (*entity.User, error) {

	user := &entity.User{}

	err := t.Get(user, "SELECT id, username, password, name, firstname, mail, phone, is_superadmin, birthday FROM Users WHERE username = ?", username)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (t *UserStore) GetUsers() ([]entity.User, error) {
	var todos []entity.User

	rows, err := t.Query("SELECT id, title, completed FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err = rows.Scan(&user.ID /* donnée à ajouter*/); err != nil {
			return []entity.User{}, err
		}
		todos = append(todos, user)
	}

	if err = rows.Err(); err != nil {
		return []entity.User{}, err
	}

	return todos, nil
}

func (t *UserStore) AddUser(item entity.User) (int, error) {
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
