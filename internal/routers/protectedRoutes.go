package routers

import (
	"inventory-management-api/internal/app"
	"inventory-management-api/internal/domain"
	"inventory-management-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func ProtectedRoutes(r *gin.RouterGroup, c *app.Container) {
	r.Use(middlewares.JWTAuthMiddleware())

	r.POST("/register", middlewares.RequireRole(domain.RolesAdminOnly), c.AuthHandler.Register)
}
