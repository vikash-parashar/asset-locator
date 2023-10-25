package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// Load .env file to read email settings
func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// sendResetEmail sends a reset email to the user.
func SendResetPasswordEmail(recipientEmail, resetToken string) error {
	// Load environment variables from .env file
	loadEnv()

	// Retrieve email settings from environment variables
	emailUsername := os.Getenv("EMAIL_USERNAME")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailSMTPServer := os.Getenv("EMAIL_SMTP_SERVER")
	emailSMTPPort, err := strconv.Atoi(os.Getenv("EMAIL_SMTP_PORT"))
	if err != nil {
		log.Println("failed to get smtp prot value from config file")
	}
	emailFrom := os.Getenv("EMAIL_FROM")

	// Compose the email message.
	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "Password Reset Request")
	m.SetBody("text/html", "To reset your password, click on the following link: "+
		"https://example.com/reset-password?token="+resetToken)

	// Create an email dialer and send the email.
	d := gomail.NewDialer(emailSMTPServer, emailSMTPPort, emailUsername, emailPassword)

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
