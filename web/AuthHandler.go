package web

import (
	"fmt"
	"net/http"
)

func (h *Handler) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// handler login from post request
		username := request.FormValue("username")
		password := request.FormValue("password")

		fmt.Println(username, password)

		// check if user exist

		// if user exist, check if password is correct

		// if password is correct, create a session

		// if password is incorrect, return error

		// if user doesn't exist, return error

	}
}

func (h *Handler) Signup() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// handler signup from post request
		username := request.FormValue("username")
		password := request.FormValue("password")

		fmt.Println(username, password)

		// check if user exist

		// if user exist, return error

		// if user doesn't exist, create user

	}
}
