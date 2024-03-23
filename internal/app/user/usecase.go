package user

import (
	"golang.org/x/crypto/bcrypt"
	"hackaton/internal/app/domain"
)

type UserUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) *UserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (u *UserUseCase) Login(loginDTO domain.LoginDTO) (*domain.JWT, error) {
	user, err := u.userRepository.GetUserByRegistration(loginDTO.Registration)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))
	if err != nil {
		return nil, err
	}
	jwt := domain.NewJWT(user)
	return jwt, nil
}
