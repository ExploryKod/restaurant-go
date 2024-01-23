package web

import (
	"net/http"
)

func (h *Handler) GetUsers() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		return
	}
}

func (h *Handler) AddUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		return
	}
}

func (h *Handler) DeleteUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		return
	}
}

func (h *Handler) ToggleIsSuperadmin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		return
	}
}
