package db

import (
	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/models"
)

// CreateDeviceAMCOwnerDetail creates a new record in the device_amc_owner table.
func (db *DB) CreateDeviceAMCOwnerDetail(data *models.DeviceAMCOwnerDetail) error {
	query := `
        INSERT INTO device_amc_owner (serial_number, device_make_model, model, po_number, po_order_date, eosl_date, amc_start_date, amc_end_date, device_owner)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.PONumber, data.POOrderDate, data.EOSLDate, data.AMCStartDate, data.AMCEndDate, data.DeviceOwner)
	if err != nil {
		logger.ErrorLogger.Printf("Error creating DeviceAMCOwnerDetail: %v", err)
		return err
	}
	logger.InfoLogger.Println("Created DeviceAMCOwnerDetail successfully")
	return nil
}

// GetAllDeviceAMCOwnerDetail retrieves all records from the device_amc_owner table.
func (db *DB) GetAllDeviceAMCOwnerDetail() ([]models.DeviceAMCOwnerDetail, error) {
	query := "SELECT * FROM device_amc_owner"
	rows, err := db.Query(query)
	if err != nil {
		logger.ErrorLogger.Printf("Error querying DeviceAMCOwnerDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DeviceAMCOwnerDetail
	for rows.Next() {
		var data models.DeviceAMCOwnerDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.PONumber, &data.POOrderDate, &data.EOSLDate, &data.AMCStartDate, &data.AMCEndDate, &data.DeviceOwner)
		if err != nil {
			logger.ErrorLogger.Printf("Error scanning DeviceAMCOwnerDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	logger.InfoLogger.Println("Retrieved all DeviceAMCOwnerDetail records successfully")
	return results, nil
}

// UpdateDeviceAMCOwnerDetail updates an existing record in the device_amc_owner table based on the ID.
func (db *DB) UpdateDeviceAMCOwnerDetail(id int, data *models.DeviceAMCOwnerDetail) error {
	query := `
        UPDATE device_amc_owner
        SET serial_number = ?, device_make_model = ?, model = ?, po_number = ?, po_order_date = ?, eosl_date = ?, amc_start_date = ?, amc_end_date = ?, device_owner = ?
        WHERE id = ?
    `
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.PONumber, data.POOrderDate, data.EOSLDate, data.AMCStartDate, data.AMCEndDate, data.DeviceOwner, id)
	if err != nil {
		logger.ErrorLogger.Printf("Error updating DeviceAMCOwnerDetail: %v", err)
		return err
	}
	logger.InfoLogger.Printf("Updated DeviceAMCOwnerDetail with ID %d successfully", id)
	return nil
}

// DeleteDeviceAMCOwnerDetail deletes a record from the device_amc_owner table based on the ID.
func (db *DB) DeleteDeviceAMCOwnerDetail(id int) error {
	query := "DELETE FROM device_amc_owner WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		logger.ErrorLogger.Printf("Error deleting DeviceAMCOwnerDetail with ID %d: %v", id, err)
		return err
	}
	logger.InfoLogger.Printf("Deleted DeviceAMCOwnerDetail with ID %d successfully", id)
	return nil
}

// FetchDataFromTable1 retrieves data from table 1.
func (db *DB) FetchDataFromDeviceOwner() ([]*models.DeviceAMCOwnerDetail, error) {
	query := "SELECT * FROM device_amc_owner"
	rows, err := db.Query(query)
	if err != nil {
		logger.ErrorLogger.Printf("Error fetching data from table 1: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []*models.DeviceAMCOwnerDetail
	for rows.Next() {
		var data models.DeviceAMCOwnerDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.PONumber, &data.POOrderDate, &data.EOSLDate, &data.AMCStartDate, &data.AMCEndDate, &data.DeviceOwner)
		if err != nil {
			logger.ErrorLogger.Printf("Error scanning data from table 1: %v", err)
			return nil, err
		}
		results = append(results, &data)
	}

	logger.InfoLogger.Println("Fetched data from device_amc_owner table successfully")
	return results, nil
}
