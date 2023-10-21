// jwt_util.go

package utils

import (
	"go-server/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret string

// Claims represents the JWT claims.
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	UserRole string `json:"user_role"`
	jwt.StandardClaims
}

func GetSecretKey() {
	jwtSecret = os.Getenv("JWT_SECRET")
}

func GenerateJWTToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWTToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

func ExtractClaims(r *http.Request) (jwt.MapClaims, bool) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, false
	}
	tokenString = tokenString[len("Bearer "):]

	token, err := ValidateJWTToken(tokenString)
	if err != nil || !token.Valid {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}

	return claims, true
}

// VerifyJWTToken verifies and decodes a JWT token and returns the claims if it's valid.
func VerifyJWTToken(tokenString string) (Claims, bool) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return claims, false
	}
	return claims, token.Valid
}
