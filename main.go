package main

import (
	config "go-server/db"
	"log"
	"net/http"
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

	db, err := config.ConnectDB()
	if err != nil {
		log.Println("failed to connect with db")
		log.Fatalln(err)
	}

	err = config.CreateDatabaseTables(db)
	if err != nil {
		log.Println("failed to create tables into db")
		log.Fatalln(err)
	}

	defer db.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Fatalln(r.Run(":" + port))
}
