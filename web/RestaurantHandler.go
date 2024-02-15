package web

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"net/url"
	"restaurantHTTP"
	"restaurantHTTP/entity"
	"strconv"
)

func (h *Handler) ShowRestaurantsPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-basic")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		restaurants, err := h.RestaurantStore.GetAllRestaurants()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Content: restaurants, Error: "", Success: ""}
			h.RenderHtml(writer, data, "pages/restaurants.client.gohtml")
			return
		}
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) ShowMenuByRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-basic")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		restaurants, err := h.RestaurantStore.GetAllRestaurants()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Content: restaurants, Error: "", Success: ""}
			h.RenderHtml(writer, data, "pages/order/create.gohtml")
			return
		}
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) ShowAddRestaurantAdminPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-basic")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		restaurants, err := h.RestaurantStore.GetAllRestaurants()
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// TODO: limiter à un accés admin seulement (il créé le restaurant à la suite d'un email > formulaire de contact restaurateur > admin)
		// TODO: Refactor (remove) session authenticated in whole project
		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Title: "Inscription d'un nouveau restaurant", Content: restaurants}

			encodedMessage := request.URL.Query().Get("success")
			decodedMessage, err := url.QueryUnescape(encodedMessage)
			if err != nil {
				log.Println("Error during decoding success msg :", err)
				decodedMessage = ""
			}
			if decodedMessage != "" {
				data.Success = decodedMessage
			}

			encodedMessage = request.URL.Query().Get("echec")
			decodedMessage, err = url.QueryUnescape(encodedMessage)
			if err != nil {
				log.Println("Error during decoding echec msg :", err)
				decodedMessage = ""
			}
			if decodedMessage != "" {
				data.Error = decodedMessage
			}

			h.RenderHtml(writer, data, "pages/restaurants.create.gohtml")
			return
		}
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) GetAllRestaurants() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restaurants, err := h.RestaurantStore.GetAllRestaurants()
		if err != nil {
			// Handle database error
			h.RenderJson(w, http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
			})
			return
		}

		h.RenderJson(w, http.StatusOK, restaurants)
	}
}

func (h *Handler) ShowBecomeRestaurantPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		session, err := storeSession.Get(request, "session-basic")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Error: "", Success: ""}
			h.RenderHtml(writer, data, "pages/restaurants.subscribe.gohtml")
			return
		}
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) RegisterRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			http.Error(writer, "cette route n'est disponible qu'en POST", http.StatusBadRequest)
			return
		}

		// Traiter tous les cas d'erreur et renvoyer ces message au front
		// Pourquoi utiliser renderHTML et non redirect - gestion des status d'erreur
		// Interdire de s'inscrire 2 fois
		// Envoyer un email de confirmation d'inscription au restaurateur
		// Attention > rectifier le uint[] dans login birthday

		restaurantName := request.FormValue("restaurant-name")
		restaurantEmail := request.FormValue("restaurant-email")
		restaurantTel := request.FormValue("restaurant-tel")
		restaurantGrade := request.FormValue("restaurant-grade")
		restaurantGradeInt, err := strconv.Atoi(restaurantGrade)
		if err != nil {
			fmt.Println("type parsing for grade failed")
			return
		}
		_, err = h.RestaurantStore.AddRestaurant(entity.Restaurant{Name: restaurantName, Phone: restaurantTel, Mail: restaurantEmail, Grade: restaurantGradeInt, IsValidated: true})
		if err != nil {
			encodedMessage := url.QueryEscape("Echec de l'inscription du restaurant")
			http.Redirect(writer, request, "/restaurant/manage-restaurants?error="+encodedMessage, http.StatusSeeOther)
			return
		}
		encodedMessage := url.QueryEscape(restaurantName + " est inscris dans le FoodCourt.")
		http.Redirect(writer, request, "/restaurant/manage-restaurants?success="+encodedMessage, http.StatusSeeOther)
	}
}

