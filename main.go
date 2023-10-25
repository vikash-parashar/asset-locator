package main

import (
	"log"
	"os"

	"go-server/db"
	"go-server/handlers"
	"go-server/middleware"

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

	r.GET("/health-check", handlers.HealthCheck)

	r.GET("/", handlers.RenderIndexPage)
	r.GET("/forget-password-page", handlers.RenderForgotPasswordPage)

	r.POST("/signup", handlers.SignUp(db))
	r.POST("/login", handlers.Login(db))
	r.POST("/logout", handlers.Logout())
	r.POST("/forget-password", handlers.ForgetPassword(db))
	r.POST("/reset-password", handlers.ResetPassword(db))

	//protected routes

	r.GET("/api/v1/home", middleware.AuthMiddleware("admin", "general"), handlers.RenderHomePage(db))

	r.GET("/api/v1/location-details", middleware.AuthMiddleware("admin", "general"), handlers.GetLocationDetails(db))
	r.POST("/api/v1/location-details", middleware.AuthMiddleware("admin"), handlers.CreateNewLocationDetails(db))
	r.PUT("/api/v1/location-details/:id", middleware.AuthMiddleware("admin"), handlers.UpdateDeviceLocationDetail(db))
	r.DELETE("/api/v1/location-details/:id", middleware.AuthMiddleware("admin"), handlers.DeleteDeviceLocationDetail(db))
	r.GET("/api/v1/location-details/pdf", middleware.AuthMiddleware("admin"), handlers.DownloadDeviceLocationDetailPDF(db))
	r.GET("/api/v1/location-details/excel", middleware.AuthMiddleware("admin"), handlers.DownloadDeviceLocationDetail(db))

	r.GET("/api/v1/owner-details", middleware.AuthMiddleware("admin", "general"), handlers.GetOwnerDetails(db))
	r.POST("/api/v1/owner-details", middleware.AuthMiddleware("admin"), handlers.CreateNewOwnerDetails(db))
	r.PUT("/api/v1/owner-details/:id", middleware.AuthMiddleware("admin"), handlers.UpdateDeviceAMCOwnerDetail(db))
	r.DELETE("/api/v1/owner-details/:id", middleware.AuthMiddleware("admin"), handlers.DeleteDeviceAMCOwnerDetail(db))
	r.GET("/api/v1/owner-details/pdf", middleware.AuthMiddleware("admin"), handlers.DownloadDeviceAMCOwnerDetailPDF(db))
	r.GET("/api/v1/owner-details/excel", middleware.AuthMiddleware("admin"), handlers.DownloadDeviceAMCOwnerDetail(db))

	r.GET("/api/v1/power-details", middleware.AuthMiddleware("admin", "general"), handlers.GetPowerDetails(db))
	r.POST("/api/v1/power-details", middleware.AuthMiddleware("admin"), handlers.CreateNewPowerDetails(db))
	r.PUT("/api/v1/power-details/:id", middleware.AuthMiddleware("admin"), handlers.UpdateDevicePowerDetail(db))
	r.DELETE("/api/v1/power-details/:id", middleware.AuthMiddleware("admin"), handlers.DeleteDevicePowerDetail(db))
	r.GET("/api/v1/power-details/pdf", middleware.AuthMiddleware("admin"), handlers.DownloadDevicePowerDetailPDF(db))
	r.GET("/api/v1/power-details/excel", middleware.AuthMiddleware("admin"), handlers.DownloadDevicePowerDetail(db))

	r.GET("/api/v1/fiber-details", middleware.AuthMiddleware("admin", "general"), handlers.GetFiberDetails(db))
	r.POST("/api/v1/fiber-details", middleware.AuthMiddleware("admin"), handlers.CreateNewFiberDetails(db))
	r.PUT("/api/v1/fiber-details/:id", middleware.AuthMiddleware("admin"), handlers.UpdateDeviceEthernetFiberDetail(db))
	r.DELETE("/api/v1/fiber-details/:id", middleware.AuthMiddleware("admin"), handlers.DeleteDeviceEthernetFiberDetail(db))
	r.GET("/api/v1/fiber-details/pdf", middleware.AuthMiddleware("admin"), handlers.DownloadDeviceEthernetFiberDetailPDF(db))
	r.GET("/api/v1/fiber-details/excel", middleware.AuthMiddleware("admin"), handlers.DownloadDeviceEthernetFiberDetail(db))

	log.Fatalln(r.Run(":" + port))
}
