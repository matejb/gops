package main

import (
	"log"
	"net/http"
	"os"

	"github.com/matejb/gops/william_kennedy/internal/platform/database"
	"github.com/matejb/gops/william_kennedy/internal/users"
)

func main() {
	db, err := database.Open(os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}

	userService := users.Service{DB: db}

	err = http.ListenAndServe(":8080", userService)
	if err != nil {
		log.Fatal(err)
	}
}
