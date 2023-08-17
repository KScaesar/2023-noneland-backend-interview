package entity

type UserRepository interface {
	// User
	GetUsers() (users []User, err error)
}

type User struct {
	Name string
}
