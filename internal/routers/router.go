package routers

import (
	"inventory-management-api/internal/app"

	"github.com/gin-gonic/gin"
)

func SetupRouter(c *app.Container) *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")
	PublicRoutes(public, c)

	protected := r.Group("/api")
	ProtectedRoutes(protected, c)

	return r
}
