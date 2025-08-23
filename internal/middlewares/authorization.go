package middlewares

import (
	"inventory-management-api/internal/domain"
	"inventory-management-api/internal/pkg/utils"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		userID, role, err := utils.ValidateJWT(strings.TrimPrefix(authHeader, "Bearer "))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
	}
}

func RequireRole(allowedRoles domain.AllowedRoles) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := domain.UserRole(c.GetString("role"))

		if role == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": domain.InvalidJWTToken})
			c.Abort()
			return
		}

		if slices.Contains(allowedRoles, role) {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, gin.H{"error": domain.AccessDenied})
		c.Abort()
	}
}
