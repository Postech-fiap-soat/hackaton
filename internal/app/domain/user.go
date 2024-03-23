package domain

type UserRepository interface {
	GetUserById(userID int) (*User, error)
	GetUserByRegistration(registration string) (*User, error)
}

type UserUseCase interface {
	Login(loginDTO LoginDTO) (*JWT, error)
}
type User struct {
	ID           int
	Name         string
	Registration string
	Email        string
	Password     string
}

type LoginDTO struct {
	Registration string `json:"registration"`
	Password     string `json:"password"`
}
type JWT struct {
	Token string
}

func NewJWT(user *User) *JWT {
	return &JWT{Token: "exemplo token"}
}
