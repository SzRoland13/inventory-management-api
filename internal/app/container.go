package app

import (
	"inventory-management-api/internal/handlers"
	"inventory-management-api/internal/repositories"
	"inventory-management-api/internal/services"

	"gorm.io/gorm"
)

type Container struct {
	AuthHandler *handlers.AuthHandler
	// TODO: add all handlers here
}

func NewContainer(db *gorm.DB) *Container {
	//Repositories
	userRepository := repositories.NewUserRepository(db)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(db)

	//Services
	authService := services.NewAuthService(userRepository, refreshTokenRepository)

	//Handlers
	authHandler := handlers.NewAuthHandler(authService)

	return &Container{
		AuthHandler: authHandler,
	}
}
