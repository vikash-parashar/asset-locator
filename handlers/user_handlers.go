package handlers

import (
	"go-server/db"
	"go-server/models"
	"go-server/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterUser handles the registration of a new user.
func SignUp(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve form data
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		role := c.PostForm("role")

		if _, err := db.GetUserByEmailID(email); err == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "user already exists with this email"})
			return
		}

		// Hash the password
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "failed to hash password"})
			return
		}

		// Create a new user with the hashed password
		newUser := models.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Role:      role,
			Password:  hashedPassword,
		}

		// Register the user in the database
		if err := db.RegisterUser(&newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "you are now registered"})
	}
}

// LoginUserHandler handles the user login and returns a JWT token upon successful login.
func Login(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve form data
		email := c.PostForm("email")
		password := c.PostForm("password")
		log.Println(email)
		log.Println(password)
		// Check if the user exists in the database
		user, err := db.GetUserByEmailID(email)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User not found"})
			return
		}

		// Verify the password
		if !utils.VerifyPassword(password, user.Password) {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Incorrect password"})
			return
		}

		// Generate a JWT token
		token, err := utils.GenerateJWTToken(user)
		if err != nil {
			log.Println("Failed to generate JWT token")
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to generate JWT token"})
			return
		}

		cookie := http.Cookie{
			Name:    "jwt-token",
			Value:   token,
			Expires: time.Now().Add(5 * time.Minute),
		}
		http.SetCookie(c.Writer, &cookie)
		c.JSON(http.StatusOK, gin.H{"success": true, "token": token, "message": "Login successful"})
	}
}

// LogoutUserHandler handles the user logout by clearing the JWT token cookie.
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Clear the JWT token cookie
		cookie := http.Cookie{
			Name:    "jwt-token",
			Value:   "",
			Expires: time.Unix(0, 0),
		}
		http.SetCookie(c.Writer, &cookie)
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Logout successful"})
	}
}

// ForgetPassword handles the "forget password" feature.
func ForgetPassword(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve form data (assuming you use the user's email for password reset)
		email := c.PostForm("email")

		// Check if the user with the provided email exists in the database
		user, err := db.GetUserByEmailID(email)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "User not found"})
			return
		}

		// Generate a password reset token (you can implement this function)
		_, err = utils.GeneratePasswordResetToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to generate password reset token"})
			return
		}

		// Send the password reset link to the user (you can implement this function)
		// You might want to send an email with a link that includes the resetToken
		// and leads the user to a reset password page.

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Password reset instructions sent to your email"})
	}
}

// ResetPassword handles the password reset feature.
func ResetPassword(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		resetToken := c.PostForm("reset_token")
		newPassword := c.PostForm("new_password")

		// Check if the token exists in your database
		user, err := db.GetUserByResetToken(resetToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid or expired reset token"})
			return
		}

		// Check if the token has expired (you should store an expiration time in your database)
		if time.Now().After(user.ResetTokenExpiry) {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Reset token has expired"})
			return
		}

		// Reset token is valid; update the user's password with the new password
		hashedPassword, err := utils.HashPassword(newPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to hash the new password"})
			return
		}

		user.Password = hashedPassword

		// Clear the reset token and its expiration time after a successful reset
		user.ResetToken = ""
		user.ResetTokenExpiry = time.Time{}

		// Update the user in the database
		if err := db.UpdateUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to update the user's password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Password reset successful"})
	}
}
