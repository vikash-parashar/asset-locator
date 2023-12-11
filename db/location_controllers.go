package db

import (
	"github.com/vikash-parashar/asset-locator/logger" // Import the logger package
	"github.com/vikash-parashar/asset-locator/models"
)

// CreateDeviceLocationDetail creates a new record in the DeviceLocationDetail table.
func (db *DB) CreateDeviceLocationDetail(data *models.DeviceLocationDetail) error {
	query := `
		INSERT INTO device_location (serial_number, device_make_model, model, device_type, data_center, region, dc_location, device_location, device_row_number, device_rack_number, device_ru_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DataCenter, data.Region, data.DCLocation, data.DeviceLocation, data.DeviceRowNumber, data.DeviceRackNumber, data.DeviceRUNumber)
	if err != nil {
		logger.ErrorLogger.Printf("Error creating DeviceLocationDetail: %v", err)
		return err
	}
	logger.InfoLogger.Println("Created DeviceLocationDetail successfully")
	return nil
}

// GetAllDeviceLocationDetail retrieves all records from the DeviceLocationDetail table.
func (db *DB) GetAllDeviceLocationDetail() ([]models.DeviceLocationDetail, error) {
	query := "SELECT * FROM device_location"
	rows, err := db.Query(query)
	if err != nil {
		logger.ErrorLogger.Printf("Error querying DeviceLocationDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DeviceLocationDetail
	for rows.Next() {
		var data models.DeviceLocationDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DataCenter, &data.Region, &data.DCLocation, &data.DeviceLocation, &data.DeviceRowNumber, &data.DeviceRackNumber, &data.DeviceRUNumber)
		if err != nil {
			logger.ErrorLogger.Printf("Error scanning DeviceLocationDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	logger.InfoLogger.Println("Retrieved all DeviceLocationDetail records successfully")
	return results, nil
}

// // UpdateDeviceLocationDetail updates an existing record in the device_location table based on the ID.
// func (db *DB) UpdateDeviceLocationDetail(id int, data *models.DeviceLocationDetail) error {
// 	query := `
//         UPDATE device_location
//         SET serial_number = $2, device_make_model = $3, model = $4, device_type = $5, data_center = $6, region = $7, dc_location = $8, device_location = $9, device_row_number = $10, device_rack_number = $11, device_ru_number = $12
//         WHERE id = $1
//     `
// 	_, err := db.Exec(query, id, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DataCenter, data.Region, data.DCLocation, data.DeviceLocation, data.DeviceRowNumber, data.DeviceRackNumber, data.DeviceRUNumber)
// 	if err != nil {
// 		logger.ErrorLogger.Printf("Error updating DeviceLocationDetail: %v", err)
// 		return err
// 	}
// 	logger.InfoLogger.Printf("Updated DeviceLocationDetail with ID %d successfully", id)
// 	return nil
// }

func (db *DB) UpdateDeviceLocationDetail(id int, data *models.DeviceLocationDetail) error {
	query := `
        UPDATE device_location
        SET serial_number = COALESCE($2, serial_number),
            device_make_model = COALESCE($3, device_make_model),
            model = COALESCE($4, model),
            device_type = COALESCE($5, device_type),
            data_center = COALESCE($6, data_center),
            region = COALESCE($7, region), 
            dc_location = COALESCE($8, dc_location),
            device_location = COALESCE($9, device_location),
            device_row_number = COALESCE($10, device_row_number),
            device_rack_number = COALESCE($11, device_rack_number),
            device_ru_number = COALESCE($12, device_ru_number)
        WHERE id = $1
    `
	_, err := db.Exec(query, id, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DataCenter, data.Region, data.DCLocation, data.DeviceLocation, data.DeviceRowNumber, data.DeviceRackNumber, data.DeviceRUNumber)
	if err != nil {
		logger.ErrorLogger.Printf("Error updating DeviceLocationDetail: %v", err)
		return err
	}
	logger.InfoLogger.Printf("Updated DeviceLocationDetail with ID %d successfully", id)
	return nil
}

// DeleteDeviceLocationDetail deletes a record from the device_location table based on the ID.
func (db *DB) DeleteDeviceLocationDetail(id int) error {
	query := "DELETE FROM device_location WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		logger.ErrorLogger.Printf("Error deleting DeviceLocationDetail with ID %d: %v", id, err)
		return err
	}
	logger.InfoLogger.Printf("Deleted DeviceLocationDetail with ID %d successfully", id)
	return nil
}

// GetDeviceLocationDetail retrieves a record from the device_location table based on the ID.
func (db *DB) GetDeviceLocationDetailById(id int) (models.DeviceLocationDetail, error) {
	query := "SELECT * FROM device_location WHERE id = $1"
	var deviceLocationDetail models.DeviceLocationDetail

	err := db.QueryRow(query, id).Scan(
		&deviceLocationDetail.Id,
		&deviceLocationDetail.SerialNumber,
		&deviceLocationDetail.DeviceMakeModel,
		&deviceLocationDetail.Model,
		&deviceLocationDetail.DeviceType,
		&deviceLocationDetail.DataCenter,
		&deviceLocationDetail.Region,
		&deviceLocationDetail.DCLocation,
		&deviceLocationDetail.DeviceLocation,
		&deviceLocationDetail.DeviceRowNumber,
		&deviceLocationDetail.DeviceRackNumber,
		&deviceLocationDetail.DeviceRUNumber,
	)

	if err != nil {
		logger.ErrorLogger.Printf("Error fetching DeviceLocationDetail with ID %d: %v", id, err)
		return models.DeviceLocationDetail{}, err
	}

	return deviceLocationDetail, nil
}

// FetchDataFromTable3 retrieves data from table 3.
func (db *DB) FetchDataFromDeviceLocation() ([]*models.DeviceLocationDetail, error) {
	query := "SELECT * FROM device_location"
	rows, err := db.Query(query)
	if err != nil {
		logger.ErrorLogger.Printf("Error fetching data from table 3: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []*models.DeviceLocationDetail
	for rows.Next() {
		var data models.DeviceLocationDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DataCenter, &data.Region, &data.DCLocation, &data.DeviceLocation, &data.DeviceRowNumber, &data.DeviceRackNumber, &data.DeviceRUNumber)
		if err != nil {
			logger.ErrorLogger.Printf("Error scanning data from table 3: %v", err)
			return nil, err
		}
		results = append(results, &data)
	}

	logger.InfoLogger.Println("Fetched data from device_location table successfully")
	return results, nil
}
