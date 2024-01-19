package entity

// RestaurantHasTags represent a struct
type RestaurantHasTags struct {
	Restaurant Restaurant `json:"restaurant"`
	Tag        Tag        `json:"tag"`
}
