package web

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"strconv"
	"time"
)

func (h *Handler) CreateOrder() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {

			ID := chi.URLParam(request, "id")

			templateData := restaurantHTTP.TemplateData{Title: "Create Order", Content: struct{ IdRestaurant string }{IdRestaurant: ID}}
			h.RenderHtml(writer, templateData, "pages/order/create.gohtml")
			return
		}

		restaurantIDUrl := request.URL.Query().Get("restaurant_id")

		_, claims, err := jwtauth.FromContext(request.Context())
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusUnauthorized, map[string]string{"error": "jwt problem"})
			return
		}

		/* products */
		var products []entity.Product

		if err := json.NewDecoder(request.Body).Decode(&products); err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusBadRequest, map[string]string{"error": "Invalid products for the order"})
			return
		}

		/*get user*/
		userID := int(claims["id"].(float64))

		user := h.UserStore.GetUserByID(userID)
		if user == nil {
			h.RenderJson(writer, http.StatusUnauthorized, map[string]string{"error": "user not found"})
			return
		}

		/*order*/
		restaurantID, _ := strconv.Atoi(restaurantIDUrl)
		restaurant := h.RestaurantStore.GetRestaurantByID(restaurantID)

		order := entity.NewOrder(*user, *restaurant, "pending", 10.23, time.Now(), time.Now())

		_, err = h.OrderStore.AddOrder(*order)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}

		orderHasProduct := entity.NewOrderHasProduct(*order, products)

		_, err = h.OrderHasProductStore.AddOrderHasProduct(*orderHasProduct)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
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
