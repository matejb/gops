package main

import (
	"database/sql"
)

type userService struct {
	DB *sql.DB
}

func (us userService) User(id int) (User, error) {
	return User{}, nil
}

func (us userService) Users() ([]User, error) {
	return nil, nil
}

func (us userService) CreateUser(user User) error {
	return nil
}

func (us userService) DeleteUser(user User) error {
	return nil
}

func (us userService) UpdateUser(user User) error {
	return nil
}
