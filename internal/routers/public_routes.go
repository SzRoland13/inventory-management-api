package routers

import (
	"inventory-management-api/internal/app"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(r *gin.RouterGroup, c *app.Container) {
	r.POST("/login", c.AuthHandler.Login)
	r.GET("/ping", c.AuthHandler.Ping)
}
