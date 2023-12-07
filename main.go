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

	// Initialize the MySQL database connection
	dbConn, err := db.NewMySQLDB(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err != nil {
		logger.ErrorLogger.Printf("Error connecting to the MySQL database: %v", err)
		return
	}
	defer dbConn.Close()

	logger.InfoLogger.Println("Connection to MySQL database is successful!")

	// Create database tables
	if err := db.CreateDatabaseTables(dbConn.DB); err != nil {
		logger.ErrorLogger.Fatalln("Failed to create tables in the database")
	}
	logger.InfoLogger.Println("Tables are created in the database successfully!")

	// File path to the insertdata.mysql file
	filePath := ".db/schema/insertdata.mysql"

	// Insert data from the file into the database
	if err := db.InsertDataFromFile(dbConn.DB, filePath); err != nil {
		logger.ErrorLogger.Printf("Error inserting dummy data into the database: %v", err)
		return
	}
	logger.InfoLogger.Println("Dummy data inserted successfully")

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
