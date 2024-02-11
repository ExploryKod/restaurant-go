package entity

// RestaurantHasUsers represent a struct
type RestaurantHasUsers struct {
	Restaurant Restaurant `json:"restaurant"`
	User       User       `json:"user"`
	IsAdmin    bool       `json:"is_admin"`
	Role       string     `json:"role"`
}

// RestaurantUserStoreInterface represent a struct
type RestaurantUserStoreInterface interface {
	AddRestaurantUser(item RestaurantHasUsers) (int, error)
	UpdateRestaurantUser(item RestaurantHasUsers) error
	DeleteRestaurantUser(userId int) error
	GetRestaurantUsers(restaurantId int) ([]RestaurantHasUsers, error)
	GetRestaurantUserByUserID(userId int) (*RestaurantHasUsers, error)
	GetRestaurantIDByUserID(userId int) (*int, error)
}
