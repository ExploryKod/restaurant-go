package web

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"strconv"
)

func (h *Handler) RestaurantUserList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		RestaurantId, _ := strconv.Atoi(chi.URLParam(request, "restaurantId"))
		Users, _ := h.RestaurantUserStore.GetRestaurantUsers(RestaurantId)
		content := struct {
			Users        []entity.RestaurantHasUsers
			RestaurantId int
		}{
			Users,
			RestaurantId,
		}
		h.RenderHtml(writer, restaurantHTTP.TemplateData{Content: content}, "pages/restaurantUser/restaurantUser.list.gohtml")
	}
}

func (h *Handler) RestaurantUserCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		RestaurantId, _ := strconv.Atoi(chi.URLParam(request, "restaurantId"))
		if request.Method == http.MethodPost {
			email := request.FormValue("email")
			user, _ := h.UserStore.GetUserByMail(email)
			role := request.FormValue("role")
			isAdminStr := request.FormValue("is_admin")
			isAdmin := false
			if isAdminStr != "" {
				isAdmin = true
			}

			restaurant := h.RestaurantStore.GetRestaurantByID(RestaurantId)
			RestaurantUser := &entity.RestaurantHasUsers{
				User:       *user,
				IsAdmin:    isAdmin,
				Restaurant: *restaurant,
				Role:       role,
			}
			//var id int
			//var err error   declare err here
			_, _ = h.RestaurantUserStore.AddRestaurantUser(*RestaurantUser)
			http.Redirect(writer, request, fmt.Sprintf("/restaurantUser/list/%d", RestaurantId), http.StatusSeeOther)
		} else {
			h.RenderHtml(writer, restaurantHTTP.TemplateData{Content: RestaurantId}, "pages/restaurantUser/restaurantUser.create.gohtml")
		}
	}
}

func (h *Handler) RestaurantUserUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		RestaurantId, _ := strconv.Atoi(chi.URLParam(request, "restaurantId"))

		id, _ := strconv.Atoi(chi.URLParam(request, "id"))
		if request.Method == http.MethodPost {
			role := request.FormValue("role")
			user := h.UserStore.GetUserByID(id)
			isAdminStr := request.FormValue("is_admin")
			isAdmin, _ := strconv.ParseBool(isAdminStr)
			restaurant := h.RestaurantStore.GetRestaurantByID(RestaurantId)
			RestaurantUser := &entity.RestaurantHasUsers{
				User:       *user,
				IsAdmin:    isAdmin,
				Restaurant: *restaurant,
				Role:       role,
			}
			_ = h.RestaurantUserStore.UpdateRestaurantUser(*RestaurantUser)
			http.Redirect(writer, request, fmt.Sprintf("/restaurantUser/list/%d", RestaurantId), http.StatusSeeOther)
		} else {
			User, _ := h.RestaurantUserStore.GetRestaurantUserByUserID(id)
			content := struct {
				User         *entity.RestaurantHasUsers
				RestaurantId int
			}{
				User,
				RestaurantId,
			}

			h.RenderHtml(writer, restaurantHTTP.TemplateData{Content: content}, "pages/restaurantUser/restaurantUser.update.gohtml")
		}
	}
}

func (h *Handler) RestaurantUserDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		restaurantId, _ := strconv.Atoi(chi.URLParam(request, "restaurantId"))

		id, _ := strconv.Atoi(chi.URLParam(request, "id"))
		fmt.Println(request.Method)

		_ = h.RestaurantUserStore.DeleteRestaurantUser(id)
		http.Redirect(writer, request, fmt.Sprintf("/restaurantUser/list/%d", restaurantId), http.StatusSeeOther)

	}
}
