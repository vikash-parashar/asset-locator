package utils

/*

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
)

func main() {
	r := gin.Default()

	r.POST("/send-otp", SendOTP)
	r.POST("/verify-otp", VerifyOTP)

	// Replace with your Twilio credentials
	twilioAccountSID := "your_twilio_account_sid"
	twilioAuthToken := "your_twilio_auth_token"

	// Set up the Twilio client
	twilioClient := twilio.NewClient(twilioAccountSID, twilioAuthToken, nil)

	r.Run(":8080")
}

// Generate a random OTP of 6 digits
func generateRandomOTP() string {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

// SendOTP generates and sends an OTP to a given phone number
func SendOTP(c *gin.Context) {
	phoneNumber := c.PostForm("phone")

	otp := generateRandomOTP()

	// Replace with your Twilio phone number
	fromNumber := "your_twilio_phone_number"

	// Use the Twilio client to send the OTP
	message := "Your OTP is: " + otp
	_, err := twilioClient.Messages.SendMessage(fromNumber, phoneNumber, message, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

// VerifyOTP compares the received OTP with the expected OTP
func VerifyOTP(c *gin.Context) {
	receivedOTP := c.PostForm("otp")
	expectedOTP := c.PostForm("expected_otp")

	if receivedOTP == expectedOTP {
		c.JSON(http.StatusOK, gin.H{"message": "OTP is valid"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
	}
}


*/
