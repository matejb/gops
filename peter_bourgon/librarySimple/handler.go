package main

import (
	"net/http"
)

type Handler struct {
	UserService *userService
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle request
}
