package postgres

import (
	"database/sql"

	"github.com/matejb/gops/ben_johnson/libraryB"
)

type UserService struct {
	DB *sql.DB
}

func (us UserService) User(id int) (libraryB.User, error) {
	return libraryB.User{}, nil
}

func (us UserService) Users() ([]libraryB.User, error) {
	return nil, nil
}

func (us UserService) CreateUser(user libraryB.User) error {
	return nil
}

func (us UserService) DeleteUser(user libraryB.User) error {
	return nil
}

func (us UserService) UpdateUser(user libraryB.User) error {
	return nil
}
