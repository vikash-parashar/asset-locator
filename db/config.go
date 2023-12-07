package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vikash-parashar/asset-locator/logger"
)

// DB represents the MySQL database.
type DB struct {
	*sql.DB
}

// NewMySQLDB creates a new database connection.
func NewMySQLDB(dbUser, dbPassword, dbHost, dbPort, dbName string) (*DB, error) {

	// Create the MySQL connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Ping the database to test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	fmt.Println("Connected to the local database!")

	return &DB{db}, nil
}

// Close closes the database connection.
func (db *DB) Close() {
	if err := db.DB.Close(); err != nil {
		logger.ErrorLogger.Printf("Error closing database connection: %v", err)
	} else {
		logger.InfoLogger.Println("Closed database connection")
	}
}
