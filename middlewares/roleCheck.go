package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleCheck(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("userRole") // Assuming the role is set in the JWT middleware
		for _, role := range allowedRoles {
			if userRole == role {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}
