package utils

import (
	"errors"
	"inventory-management-api/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your-secret-key")

func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (uint, string, error) {
	if tokenString == "" {
		return 0, "", errors.New(domain.MissingJWTToken)
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(domain.InvalidSigningMethod)
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, "", errors.New(domain.InvalidJWTToken)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		role, _ := claims["role"].(string)

		var userID uint
		if uid, ok := claims["user_id"].(float64); ok {
			userID = uint(uid)
		}

		return userID, role, nil
	}

	return 0, "", errors.New(domain.InvalidJWTToken)
}
