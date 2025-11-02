package web

import (
	"database/sql"
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

		restaurantIDUrl := chi.URLParam(request, "id")

		_, claims, err := jwtauth.FromContext(request.Context())
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusUnauthorized, map[string]string{"error": "jwt problem"})
			return
		}

		/* products */
		// Structure temporaire pour décoder les données du frontend
		type CartProduct struct {
			ID          int     `json:"id"`
			MenuID      int     `json:"menuid"`
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Price       float64 `json:"price"`
			Image       string  `json:"image"`
			Count       int     `json:"count"`
			Total       float64 `json:"total"`
		}

		var cartProducts []CartProduct
		if err := json.NewDecoder(request.Body).Decode(&cartProducts); err != nil {
			log.Println("Error decoding products:", err)
			h.RenderJson(writer, http.StatusBadRequest, map[string]string{"error": "Invalid products for the order"})
			return
		}

		// Convertir les produits du panier en produits avec quantité
		var products []entity.Product
		for _, cartProduct := range cartProducts {
			// Ajouter le produit autant de fois que sa quantité
			for i := 0; i < cartProduct.Count; i++ {
				products = append(products, entity.Product{
					ID:          cartProduct.ID,
					Name:        cartProduct.Name,
					Description: cartProduct.Description,
					Price:       cartProduct.Price,
					Image:       cartProduct.Image,
				})
			}
		}

		/*get user*/
		userID := int(claims["id"].(float64))

		user := h.UserStore.GetUserByID(userID)
		if user == nil {
			h.RenderJson(writer, http.StatusUnauthorized, map[string]string{"error": "user not found"})
			return
		}

		/*order*/
		restaurantID, err := strconv.Atoi(restaurantIDUrl)
		if err != nil {
			log.Println("Error parsing restaurant ID:", restaurantIDUrl, err)
			h.RenderJson(writer, http.StatusBadRequest, map[string]string{"error": "Invalid restaurant ID"})
			return
		}
		
		restaurant := h.RestaurantStore.GetRestaurantByID(restaurantID)
		if restaurant == nil {
			log.Printf("Restaurant with ID %d not found", restaurantID)
			h.RenderJson(writer, http.StatusNotFound, map[string]string{"error": fmt.Sprintf("Restaurant with ID %d not found", restaurantID)})
			return
		}
		
		log.Printf("Creating order for restaurant ID: %d, Name: %s", restaurant.ID, restaurant.Name)

		// Calculer le total en utilisant les données du panier (prend en compte les quantités)
		totalPrice := 0.0
		for _, cartProduct := range cartProducts {
			totalPrice += cartProduct.Total
		}

		order := entity.NewOrder(*user, *restaurant, "pending", totalPrice, 0, time.Now(), sql.NullTime{})

		var lastOrderID int
		lastOrderID, err = h.OrderStore.AddOrder(*order)
		if err != nil {
			log.Println("Error adding order:", err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}
		order.ID = lastOrderID

		orderHasProduct := entity.NewOrderHasProduct(*order, products)

		_, err = h.OrderHasProductStore.AddOrderHasProduct(orderHasProduct)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}

		h.RenderJson(writer, http.StatusOK, map[string]any{"message": "Order created successfully!", "data": orderHasProduct})
	}
}

func (h *Handler) GetOrder() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ID := chi.URLParam(request, "id")

		orderID, _ := strconv.Atoi(ID)

		order := h.OrderStore.GetOrderByID(orderID)
		if order == nil {
			h.RenderJson(writer, http.StatusNotFound, map[string]string{"error": "order not found"})
			return
		}

		orderHasProduct := h.OrderHasProductStore.GetOrderHasProductByOrderID(orderID)
		if orderHasProduct == nil {
			h.RenderJson(writer, http.StatusNotFound, map[string]string{"error": "order has product not found"})
			return
		}

		orderHasProduct.Order = *order

		h.RenderJson(writer, http.StatusOK, map[string]any{"message": "Order found", "data": orderHasProduct})
	}
}

func (h *Handler) GetAllOrders() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		orders, err := h.OrderHasProductStore.GetAllOrderHasProducts()
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "getallorderhasproduct Internal Server Error"})
			return
		}

		data := restaurantHTTP.TemplateData{Title: "Mes commandes", Content: orders}

		h.RenderHtml(writer, data, "pages/order/index.gohtml")
		//h.RenderJson(writer, http.StatusOK, map[string]any{"message": "Orders found", "data": orders})

	}
}

func (h *Handler) ShowOrdersPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := restaurantHTTP.TemplateData{Title: "Mes commandes"}
		h.RenderHtml(writer, data, "pages/order/index.gohtml")
	}
}
