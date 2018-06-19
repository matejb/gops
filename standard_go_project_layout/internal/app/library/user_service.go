package library

import (
	"database/sql"
	"net/http"
)

type UserService struct {
	DB *sql.DB
}

func (us UserService) User(id int) (User, error) {
	return User{}, nil
}

func (us UserService) Users() ([]User, error) {
	return nil, nil
}

func (us UserService) CreateUser(user User) error {
	return nil
}

func (us UserService) DeleteUser(user User) error {
	return nil
}

func (us UserService) UpdateUser(user User) error {
	return nil
}

func (us UserService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
