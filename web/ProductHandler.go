package web

import (
	"fmt"
	"net/http"
	"restaurantHTTP"
)

func (h *Handler) AddProductType() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			proType := request.FormValue("productType")
			// restaurant := "Constant Restaurant"
			// productType := &entity.User{
			// 	Username:     username,
			// 	Password:     hashedPassword,
			// 	Name:         name,
			// 	Firstname:    firstname,
			// 	Mail:         mail,
			// 	Phone:        phone,
			// 	IsSuperadmin: false,
			// 	Birthday:     []uint8{65, 66, 67, 68, 69},
			// }
			fmt.Println("Error parsing form:", proType)
		}
		fmt.Println("Error adding product type to the database:", request.Method)
		data := restaurantHTTP.TemplateData{Error: "Echec de l'inscription du restaurant"}
		h.RenderHtml(writer, data, "pages/product/productType.create.gohtml")
		return

	}
}
func (h *Handler) AddProduct() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			proType := request.FormValue("productType")
			// restaurant := "Constant Restaurant"
			// productType := &entity.User{
			// 	Username:     username,
			// 	Password:     hashedPassword,
			// 	Name:         name,
			// 	Firstname:    firstname,
			// 	Mail:         mail,
			// 	Phone:        phone,
			// 	IsSuperadmin: false,
			// 	Birthday:     []uint8{65, 66, 67, 68, 69},
			// }
			fmt.Println("Error parsing form:", proType)
		}
		fmt.Println("Error adding product type to the database:", request.Method)
		data := restaurantHTTP.TemplateData{Error: "Echec de l'inscription du restaurant"}
		h.RenderHtml(writer, data, "pages/product/product.create.gohtml")
		return

	}
}
