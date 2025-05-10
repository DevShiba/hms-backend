package middleware

import (
	"hms-api/domain"
	tokenutil "hms-api/internal"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userIDString, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}

				parsedUserID, err := uuid.Parse(userIDString)
				if err != nil {
					log.Printf("[ERROR] Middleware: Failed to parse userID string from token as UUID: %v, userIDString from token: %s\n", err, userIDString)
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid user ID format in token"})
					c.Abort()
					return
				}

				c.Set("x-user-id", parsedUserID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
