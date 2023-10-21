package middleware

import (
	"go-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define an authentication middleware function that checks JWT tokens and user roles
func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, valid := utils.ExtractClaims(c.Request)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the user has the required role
		requiredRoles := make(map[string]bool)
		for _, role := range roles {
			requiredRoles[role] = true
		}

		userRole, ok := claims["user_role"].(string)
		if !ok || !requiredRoles[userRole] {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
