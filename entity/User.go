package entity

import "time"

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Firstname    string    `json:"firstname"`
	Mail         string    `json:"mail"`
	Phone        string    `json:"phone"`
	IsSuperAdmin bool      `json:"is_super_admin"`
	Birthday     time.Time `json:"birthday"`
}
