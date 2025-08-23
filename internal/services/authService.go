package services

import (
	"errors"
	"inventory-management-api/internal/models"
	"inventory-management-api/internal/pkg/utils"
	"inventory-management-api/internal/repositories"
	"time"
)

type AuthService interface {
	Register(username, email, password, role string) error
	Login(email, password string) (accessToken string, refreshToken string, err error)
	RefreshAccessToken(refreshToken string) (newAccessToken string, err error)
	Logout(refreshToken string) error
}

type authService struct {
	userRepository         repositories.UserRepository
	refreshTokenRepository repositories.RefreshTokenRepository
}

func NewAuthService(userRepository repositories.UserRepository, refreshTokenRepository repositories.RefreshTokenRepository) AuthService {
	return &authService{userRepository, refreshTokenRepository}
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

func (s *authService) Login(email, password string) (accessToken string, refreshToken string, err error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}
	accessToken, _ = utils.GenerateJWT(user.ID, user.Role)

	refreshToken = utils.GenerateRefreshToken()

	rt := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = s.refreshTokenRepository.Create(rt)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) RefreshAccessToken(refreshToken string) (newAccessToken string, err error) {
	rt, err := s.refreshTokenRepository.FindByToken(refreshToken)

	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	if time.Now().After(rt.ExpiresAt) {
		_ = s.refreshTokenRepository.DeleteById(rt.ID)
		return "", errors.New("refresh token expired")
	}

	newAccessToken, _ = utils.GenerateJWT(rt.User.ID, rt.User.Role)

	return newAccessToken, nil
}

func (s *authService) Logout(refreshToken string) error {
	rt, err := s.refreshTokenRepository.FindByToken(refreshToken)

	if err != nil {
		return err
	}

	return s.refreshTokenRepository.DeleteById(rt.ID)
}
