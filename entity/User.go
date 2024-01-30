package entity

type User struct {
	ID           int     `json:"id" db:"id"`
	Username     string  `json:"username" db:"username"`
	Password     string  `json:"password" db:"password"`
	Name         string  `json:"name" db:"name"`
	Firstname    string  `json:"firstname" db:"firstname"`
	Mail         string  `json:"mail" db:"mail"`
	Phone        string  `json:"phone" db:"phone"`
	IsSuperadmin bool    `json:"is_superadmin" db:"is_superadmin"`
	Birthday     []uint8 `json:"birthday" db:"birthday"`
}

func NewUser(username string, password string, name string, firstname string, mail string, phone string, isSuperadmin bool, birthday sql.NullTime) *User {
	return &User{
		Username:     username,
		Password:     password,
		Name:         name,
		Firstname:    firstname,
		Mail:         mail,
		Phone:        phone,
		IsSuperadmin: isSuperadmin,
		Birthday:     birthday,
	}
}

type UserStoreInterface interface {
	GetUserByUsername(username string) (*User, error)
	GetUserByMail(mail string) (*User, error)
	GetUsers() ([]User, error)
	AddUser(item *User) (int, error)
	DeleteUser(id int) error
}
