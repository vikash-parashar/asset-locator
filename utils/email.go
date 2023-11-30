package utils

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

// SendResetPasswordEmail sends a reset email to the user using Gmail SMTP.
func SendResetPasswordEmail(recipientEmail, resetToken string) error {

	// Retrieve email settings from environment variables
	emailUsername := os.Getenv("EMAIL_USERNAME")
	emailPassword := os.Getenv("EMAIL_PASSWORD")

	// Gmail SMTP server and port with TLS
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	// Set up authentication
	auth := smtp.PlainAuth("", emailUsername, emailPassword, smtpServer)

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		ServerName: smtpServer, // Specify the server name
	}

	// Connect to the SMTP server
	client, err := smtp.Dial(smtpServer + ":" + strconv.Itoa(smtpPort))
	if err != nil {
		log.Println("Failed to connect to SMTP server:", err)
		return err
	}
	defer client.Close()

	// Upgrade the connection to use TLS
	if err := client.StartTLS(tlsConfig); err != nil {
		log.Println("Failed to start TLS:", err)
		return err
	}

	// Authenticate and send the email
	if err := client.Auth(auth); err != nil {
		log.Println("SMTP authentication failed:", err)
		return err
	}

	if err := client.Mail(emailUsername); err != nil {
		log.Println("Failed to set sender:", err)
		return err
	}

	if err := client.Rcpt(recipientEmail); err != nil {
		log.Println("Failed to set recipient:", err)
		return err
	}

	wc, err := client.Data()
	if err != nil {
		log.Println("Failed to open data connection:", err)
		return err
	}
	defer wc.Close()

	message := "To: " + recipientEmail + "\r\n" +
		"<h2>Subject: Password Reset Request</h2>\r\n" +
		"\r\n" +
		"<p>To reset your password, click on the following link:</p>\r\n" +
		"http://localhost:8080/reset-password?token=" + resetToken

	_, err = wc.Write([]byte(message))
	if err != nil {
		log.Println("Failed to send email data:", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}
