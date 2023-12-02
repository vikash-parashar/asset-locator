package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vikash-parashar/asset-locator/logger"
)

// DB represents the PostgreSQL database.
type DB struct {
	*sql.DB
}

// NewDB creates a new database connection.
func NewDB(host, port, user, password, dbName string) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.ErrorLogger.Printf("Error opening database connection: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.ErrorLogger.Printf("Error pinging database: %v", err)
		return nil, err
	}

	logger.InfoLogger.Println("Connected to the database")

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
