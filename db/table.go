package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/vikash-parashar/asset-locator/logger" // Import the logger package
)

func CreateDatabaseTables(db *sql.DB) error {

	// Create the tables
	createTables := `
		CREATE TABLE IF NOT EXISTS device_location_detail (
			serial_number VARCHAR(255) PRIMARY KEY,
			device_make_model VARCHAR(255),
			model VARCHAR(255),
			device_type VARCHAR(255),
			data_center VARCHAR(255),
			region VARCHAR(255),
			dc_location VARCHAR(255),
			device_location VARCHAR(255),
			device_row_number INT,
			device_rack_number INT,
			device_ru_number VARCHAR(255)
		);
		
		CREATE TABLE IF NOT EXISTS device_amc_owner_detail (
			serial_number VARCHAR(255) PRIMARY KEY,
			device_make_model VARCHAR(255),
			model VARCHAR(255),
			po_number VARCHAR(255),
			po_order_date DATE,
			eosl_date DATE,
			amc_start_date DATE,
			amc_end_date DATE,
			device_owner VARCHAR(255)
		);
		
		CREATE TABLE IF NOT EXISTS device_power_detail (
			serial_number VARCHAR(255) PRIMARY KEY,
			device_make_model VARCHAR(255),
			model VARCHAR(255),
			device_type VARCHAR(255),
			total_power_watt INT,
			total_btu DOUBLE PRECISION,
			total_power_cable INT,
			power_socket_type VARCHAR(255)
		);
		
		CREATE TABLE IF NOT EXISTS device_ethernet_fiber_detail (
			serial_number VARCHAR(255) PRIMARY KEY,
			device_make_model VARCHAR(255),
			model VARCHAR(255),
			device_type VARCHAR(255),
			device_physical_port VARCHAR(255),
			device_port_type VARCHAR(255),
			device_port_macwwn VARCHAR(255),
			connected_device_port VARCHAR(255)
		);
		
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(255),
			reset_token VARCHAR(255),
			reset_token_expiry TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`
	res, err := db.Exec(createTables)
	if err != nil {
		logger.ErrorLogger.Printf("Error creating database tables: %v", err)
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		logger.ErrorLogger.Printf("Error getting rows affected: %v", err)
		return err
	}
	logger.InfoLogger.Printf("Database tables created successfully. Rows Affected: %d", count)

	return nil
}

// InsertDataFromFile reads SQL statements from a file and executes them to insert data into the database.
func InsertDataFromFile(db *sql.DB, filePath string) error {
	// Read SQL statements from the file
	sqlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %v", err)
	}

	// Split SQL statements based on semicolons
	sqlStatements := strings.Split(string(sqlContent), ";")

	// Execute each SQL statement
	for _, statement := range sqlStatements {
		trimmedStatement := strings.TrimSpace(statement)
		if trimmedStatement != "" {
			_, err := db.Exec(trimmedStatement)
			if err != nil {
				logger.ErrorLogger.Printf("Error executing SQL statement: %v", err)
				return fmt.Errorf("error executing SQL statement: %v", err)
			}
		}
	}

	return nil
}
