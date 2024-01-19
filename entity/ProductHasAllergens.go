package entity

// ProductHasAllergens represent a struct
type ProductHasAllergens struct {
	Product  Product  `json:"product"`
	Allergen Allergen `json:"allergen"`
}
