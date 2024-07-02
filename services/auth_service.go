package services

import (
	"go-gin-productManagerPro/models"
	"go-gin-productManagerPro/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(email string, password string) error
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthSevice(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Signup(email string, password string) error {
	// passwordのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email: email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}
