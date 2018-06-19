package http

import (
	"net/http"

	"github.com/matejb/gops/ben_johnson/libraryA/library"
)

type Handler struct {
	UserService library.UserService
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle request
}
