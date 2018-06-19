package postgres

import (
	"database/sql"

	"github.com/matejb/gops/ben_johnson/libraryA/library"
)

type UserService struct {
	DB *sql.DB
}

func (us UserService) User(id int) (library.User, error) {
	return library.User{}, nil
}

func (us UserService) Users() ([]library.User, error) {
	return nil, nil
}

func (us UserService) CreateUser(user library.User) error {
	return nil
}

func (us UserService) DeleteUser(user library.User) error {
	return nil
}

func (us UserService) UpdateUser(user library.User) error {
	return nil
}
