package entity

// RestaurantHasUsers represent a struct
type RestaurantHasUsers struct {
	Restaurant Restaurant `json:"restaurant"`
	User       User       `json:"user"`
	IsAdmin    bool       `json:"is_admin"`
	Role       string     `json:"role"`
}
