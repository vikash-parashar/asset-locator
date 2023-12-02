package middleware

import (
	"net/http"

	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin" // Update with your actual import path
)

// Claims represents the JWT claims.
type Claims struct {
	UserId    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	UserRole  string `json:"user_role"`
	jwt.StandardClaims
}

// AuthMiddleware checks JWT tokens from cookies and enforces user roles.
func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the cookie
		cookie, err := c.Request.Cookie("jwt-token")
		if err != nil {
			// Token not found, redirect to login page
			logger.ErrorLogger.Printf("Token not found, redirecting to login page: %s\n", err)
			c.Redirect(http.StatusSeeOther, "http://localhost:8080")
			c.Abort()
			return
		}

		token := cookie.Value

		claims, valid := utils.VerifyJWTToken(token)
		if !valid {
			// Token is invalid or expired, redirect to login page
			logger.WarningLogger.Printf("Invalid or expired token, redirecting to login page\n")
			c.Redirect(http.StatusSeeOther, "http://localhost:8080/")
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
			logger.ErrorLogger.Printf("Access Forbidden for role: %s\n", userRole)
			c.JSON(http.StatusForbidden, gin.H{"message": "Access Forbidden"})
			c.Abort()
			return
		}

		logger.InfoLogger.Printf("User %s has access\n", claims.UserEmail)
		c.Next()
	}
}
