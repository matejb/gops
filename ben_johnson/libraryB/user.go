package libraryB

type User struct {
	ID    int
	Name  string
	Email string
}

type UserService interface {
	User(id int) (User, error)
	Users() ([]User, error)
	CreateUser(user User) error
	DeleteUser(user User) error
	UpdateUser(user User) error
}
