package models

// User is a user of the system
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)
