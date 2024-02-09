package web

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"restaurantHTTP"
	"restaurantHTTP/entity"
)

func (h *Handler) CreateOrder() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {

			// recup param id
			ID := chi.URLParam(request, "id")

			templateData := restaurantHTTP.TemplateData{Title: "Create Order", Content: struct{ IdRestaurant string }{IdRestaurant: ID}}
			h.RenderHtml(writer, templateData, "pages/order/create.gohtml")
			return
		}

		// extract products (json) from request post

		var products []entity.Product

		if err := json.NewDecoder(request.Body).Decode(&products); err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusBadRequest, map[string]string{"error": "Invalid products"})
			return
		}
		for _, product := range products {
			fmt.Printf("ID: %d\n", product.ID)
			fmt.Printf("Name: %s\n", product.Name)
			fmt.Printf("Price: %d\n", product.Price)
			fmt.Printf("Image: %s\n", product.Image)
			fmt.Printf("Description: %s\n", product.Description)
			fmt.Println("-------------------------------------")
		}

		h.RenderJson(writer, http.StatusOK, map[string]any{"message": "Order created successfully!", "products": products})

	}
}
