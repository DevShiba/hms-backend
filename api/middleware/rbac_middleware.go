package middleware

import (
	"hms-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware(allowedRoles ...domain.UserRole) gin.HandlerFunc {
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

			for _, role := range allowedRoles {
					if userRole == role {
							c.Next()
							return
					}
			}

			c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "Access denied. You don't have permission to access this resource."})
			c.Abort()
	}
}