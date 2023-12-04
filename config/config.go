package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config struct to hold configuration parameters
type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	Port           string
	JWTSecret      string
	EmailPassword  string
	EmailUsername  string
	UseHTTPS       bool
	CertFile       string
	KeyFile        string
	ExternalServer string
	ExternalPort   int
	ExternalUser   string
	ExternalPass   string
	Env            string // New field to store the environment name
}

// LoadConfig loads configuration from environment variables and a specific config file
func LoadConfig() *Config {
	// Default to development environment
	env := getEnv("APP_ENV", "development")

	// Load environment-specific configuration
	configFile := ".env." + env
	if env != "production" {
		// Load additional configuration for non-production environments
		err := godotenv.Load(configFile)
		if err != nil {
			// Log or handle the error if needed
			println("Error loading additional configuration:", err)
		}
	}

	return &Config{
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "asset-locator"),
		Port:           getEnv("PORT", "8080"),
		JWTSecret:      getEnv("JWT_SECRET", "go-server-secret"),
		EmailPassword:  getEnv("EMAIL_PASSWORD", ""),
		EmailUsername:  getEnv("EMAIL_USERNAME", ""),
		UseHTTPS:       getEnvAsBool("USE_HTTPS", false),
		CertFile:       getEnv("CERT_FILE", ""),
		KeyFile:        getEnv("KEY_FILE", ""),
		ExternalServer: getEnv("S_SERVER", ""),
		ExternalPort:   getEnvAsInt("S_PORT", 0),
		ExternalUser:   getEnv("S_USER", ""),
		ExternalPass:   getEnv("S_PASS", ""),
		Env:            env,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}
