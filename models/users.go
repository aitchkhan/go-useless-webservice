package models

type User struct {
	ID        int
	firstName string
	lastName  string
}

var (
	users  []*User
	nextID = 1
)
