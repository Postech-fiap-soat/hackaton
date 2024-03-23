package user

import (
	"golang.org/x/crypto/bcrypt"
	"hackaton/internal/app/domain"
	"hackaton/internal/config"
)

type UserUseCase struct {
	userRepository domain.UserRepository
	cfg            *config.Config
}

func NewUserUseCase(userRepository domain.UserRepository, cfg *config.Config) *UserUseCase {
	return &UserUseCase{userRepository: userRepository, cfg: cfg}
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
	jwt, err := domain.NewJWT(user, u.cfg.JWTSecretKey)
	if err != nil {
		return nil, err
	}
	return jwt, nil
}
