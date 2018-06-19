package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/matejb/gops/standard_go_project_layout/internal/app/library"
)

func main() {
	db, err := sql.Open("postgres", "connection string")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	us := library.UserService{DB: db}

	err = http.ListenAndServe(":8080", us)
	if err != nil {
		log.Fatal(err)
	}

}
