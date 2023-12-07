package db

import (
	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/models"
)

// CreateDevicePowerDetail creates a new record in the DevicePowerDetail table.
func (db *DB) CreateDevicePowerDetail(data *models.DevicePowerDetail) error {
	query := `
		INSERT INTO device_power (serial_number, device_make_model, model, device_type, total_power_watt, total_btu, total_power_cable, power_socket_type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.TotalPowerWatt, data.TotalBTU, data.TotalPowerCable, data.PowerSocketType)
	if err != nil {
		logger.ErrorLogger.Printf("Error creating DevicePowerDetail: %v", err)
		return err
	}
	logger.InfoLogger.Println("Created DevicePowerDetail successfully")
	return nil
}

// GetAllDevicePowerDetail retrieves all records from the DevicePowerDetail table.
func (db *DB) GetAllDevicePowerDetail() ([]models.DevicePowerDetail, error) {
	query := "SELECT * FROM device_power"
	rows, err := db.Query(query)
	if err != nil {
		logger.ErrorLogger.Printf("Error querying DevicePowerDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DevicePowerDetail
	for rows.Next() {
		var data models.DevicePowerDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.TotalPowerWatt, &data.TotalBTU, &data.TotalPowerCable, &data.PowerSocketType)
		if err != nil {
			logger.ErrorLogger.Printf("Error scanning DevicePowerDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	logger.InfoLogger.Println("Retrieved all DevicePowerDetail records successfully")
	return results, nil
}

// UpdateDevicePowerDetail updates an existing record in the device_power table based on the ID.
func (db *DB) UpdateDevicePowerDetail(id int, data *models.DevicePowerDetail) error {
	query := `
        UPDATE device_power
        SET serial_number = ?, device_make_model = ?, model = ?, device_type = ?, total_power_watt = ?, total_btu = ?, total_power_cable = ?, power_socket_type = ?
        WHERE id = ?
    `
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.TotalPowerWatt, data.TotalBTU, data.TotalPowerCable, data.PowerSocketType, id)
	if err != nil {
		logger.ErrorLogger.Printf("Error updating DevicePowerDetail: %v", err)
		return err
	}
	logger.InfoLogger.Printf("Updated DevicePowerDetail with ID %d successfully", id)
	return nil
}

// DeleteDevicePowerDetail deletes a record from the device_power table based on the ID.
func (db *DB) DeleteDevicePowerDetail(id int) error {
	query := "DELETE FROM device_power WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		logger.ErrorLogger.Printf("Error deleting DevicePowerDetail with ID %d: %v", id, err)
		return err
	}
	logger.InfoLogger.Printf("Deleted DevicePowerDetail with ID %d successfully", id)
	return nil
}

// FetchDataFromDevicePower retrieves data from table 4.
func (db *DB) FetchDataFromDevicePower() ([]*models.DevicePowerDetail, error) {
	query := "SELECT * FROM device_power"
	rows, err := db.Query(query)
	if err != nil {
		logger.ErrorLogger.Printf("Error fetching data from table 4: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []*models.DevicePowerDetail
	for rows.Next() {
		var data models.DevicePowerDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.TotalPowerWatt, &data.TotalBTU, &data.TotalPowerCable, &data.PowerSocketType)
		if err != nil {
			logger.ErrorLogger.Printf("Error scanning data from table 4: %v", err)
			return nil, err
		}
		results = append(results, &data)
	}

	logger.InfoLogger.Println("Fetched data from device_power table successfully")
	return results, nil
}
