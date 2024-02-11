package web

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"strconv"
)

func (h *Handler) RestaurantUserList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		restaurantId := 1
		users, _ := h.RestaurantUserStore.GetRestaurantUsers(restaurantId)
		h.RenderHtml(writer, restaurantHTTP.TemplateData{Content: users}, "pages/restaurantUser/restaurantUser.list.gohtml")
	}
}

func (h *Handler) RestaurantUserCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			email := request.FormValue("email")
			user, _ := h.UserStore.GetUserByMail(email)
			role := request.FormValue("role")
			//isAdminStr := request.FormValue("is_admin")
			//isAdmin, _ := strconv.ParseBool(isAdminStr)
			isAdminStr := request.FormValue("is_admin")
			isAdmin := false
			if isAdminStr != "" {
				isAdmin = true
			}
			restaurantId := 1
			restaurant := h.RestaurantStore.GetRestaurantByID(restaurantId)
			RestaurantUser := &entity.RestaurantHasUsers{
				User:       *user,
				IsAdmin:    isAdmin,
				Restaurant: *restaurant,
				Role:       role,
			}
			//var id int
			//var err error   declare err here
			_, _ = h.RestaurantUserStore.AddRestaurantUser(*RestaurantUser)
			http.Redirect(writer, request, "/restaurantUser/list", http.StatusSeeOther)
		} else {
			h.RenderHtml(writer, restaurantHTTP.TemplateData{}, "pages/restaurantUser/restaurantUser.create.gohtml")
		}
	}
}

func (h *Handler) RestaurantUserUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(request, "id"))
		if request.Method == http.MethodPost {
			role := request.FormValue("role")
			user := h.UserStore.GetUserByID(id)
			isAdminStr := request.FormValue("is_admin")
			isAdmin, _ := strconv.ParseBool(isAdminStr)
			restaurantId := 1
			restaurant := h.RestaurantStore.GetRestaurantByID(restaurantId)
			RestaurantUser := &entity.RestaurantHasUsers{
				User:       *user,
				IsAdmin:    isAdmin,
				Restaurant: *restaurant,
				Role:       role,
			}
			_ = h.RestaurantUserStore.UpdateRestaurantUser(*RestaurantUser)
			http.Redirect(writer, request, "/restaurantUser/list", http.StatusSeeOther)
		} else {
			h.RenderHtml(writer, restaurantHTTP.TemplateData{}, "pages/restaurantUser/restaurantUser.update.gohtml")
		}
	}
}

func (h *Handler) RestaurantUserDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(request, "id"))
		if request.Method == http.MethodPost {
			_ = h.RestaurantUserStore.DeleteRestaurantUser(id)
			http.Redirect(writer, request, "/restaurantUser/list", http.StatusSeeOther)
		}
	}
}
