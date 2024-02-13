package web

import (
	"fmt"
	"net/http"
	"net/url"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"strconv"

	"strings"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) AddProductType() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		restaurantId := chi.URLParam(request, "restaurantId")
		restaurantIdInt, _ := strconv.Atoi(restaurantId)
		if request.Method == http.MethodPost {
			menu := request.FormValue("menu")
			icon := request.FormValue("icon")
			productType := &entity.ProductType{
				Name:         menu,
				Icon:         icon,
				RestaurantId: restaurantIdInt,
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
		restaurantId := chi.URLParam(request, "restaurantId")
		restaurantIdInt, _ := strconv.Atoi(restaurantId)
		if request.Method == http.MethodPost {
			productName := request.FormValue("productName")
			menuVal := request.FormValue("menuDropdown")
			priceString := request.FormValue("price")
			image := "https://source.unsplash.com/300x150/?" + request.FormValue("image") + strings.ReplaceAll(productName, " ", "")
			description := request.FormValue("description")
			// allergyVal := request.FormValue("allergyDropDown")
			menuValInt, _ := strconv.Atoi(menuVal)
			priceFloat, _ := strconv.ParseFloat(priceString, 64)
			restaurant := h.RestaurantStore.GetRestaurantByID(restaurantIdInt)
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
		productsType, _ := h.ProductTypeStore.GetProductTypeByRestaurantId(restaurantIdInt)
		//  Allergy
		// allergies, err := h.ProductStore.GetAllergiesList()
		// obj := ProductTypeAllergy{
		// 	productType: productsType,
		// 	allergies:   allergies,
		// }
		data := restaurantHTTP.TemplateData{Error: "Echec de l'ajouter un produit", Content: productsType}
		h.RenderHtml(writer, data, "pages/product/product.create.gohtml")
		return
	}
}
func (h *Handler) ListProducts() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		restaurantId := chi.URLParam(request, "restaurantId")
		restaurantIdInt, _ := strconv.Atoi(restaurantId)
		productsList, _ := h.ProductStore.GetProductByRestaurantId(restaurantIdInt)
		//  Allergy
		// allergies, err := h.ProductStore.GetAllergiesList()
		// obj := ProductTypeAllergy{
		// 	productType: productsType,
		// 	allergies:   allergies,
		// }
		// h.RenderJson(writer, http.StatusOK, map[string]any{"message": "Orders found", "data": productsList})
		data := restaurantHTTP.TemplateData{Error: "Echec de l'ajouter un produit", Content: productsList}
		h.RenderHtml(writer, data, "pages/product/product.list.gohtml")
		return
	}
}
func (h *Handler) DeleteProducts() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		id := chi.URLParam(request, "id")
		resId := chi.URLParam(request, "restaurantId")
		idInt, _ := strconv.Atoi(id)
		resIdInt, _ := strconv.Atoi(resId)
		productsList, _ := h.ProductStore.DeleteProduct(idInt)
		if productsList {
			fmt.Println("Error in productlist")
		}
		http.Redirect(writer, request, fmt.Sprintf("/product/list/%d", resIdInt), http.StatusSeeOther)
		//  Allergy

		// allergies, err := h.ProductStore.GetAllergiesList()
		// obj := ProductTypeAllergy{
		// 	productType: productsType,
		// 	allergies:   allergies,
		// }
		// h.RenderJson(writer, http.StatusOK, map[string]any{"message": "Orders found", "data": productsList})
	}
}
