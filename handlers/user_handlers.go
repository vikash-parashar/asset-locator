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

// ForgotPasswordHandler handles the process of resetting a user's forgotten password.
func ForgotPassword(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve email address from the user input
		email := c.PostForm("email")

		// Check if the user exists in the database
		user, err := db.GetUserByEmailID(email)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "User not found"})
			return
		}

		// Generate a unique reset token and set an expiration time for it (e.g., 1 hour)
		resetToken, err := utils.GeneratePasswordResetToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to generate reset token"})
			return
		}

		// Save the reset token in the database associated with the user's account
		if err := db.SetResetToken(int(user.ID), resetToken); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to save reset token"})
			return
		}

		// Send an email to the user with a link to reset their password
		err = utils.SendResetPasswordEmail(user.Email, resetToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to send reset email"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Reset instructions sent to your email"})
	}
}

// ResetPasswordHandler handles the user's password reset by verifying the reset token.
func ResetPassword(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve reset token and new password from the user input
		resetToken := c.PostForm("reset_token")
		newPassword := c.PostForm("new_password")

		// Verify the reset token
		user, err := db.VerifyResetToken(resetToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid or expired reset token"})
			return
		}

		// Hash the new password
		hashedPassword, err := utils.HashPassword(newPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to hash the new password"})
			return
		}

		// Update the user's password in the database
		user.Password = hashedPassword
		if err := db.UpdateUserPassword(int(user.ID), hashedPassword); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to update the password"})
			return
		}

		// Clear the reset token from the database
		if err := db.ClearResetToken(int(user.ID)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to clear the reset token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Password reset successful"})
	}
}
