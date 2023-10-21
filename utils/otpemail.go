package utils

/*
package main

import (
    "fmt"
    "math/rand"
    "net/smtp"
    "strconv"
    "strings"
    "time"
)

// Email configuration
const (
    smtpServer   = "smtp.example.com"
    smtpPort     = "587"
    smtpUsername = "your_username"
    smtpPassword = "your_password"
)

// OTP length and validity period
const (
    otpLength         = 6
    otpValidityPeriod = 5 * time.Minute
)

// Store OTPs and their creation time
var otpStore = make(map[string]otpData)

type otpData struct {
    code       string
    createTime time.Time
}

// Generate a random OTP of specified length
func generateOTP(length int) string {
    rand.Seed(time.Now().UnixNano())
    characters := "0123456789"
    otp := make([]byte, length)
    for i := range otp {
        otp[i] = characters[rand.Intn(len(characters))]
    }
    return string(otp)
}

// Send OTP via email
func sendOTPEmail(email, otp string) error {
    from := "your_email@example.com"
    to := []string{email}
    subject := "Your OTP Code"
    body := "Your OTP code is: " + otp

    message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, strings.Join(to, ","), subject, body)

    auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)

    err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, []byte(message))
    if err != nil {
        return err
    }

    return nil
}

// Generate and send OTP to the user
func sendOTPToUser(email string) (string, error) {
    otp := generateOTP(otpLength)
    if err := sendOTPEmail(email, otp); err != nil {
        return "", err
    }

    otpStore[email] = otpData{
        code:       otp,
        createTime: time.Now(),
    }

    return otp, nil
}

// Verify OTP
func verifyOTP(email, code string) bool {
    otpData, found := otpStore[email]
    if !found {
        return false
    }

    // Check if the OTP is expired
    if time.Since(otpData.createTime) > otpValidityPeriod {
        delete(otpStore, email)
        return false
    }

    return otpData.code == code
}

// Resend OTP
func resendOTP(email string) (string, error) {
    // Check if an OTP already exists
    if _, found := otpStore[email]; found {
        return "", fmt.Errorf("An OTP for this email is already in use")
    }

    return sendOTPToUser(email)
}

func main() {
    // Example usage
    email := "user@example.com"

    // Send OTP to the user
    otp, err := sendOTPToUser(email)
    if err != nil {
        fmt.Println("Error sending OTP:", err)
        return
    }

    fmt.Println("OTP sent successfully")

    // Simulate OTP verification
    userEnteredOTP := "123456"
    if verifyOTP(email, userEnteredOTP) {
        fmt.Println("OTP is valid")
    } else {
        fmt.Println("Invalid OTP")
    }

    // Resend OTP
    newOTP, err := resendOTP(email)
    if err != nil {
        fmt.Println("Error resending OTP:", err)
    } else {
        fmt.Println("Resent OTP:", newOTP)
    }
}

*/
