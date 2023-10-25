// auth_middleware.go

package middleware

import (
	"go-server/utils"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret string

func init() {
	GetSecretKey()
}

// Claims represents the JWT claims.
type Claims struct {
	UserId    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	UserRole  string `json:"user_role"`
	jwt.StandardClaims
}

func GetSecretKey() {
	jwtSecret = os.Getenv("JWT_SECRET")
}

// AuthMiddleware checks JWT tokens from cookies and enforces user roles.
func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the cookie
		cookie, err := c.Request.Cookie("jwt-token")
		if err != nil {
			// Token not found, redirect to login page
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}

		token := cookie.Value

		claims, valid := utils.VerifyJWTToken(token)
		if !valid {
			// Token is invalid or expired, redirect to login page
			c.Redirect(http.StatusSeeOther, "/")
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

// func validateJWTToken(token string) (*Claims, error) {
// 	claims := &Claims{}
// 	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(jwtSecret), nil
// 	})
// 	return claims, err
// }

// func hasRequiredRole(userRole string, roles []string) bool {
// 	for _, role := range roles {
// 		if userRole == role {
// 			return true
// 		}
// 	}
// 	return false
// }

// func handleUnauthorized(c *gin.Context) {
// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
// 	c.Abort()
// }

// func handleAccessForbidden(c *gin.Context) {
// 	c.JSON(http.StatusForbidden, gin.H{"message": "Access Forbidden"})
// 	c.Abort()
// }
