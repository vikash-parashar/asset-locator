package config

import "os"

// Config struct holds the configuration variables
type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	Port          string
	JWTSecret     string
	EmailPassword string
	EmailUsername string
	SServer       string
	SPort         string
	SUser         string
	SPass         string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		Port:          os.Getenv("PORT"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),
		EmailUsername: os.Getenv("EMAIL_USERNAME"),
		SServer:       os.Getenv("S_SERVER"),
		SPort:         os.Getenv("S_PORT"),
		SUser:         os.Getenv("S_USER"),
		SPass:         os.Getenv("S_PASS"),
	}
}
