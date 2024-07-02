package services

import (
	"go-gin-productManagerPro/models"
	"go-gin-productManagerPro/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(email string, password string) error
	Login(email string, password string) (*string, error)
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
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}

func (s *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil, err
	}

	token, err := CreateToken(foundUser.ID, foundUser.ID)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func CreateToken(userId uint, email uint) (*string, error) {
	// HS256でトークン情報をハッシュ化
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
