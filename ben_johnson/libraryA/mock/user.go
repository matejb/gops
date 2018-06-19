package mock

import "github.com/matejb/gops/ben_johnson/libraryA/library"

// UserService represents a mock implementation of myapp.UserService.
type UserService struct {
	UserFn      func(id int) (*library.User, error)
	UserInvoked bool

	UsersFn      func() ([]*library.User, error)
	UsersInvoked bool

	// additional function implementations...
}

// User invokes the mock implementation and marks the function as invoked.
func (s *UserService) User(id int) (*library.User, error) {
	s.UserInvoked = true
	return s.UserFn(id)
}
