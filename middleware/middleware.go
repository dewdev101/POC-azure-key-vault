package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthorizationMiddleware extracts the Authorization header and removes the "Bearer " prefix
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		// Set token in context for use in handlers
		c.Set("token", token)
		// Add Content-Type header

		c.Next()
	}
}
