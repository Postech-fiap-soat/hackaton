package domain

type UserRepository interface {
	GetUserById(userID int) (*User, error)
}

type User struct {
	ID           int
	Name         string
	Registration string
	Email        string
}
