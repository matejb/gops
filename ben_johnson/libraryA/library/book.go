package library

type Book struct {
	ID       int
	Name     string
	Author   string
	Borrowed *User
}

type BookService interface {
	Book(id int) (Book, error)
	Books() ([]Book, error)
	Borrow(user User) error
	Return() error
}
