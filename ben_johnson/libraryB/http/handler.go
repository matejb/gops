package http

import (
	"net/http"

	"github.com/matejb/gops/ben_johnson/libraryB"
)

type Handler struct {
	UserService libraryB.UserService
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle request
}
