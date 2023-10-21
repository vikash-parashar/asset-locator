package main

import (
	"log"
	"os"

	"go-server/db"
	"go-server/handlers"

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
	db, err := db.NewDB(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", handlers.Dashboard)

	//TODO: Get Handlers
	r.GET("/api/v1/location-details", handlers.GetLocationDetails(db))
	r.GET("/api/v1/owner-details", handlers.GetOwnerDetails(db))
	r.GET("/api/v1/power-details", handlers.GetPowerDetails(db))
	r.GET("/api/v1/fiber-details", handlers.GetFiberDetails(db))

	//TODO: Post Handlers
	r.POST("/api/v1/location-details", handlers.CreateNewLocationDetails(db))
	r.POST("/api/v1/owner-details", handlers.CreateNewOwnerDetails(db))
	r.POST("/api/v1/power-details", handlers.CreateNewPowerDetails(db))
	r.POST("/api/v1/fiber-details", handlers.CreateNewFiberDetails(db))

	//TODO: Delete Handlers
	r.DELETE("/api/v1/fiber-details/:id", handlers.DeleteDeviceEthernetFiberDetail(db))
	r.DELETE("/api/v1/power-details/:id", handlers.DeleteDevicePowerDetail(db))
	r.DELETE("/api/v1/owner-details/:id", handlers.DeleteDeviceAMCOwnerDetail(db))
	r.DELETE("/api/v1/location-details/:id", handlers.DeleteDeviceLocationDetail(db))

	//TODO: Update Handlers
	r.PUT("/api/v1/location-details/:id", handlers.UpdateDeviceLocationDetail(db))
	r.PUT("/api/v1/owner-details/:id", handlers.UpdateDeviceAMCOwnerDetail(db))
	r.PUT("/api/v1/power-details/:id", handlers.UpdateDevicePowerDetail(db))
	r.PUT("/api/v1/fiber-details/:id", handlers.UpdateDeviceEthernetFiberDetail(db))

	log.Fatalln(r.Run(":" + port))
}
