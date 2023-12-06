package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/vikash-parashar/asset-locator/config"
	"github.com/vikash-parashar/asset-locator/db"
	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/routes"

	"github.com/gin-gonic/gin"
)

// automate db related operations like create db/table or drop db/table etc.
func runMigration(dbConn *db.DB, migrationType string) error {
	var migrationFile string

	switch migrationType {
	case "up":
		migrationFile = "./db/schema.up.sql"
	case "down":
		migrationFile = "./db/schema.down.sql"
	default:
		return fmt.Errorf("unsupported migration type: %s", migrationType)
	}

	content, err := os.ReadFile(migrationFile)
	if err != nil {
		return err
	}

	_, err = dbConn.Exec(string(content))
	return err
}

// main function
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the database connection
	dbConn, err := db.NewDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		logger.ErrorLogger.Printf("Error connecting to the database: %v", err)
	}
	defer dbConn.Close()

	// Parse command-line arguments
	migrateType := flag.String("migrate", "", "Specify migration type: up or down")
	flag.Parse()

	// Run migration if the -migrate flag is provided
	if *migrateType != "" {
		fmt.Printf("Running %s migration...\n", *migrateType)
		if err := runMigration(dbConn, *migrateType); err != nil {
			logger.ErrorLogger.Printf("Error running migration: %v", err)
		}
		fmt.Printf("%s migration completed.\n", *migrateType)
		return
	}

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

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

	// Log only critical errors in production
	if cfg.Env == "production" {
		r.Use(gin.Recovery())
	}
	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	// Set up routes from the routes package
	routes.SetupRoutes(r, dbConn)

	// Run the server with HTTPS if configured
	if cfg.UseHTTPS {
		err := r.RunTLS(":"+cfg.Port, cfg.CertFile, cfg.KeyFile)
		if err != nil {
			logger.ErrorLogger.Printf("Error running the server: %v", err)
		}
		log.Println("server is running")
	} else {
		err := r.Run(":" + cfg.Port)
		if err != nil {
			logger.ErrorLogger.Printf("Error running the server: %v", err)
		}
		logger.InfoLogger.Println("server is running")
	}
}
