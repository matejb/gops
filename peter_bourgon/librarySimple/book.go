package main

type Book struct {
	ID       int
	Name     string
	Author   string
	Borrowed *User
}
