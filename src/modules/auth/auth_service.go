package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	commons "go-rest-api-boilerplate/src/commons/config"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type IAuthService interface {
	Register(registerRequest SRegisterRequest) (User, error)
	Login(loginRequest SLoginRequest) (User, string, error)
}

type SAuthService struct {
	authRepository IAuthRepository
}

func NewAuthService(authRepository IAuthRepository) *SAuthService {
	return &SAuthService{authRepository}
}

func (s *SAuthService) Register(authRequest SRegisterRequest) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		Name:     authRequest.Name,
		Username: authRequest.Username,
		Email:    authRequest.Email,
		Password: string(hashedPassword),
	}
	createdUser, err := s.authRepository.Create(user)
	return createdUser, err
}

func (s *SAuthService) Login(loginRequest SLoginRequest) (User, string, error) {
	user, err := s.authRepository.FindByUsername(loginRequest.Username)
	if err != nil {
		return User{}, "", fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return User{}, "", fmt.Errorf("wrong password")
	}

	token, err := generateToken(user.ID, user.Username)
	if err != nil {
		return User{}, "", err
	}

	return user, token, err
}

func generateToken(ID int, username string) (string, error) {
	key := commons.LoadConfig().JwtSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      ID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(), // 1month
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}
