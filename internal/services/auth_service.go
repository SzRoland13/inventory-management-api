package services

import (
	"errors"
	"inventory-management-api/internal/models"
	"inventory-management-api/internal/pkg/utils"
	"inventory-management-api/internal/repositories"
)

type AuthService interface {
	Register(username, email, password, role string) error
	Login(email, password string) (string, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{userRepository}
}

func (s *authService) Register(username, email, password, role string) error {
	hashed, _ := utils.HashPassword(password)
	return s.userRepository.Create(&models.User{
		Username: username,
		Email:    email,
		Password: hashed,
		Role:     role,
	})
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}
	token, _ := utils.GenerateJWT(user.ID, user.Role)
	return token, nil
}
