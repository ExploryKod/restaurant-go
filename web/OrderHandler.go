package web

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
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
		var products *[]entity.Product

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
		if restaurant == nil {
			h.RenderJson(writer, http.StatusNotFound, map[string]string{"error": "restaurant not found"})
			return
		}

		totalPrice := 0.0
		for _, product := range *products {
			totalPrice += float64(product.Price)
		}

		order := entity.NewOrder(*user, *restaurant, "pending", totalPrice, 0, time.Now(), sql.NullTime{})

		var lastOrderID int
		lastOrderID, err = h.OrderStore.AddOrder(*order)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}
		order.ID = lastOrderID

		orderHasProduct := entity.NewOrderHasProduct(*order, *products)

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

func (h *Handler) GetAllOrdersByRestaurantId() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		restaurantId := 1
		orders, err := h.OrderStore.GetOrderByRestaurantID(restaurantId)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "getallorderhasproduct Internal Server Error"})
			return
		}

		data := restaurantHTTP.TemplateData{Title: "Mes commandes", Content: orders}

		h.RenderHtml(writer, data, "pages/order/list.admin.gohtml")
		//h.RenderJson(writer, http.StatusOK, map[string]any{"message": "Orders found", "data": orders})

	}
}

func (h *Handler) ShowOrdersPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := restaurantHTTP.TemplateData{Title: "Mes commandes"}
		h.RenderHtml(writer, data, "pages/order/index.gohtml")
	}
}
