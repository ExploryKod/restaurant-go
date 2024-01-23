package entity

import (
	"database/sql"
)

type User struct {
	ID           int          `json:"id" db:"id"`
	Username     string       `json:"username" db:"username"`
	Password     string       `json:"password" db:"password"`
	Name         string       `json:"name" db:"name"`
	Firstname    string       `json:"firstname" db:"firstname"`
	Mail         string       `json:"mail" db:"mail"`
	Phone        string       `json:"phone" db:"phone"`
	IsSuperadmin bool         `json:"is_superadmin" db:"is_superadmin"`
	Birthday     sql.NullTime `json:"birthday" db:"birthday"`
}

type UserStoreInterface interface {
	GetUserByUsername(username string) (*User, error)
	GetUsers() ([]User, error)
	AddUser(item User) (int, error)
	DeleteUser(id int) error
}
