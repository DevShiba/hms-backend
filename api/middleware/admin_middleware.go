package middleware

import (
	"hms-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleFromCtx, exists := c.Get("x-user-role")
		if !exists {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User role not found in context"})
			c.Abort()
			return
		}

		userRole, ok := roleFromCtx.(domain.UserRole)
		if !ok {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Invalid user role type in context"})
			c.Abort()
			return
		}

		if userRole != domain.AdminRole {
			c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "Access denied. Admin role required."})
			c.Abort()
			return
		}
		c.Next()
	}
}