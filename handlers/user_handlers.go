// handlers.go

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

// SignUp handles the registration of a new user.
func SignUp(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse form data into a User struct
		var newUser models.User
		newUser.FirstName = c.PostForm("first_name")
		newUser.LastName = c.PostForm("last_name")
		newUser.Phone = c.PostForm("phone")
		newUser.Email = c.PostForm("email")
		newUser.Password = c.PostForm("password")
		newUser.Role = "general" // You can set the role as needed

		// Check if the user already exists with this email
		_, err := db.GetUserByEmailID(newUser.Email)
		if err == nil {
			log.Println(err)
			c.JSON(http.StatusConflict, gin.H{"status": "error", "message": "User already exists with this email"})
			return
		}

		// Hash the password
		hashedPassword, err := utils.HashPassword(newUser.Password)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to hash password"})
			return
		}
		newUser.Password = hashedPassword

		// Register the user in the database
		if err := db.RegisterUser(&newUser); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "You are now registered"})
	}
}

// Login handles the user login and returns a JWT token upon successful login.
func Login(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse request body into a User struct
		var loginRequest struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		loginRequest.Email = c.PostForm("email")
		loginRequest.Password = c.PostForm("password")

		// Check if the user exists in the database
		user, err := db.GetUserByEmailID(loginRequest.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Incorrect email or password"})

			return
		}

		// Verify the password
		if !utils.VerifyPassword(loginRequest.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Incorrect password"})
			return
		}

		// Generate a JWT token
		token, err := utils.GenerateJWTToken(user)
		if err != nil {
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

// Logout handles the user logout by clearing the JWT token cookie.
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Clear the JWT token cookie
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("jwt-token", "", -1, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Logout successful"})
	}
}

// ForgotPassword handles the process of resetting a user's forgotten password.
func ForgotPassword(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve email address from the user input
		var resetRequest struct {
			Email string `json:"email" binding:"required"`
		}
		if err := c.ShouldBindJSON(&resetRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid input data"})
			return
		}

		// Check if the user exists in the database
		user, err := db.GetUserByEmailID(resetRequest.Email)
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

// ResetPassword handles the user's password reset by verifying the reset token.
func ResetPassword(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse request body into a struct
		var resetRequest struct {
			ResetToken  string `json:"reset_token" binding:"required"`
			NewPassword string `json:"new_password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&resetRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid input data"})
			return
		}

		// Verify the reset token
		user, err := db.VerifyResetToken(resetRequest.ResetToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid or expired reset token"})
			return
		}

		// Hash the new password
		hashedPassword, err := utils.HashPassword(resetRequest.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to hash the new password"})
			return
		}

		// Update the user's password in the database
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
