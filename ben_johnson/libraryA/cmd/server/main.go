package main

import (
	"log"
	"os"

	"github.com/matejb/gops/ben_johnson/libraryA/http"
	"github.com/matejb/gops/ben_johnson/libraryA/postgres"
)

func main() {
	// Init dependencies like DB connection.
	db, err := postgres.Open(os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create services.
	us := &postgres.UserService{DB: db}

	// Attach to HTTP handler.
	var h http.Handler
	h.UserService = us

	// Start the server.
	err = http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatal(err)
	}
}
