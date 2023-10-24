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

func init() {
	GetSecretKey()
}

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
	claims := Claims{
		UserID:   user.Id,
		Username: user.Username,
		UserRole: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
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

func VerifyJWTToken(tokenString string) (Claims, bool) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return claims, false
	}
	return claims, token.Valid
}
