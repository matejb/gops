package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "connection string")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create services.
	us := &userService{DB: db}

	// Attach to HTTP handler.
	var h Handler
	h.UserService = us

	// Start the server.
	err = http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatal(err)
	}
}
