package main

import (
	"html/template"

	"github.com/vikash-parashar/asset-locator/config"
	"github.com/vikash-parashar/asset-locator/db"
	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// load's environment variables values from .env file
func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		logger.ErrorLogger.Printf("Error loading .env file: %v", err)
	}
}

// main function
func main() {
	loadEnvVariables()

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the database connection
	dbConn, err := db.NewDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		logger.ErrorLogger.Printf("Error connecting to the database: %v", err)
	}
	defer dbConn.Close()

	// Setting server mux as default mux
	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Define a custom template function
	r.SetFuncMap(template.FuncMap{
		"add1": func(i int) int {
			return i + 1
		},
	})

	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	// Set up routes from the routes package
	routes.SetupRoutes(r, dbConn)

	logger.ErrorLogger.Println(r.Run(":" + cfg.Port))
}
