package restaurantHTTP

import "time"

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Firstname    string    `json:"firstname"`
	Mail         string    `json:"mail"`
	Phone        string    `json:"phone"`
	IsSuperadmin bool      `json:"is_superadmin"`
	Birthday     time.Time `json:"birthday"`
}

type Order struct {
	ID   int  `json:"id"`
	User User `json:"user"`
}

type UserStoreInterface interface {
	GetUsers() ([]User, error)
	AddUser(item User) (int, error)
	DeleteUser(id int) error
}
