package library

type Book struct {
	ID       int
	Name     string
	Author   string
	Borrowed *User
}
