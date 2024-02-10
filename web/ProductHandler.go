package web

import (
	"fmt"
	"net/http"
	"net/url"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"strconv"
)

func (h *Handler) AddProductType() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			menu := request.FormValue("menu")
			icon := request.FormValue("icon")
			// restaurantId := 12
			productType := &entity.ProductType{
				Name:         menu,
				Icon:         icon,
				RestaurantId: 1,
			}
			var id int
			var err error // declare err here
			id, err = h.ProductTypeStore.AddProductType(*productType)
			if err != nil {
				fmt.Println(err)
				return
			}
			message := "Account created successfully with id: " + fmt.Sprintf("%d", id)
			encodedMessage := url.QueryEscape(message)
			fmt.Println("Error parsing form:", encodedMessage)
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
			fmt.Println("Error parsing formsss:")
			productName := request.FormValue("productName")
			menuVal := request.FormValue("menuDropdown")
			priceString := request.FormValue("price")
			image := request.FormValue("image")
			description := request.FormValue("description")
			// allergyVal := request.FormValue("allergyDropDown")
			restaurantId := 1
			menuValInt, _ := strconv.Atoi(menuVal)
			priceFloat, _ := strconv.ParseFloat(priceString, 64)
			restaurant := h.RestaurantStore.GetRestaurantByID(restaurantId)
			productType, _ := h.ProductTypeStore.GetProductTypeById(menuValInt)

			product := &entity.Product{
				Name:        productName,
				ProductType: *productType,
				Restaurant:  *restaurant,
				Price:       priceFloat,
				Image:       image,
				Description: description,
			}

			var id int
			var err error // declare err here
			id, err = h.ProductStore.AddProduct(*product)
			if err != nil {
				fmt.Println(err)
				return
			}
			message := "Product created successfully with id: " + fmt.Sprintf("%d", id)
			encodedMessage := url.QueryEscape(message)
			fmt.Println("Error parsing form:", encodedMessage)
		}
		allergies, err := h.ProductStore.GetAllergiesList()
		if err != nil {

		}
		fmt.Println("Error parsing form:", allergies)
		data := restaurantHTTP.TemplateData{Error: "Echec de l'inscription du restaurant", Content: allergies}
		h.RenderHtml(writer, data, "pages/product/product.create.gohtml")
		return

	}
}
