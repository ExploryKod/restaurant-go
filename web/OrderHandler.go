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

		lastOrderID, orderNumber, err := h.OrderStore.AddOrder(order)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}
		order.ID = lastOrderID
		order.Number = orderNumber

		orderHasProduct := entity.NewOrderHasProduct(*order, *products)

		_, err = h.OrderHasProductStore.AddOrderHasProduct(orderHasProduct)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}

		err = h.Client.Trigger("restaurant-"+strconv.Itoa(restaurantID), "new-order", map[string]any{"data": orderHasProduct, "message": "New order !"})
		if err != nil {
			log.Println(err)
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
		restaurantIDUrl := chi.URLParam(request, "restaurantId")
		restaurantId, _ := strconv.Atoi(restaurantIDUrl)

		req := request.URL.Query().Get("req")

		orders, err := h.OrderHasProductStore.GetAllOrderHasProductsByRestaurantId(restaurantId)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "getallorderhasproduct Internal Server Error"})
			return
		}

		if req == "json" {
			h.RenderJson(writer, http.StatusOK, map[string]any{"message": "Orders found", "data": orders})
			return
		}

		data := restaurantHTTP.TemplateData{Title: "Mes commandes", Content: struct {
			Orders []entity.OrderHasProduct `json:"orders"`
		}{
			Orders: orders,
		}}

		h.RenderHtml(writer, data, "pages/order/list.admin.gohtml")

	}
}
func (h *Handler) ValidateOrderById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ID := chi.URLParam(request, "id")
		idInt, _ := strconv.Atoi(ID)
		// restaurantId := 1
		_, err := h.OrderStore.ValidateOrder(idInt)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "getallorderhasproduct Internal Server Error"})
			return
		}
		http.Redirect(writer, request, "restaurant/order/get", http.StatusSeeOther)
	}
}
func (h *Handler) CompleteOrderById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ID := chi.URLParam(request, "id")
		idInt, _ := strconv.Atoi(ID)
		// restaurantId := 1
		_, err := h.OrderStore.CompleteOrder(idInt)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "getallorderhasproduct Internal Server Error"})
			return
		}
		http.Redirect(writer, request, "restaurant/order/get", http.StatusSeeOther)
	}
}
func (h *Handler) ReadyOrderById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ID := chi.URLParam(request, "id")
		idInt, _ := strconv.Atoi(ID)
		// restaurantId := 1
		_, err := h.OrderStore.ReadyOrder(idInt)
		if err != nil {
			log.Println(err)
			h.RenderJson(writer, http.StatusInternalServerError, map[string]string{"error": "getallorderhasproduct Internal Server Error"})
			return
		}
		http.Redirect(writer, request, "restaurant/order/get", http.StatusSeeOther)
	}
}
func (h *Handler) ShowOrdersPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := restaurantHTTP.TemplateData{Title: "Mes commandes"}
		h.RenderHtml(writer, data, "pages/order/index.gohtml")
	}
}
