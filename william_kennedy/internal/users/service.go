package users

import (
	"database/sql"
	"net/http"
)

type Service struct {
	DB *sql.DB
}

func (s Service) User(id int) (User, error) {
	return User{}, nil
}

func (s Service) Users() ([]User, error) {
	return nil, nil
}

func (s Service) CreateUser(user User) error {
	return nil
}

func (s Service) DeleteUser(user User) error {
	return nil
}

func (s Service) UpdateUser(user User) error {
	return nil
}

func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle request
}
