package mock

import "github.com/matejb/gops/ben_johnson/libraryB"

// UserService represents a mock implementation of myapp.UserService.
type UserService struct {
	UserFn      func(id int) (*libraryB.User, error)
	UserInvoked bool

	UsersFn      func() ([]*libraryB.User, error)
	UsersInvoked bool

	// additional function implementations...
}

// User invokes the mock implementation and marks the function as invoked.
func (s *UserService) User(id int) (*libraryB.User, error) {
	s.UserInvoked = true
	return s.UserFn(id)
}
