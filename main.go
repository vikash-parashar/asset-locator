package main

import (
	"html/template"
	"log"
	"os"

	"github.com/vikash-parashar/asset-locator/db"
	"github.com/vikash-parashar/asset-locator/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	loadEnvVariables()

	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Initialize the database connection
	dbConn, err := db.NewDB(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbConn.Close()

	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/static", "./static")
	//  Define a custom template function
	r.SetFuncMap(template.FuncMap{
		"add1": func(i int) int {
			return i + 1
		},
	})
	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")
	// Set up routes from the routes package
	routes.SetupRoutes(r, dbConn)

	log.Fatal(r.Run(":" + port))

}
