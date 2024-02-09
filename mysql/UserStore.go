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

func (u *UserStore) GetUserByID(id int) *entity.User {
	user := &entity.User{}
	err := u.Get(user, "SELECT id, username, password, name, firstname, mail, phone, is_superadmin, birthday FROM Users WHERE id = ?", id)
	if err != nil {
		return nil
	}
	return user
}

func (u *UserStore) GetUserByUsername(username string) (*entity.User, error) {

	user := &entity.User{}

	err := u.Get(user, "SELECT id, username, password, name, firstname, mail, phone, is_superadmin, birthday FROM Users WHERE username = ?", username)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserStore) GetUserByMail(mail string) (*entity.User, error) {

	user := &entity.User{}

	err := u.Get(user, "SELECT id, username, password, name, firstname, mail, phone, is_superadmin, birthday FROM Users WHERE mail = ?", mail)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserStore) GetUsers() ([]entity.User, error) {
	var todos []entity.User

	rows, err := u.Query("SELECT id, title, completed FROM users")
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

func (u *UserStore) AddUser(item *entity.User) (int, error) {
	res, err := u.DB.Exec("INSERT INTO Users (username, password, name, firstname, mail, phone, is_superadmin, birthday) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", item.Username, item.Password, item.Name, item.Firstname, item.Mail, item.Phone, item.IsSuperadmin, item.Birthday)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (u *UserStore) DeleteUser(id int) error {
	_, err := u.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserStore) ToggleIsAdmin(id int) error {
	_, err := u.DB.Exec("UPDATE users SET `is_superadmin` = IF (`is_superadmin`, 0, 1) WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