func (h *Handler) AddTagToRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			http.Error(writer, "cette route n'est disponible qu'en POST", http.StatusBadRequest)
			return
		}

		restaurantTag := request.FormValue("restaurant-tag")

		_, err := h.RestaurantStore.AddTagToRestaurant(entity.Tag{Name: restaurantTag})
		if err != nil {
			data := restaurantHTTP.TemplateData{Error: "Echec de l'ajout de tag au restaurant"}
			h.RenderHtml(writer, data, "pages/restaurants.admin.gohtml")
			return
		} else {
			// get tag id of Tag (créer méthode dans model getTags)
			// get id of the restaurant adding the tag > slug via gohtml
			// Jointure
			//_, err h.AddTagToRestaurantHasTag()
		}
		data := restaurantHTTP.TemplateData{Title: "", Success: "Nouveau tag" + restaurantTag + "ajouté"}
		h.RenderHtml(writer, data, "pages/restaurants.admin.gohtml")
		return
	}
}

func (h *Handler) ShowAdminRestaurantPage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		restaurantId := 1
		// TODO: get id to have it in the page and link to restaurant / tag
		session, err := storeSession.Get(request, "session-basic")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		restaurant := h.RestaurantStore.GetRestaurantByID(restaurantId)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Title: "", Content: restaurant}
			h.RenderHtml(writer, data, "pages/restaurants.admin.gohtml")
			return
		}
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
	}
}

func (h *Handler) ShowRestaurantUpdatePage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		restaurantId := chi.URLParam(request, "id")
		restaurantIdInt, _ := strconv.Atoi(restaurantId)
		// TODO: get id to have it in the page and link to restaurant / tag
		session, err := storeSession.Get(request, "session-basic")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		restaurant := h.RestaurantStore.GetRestaurantByID(restaurantIdInt)
		if restaurant == nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["authenticated"] != nil && session.Values["authenticated"].(bool) {
			data := restaurantHTTP.TemplateData{Title: "", Content: restaurant}
			h.RenderHtml(writer, data, "pages/restaurants.update.gohtml")
		}
		http.Redirect(writer, request, "/restaurant/manage-restaurants", http.StatusSeeOther)
	}
}

func (h *Handler) UpdateRestaurantHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		restaurantName := request.FormValue("restaurant-name")
		restaurantID, err := strconv.Atoi(request.FormValue("restaurant-id"))
		if err != nil {
			println("restaurantId parsing failed %s", err)
			return
		}
		//restaurantLogo := r.FormValue("restaurant-logo")
		//restaurantImage := r.FormValue("restaurant-image")
		//restaurantPhone := r.FormValue("restaurant-phone")
		//restaurantMail := r.FormValue("restaurant-mail")
		//restaurantIsOpen := r.FormValue("restaurant-isopen") == "open"
		//// TODO: need conversion time.tIME
		//restaurantOpeningTime := r.FormValue("restaurant-openingtime")
		//restaurantClosingTime := r.FormValue("restaurant-closingtime")
		//
		//restaurantGrade, err := strconv.Atoi(r.FormValue("restaurant-grade"))
		//if err != nil {
		//	fmt.Println("restaurant grade parsing type failed")
		//	return
		//}
		//restaurantIsValidated := r.FormValue("restaurant-isvalidated") == "validated"

		//layout := "2006-01-02 15:04:05"
		//layout := "15:04:05"
		//restaurantOpeningHours, _ := time.Parse(layout, restaurantOpeningTime)
		//restaurantClosingHours, _ := time.Parse(layout, restaurantClosingTime)
		//fmt.Printf("opening time parsing: %v\n", restaurantOpeningHours)
		//fmt.Printf("closing time parsing: %v\n", restaurantOpeningHours)

		err = h.RestaurantStore.UpdateRestaurant(entity.Restaurant{
			ID:   restaurantID,
			Name: restaurantName,
			//Logo:        restaurantLogo,
			//Image:       restaurantImage,
			//Phone:       restaurantPhone,
			//Mail:        restaurantMail,
			//IsOpen:      restaurantIsOpen,
			//Grade:       restaurantGrade,
			//ClosingTime: restaurantClosingHours,
			//OpeningTime: restaurantOpeningHours,
			//IsValidated: restaurantIsValidated
		},
			restaurantID)
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		data := restaurantHTTP.TemplateData{Title: "", Success: restaurantName + "modifié"}
		h.RenderHtml(writer, data, "pages/restaurants.update.gohtml")
		return
	}
}

func (h *Handler) DeleteRestaurantHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")

		id, _ := strconv.Atoi(QueryId)

		err := h.RestaurantStore.DeleteRestaurantById(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/restaurant/manage-restaurants", http.StatusSeeOther)

	}
}
