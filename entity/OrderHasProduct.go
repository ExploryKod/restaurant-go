package entity

// OrderHasProduct represent a struct
type OrderHasProduct struct {
	Order   Order   `json:"order"`
	Product Product `json:"product"`
}
