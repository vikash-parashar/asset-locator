package middleware

import (
	"go-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Get Token From Auth Headers

// Define an authentication middleware function that checks JWT tokens and user roles
// func AuthMiddleware(roles ...string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		claims, valid := utils.ExtractClaims(c.Request)
// 		if !valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}

// 		// Check if the user has the required role
// 		requiredRoles := make(map[string]bool)
// 		for _, role := range roles {
// 			requiredRoles[role] = true
// 		}

// 		userRole, ok := claims["user_role"].(string)
// 		if !ok || !requiredRoles[userRole] {
// 			c.JSON(http.StatusForbidden, gin.H{"message": "Access Forbidden"})
// 			c.Abort()
// 			return
// 		}

//			c.Next()
//		}
//	}

// TODO: AuthMiddleware checks JWT tokens and user roles from cookies
func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the cookie
		cookie, err := c.Request.Cookie("jwt-token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			// If JWT is not present, redirect to the "/" path
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		token := cookie.Value

		claims, valid := utils.VerifyJWTToken(token)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the user has the required role
		hasRequiredRole := false
		userRole := claims.UserRole // Access the user role from the claims

		for _, role := range roles {
			if userRole == role {
				hasRequiredRole = true
				break
			}
		}

		if !hasRequiredRole {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
