package main

import (
	"go-server/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")

	// db, err := config.ConnectDB()
	// if err != nil {
	// 	log.Println("failed to connect with db")
	// 	log.Fatalln(err)
	// }

	// err = config.CreateDatabaseTables(db)
	// if err != nil {
	// 	log.Println("failed to create tables into db")
	// 	log.Fatalln(err)
	// }

	// defer db.Close()

	r := gin.Default()

	r.GET("/", handlers.Dashboard)
	r.GET("/api/location-details", handlers.GetLocationDetails)
	r.GET("/api/owner-details", handlers.GetOwnerDetails)
	r.GET("/api/power-details", handlers.GetPowerDetails)
	r.GET("/api/fiber-details", handlers.GetFiberDetails)
	r.POST("/api/location-details", handlers.CreateNewLocationDetails)
	r.POST("/api/owner-details", handlers.CreateNewOwnerDetails)
	r.POST("/api/power-details", handlers.CreateNewPowerDetails)
	r.POST("/api/fiber-details", handlers.CreateNewFiberDetails)

	log.Fatalln(r.Run(":" + port))
}
