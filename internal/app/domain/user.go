package domain

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

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

func NewJWT(user *User, secretKey string) (*JWT, error) {
	myJwt := &JWT{}
	err := myJwt.CreateToken(user.ID, secretKey)
	if err != nil {
		return nil, err
	}
	return myJwt, nil
}

func (j *JWT) CreateToken(userID int, secretKey string) error {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return err
	}
	j.Token = tokenStr
	return nil
}
