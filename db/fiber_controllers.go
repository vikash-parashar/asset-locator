package db

import (
	"database/sql"
	"fmt"

	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/models"
)

// CreateDeviceEthernetFiberDetail creates a new record in the DeviceEthernetFiberDetail table.
func (db *DB) CreateDeviceEthernetFiberDetail(data *models.DeviceEthernetFiberDetail) error {
	query := `
		INSERT INTO device_ethernet_fiber (serial_number, device_make_model, model, device_type, device_physical_port, device_port_type, device_port_mac_address_wwn, connected_device_port)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DevicePhysicalPort, data.DevicePortType, data.DevicePortMACWWN, data.ConnectedDevicePort)
	if err != nil {
		logger.ErrorLogger.Printf("Error creating DeviceEthernetFiberDetail: %v", err)
		return err
	}
	logger.InfoLogger.Println("Created DeviceEthernetFiberDetail successfully")
	return nil
}

// GetAllDeviceEthernetFiberDetail retrieves all records from the DeviceEthernetFiberDetail table.
func (db *DB) GetAllDeviceEthernetFiberDetail() ([]models.DeviceEthernetFiberDetail, error) {
	query := "SELECT * FROM device_ethernet_fiber"
	rows, err := db.Query(query)
	if err != nil {
		logger.ErrorLogger.Printf("Error querying DeviceEthernetFiberDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DeviceEthernetFiberDetail
	for rows.Next() {
		var data models.DeviceEthernetFiberDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DevicePhysicalPort, &data.DevicePortType, &data.DevicePortMACWWN, &data.ConnectedDevicePort)
		if err != nil {
			logger.ErrorLogger.Printf("Error scanning DeviceEthernetFiberDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	logger.InfoLogger.Println("Retrieved all DeviceEthernetFiberDetail records successfully")
	return results, nil
}

// Your GetFiberDetailByID function
func (db *DB) GetFiberDetailByID(id int) (models.DeviceEthernetFiberDetail, error) {
	var fiberDetail models.DeviceEthernetFiberDetail
	query := "SELECT * FROM device_ethernet_fiber WHERE id = ?"
	err := db.QueryRow(query, id).Scan(
		&fiberDetail.Id,
		&fiberDetail.SerialNumber,
		&fiberDetail.DeviceMakeModel,
		&fiberDetail.Model,
		&fiberDetail.DeviceType,
		&fiberDetail.DevicePhysicalPort,
		&fiberDetail.DevicePortType,
		&fiberDetail.DevicePortMACWWN,
		&fiberDetail.ConnectedDevicePort,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.ErrorLogger.Printf("Fiber detail with ID %d not found", id)
			return fiberDetail, fmt.Errorf("fiber detail with ID %d not found", id)
		}
		logger.ErrorLogger.Printf("Error retrieving FiberDetail: %v", err)
		return fiberDetail, err
	}
	logger.InfoLogger.Printf("Retrieved FiberDetail with ID %d successfully", id)
	return fiberDetail, nil
}

func (db *DB) DeleteDeviceEthernetFiberDetail(id int) error {
	query := "DELETE FROM device_ethernet_fiber WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		logger.ErrorLogger.Printf("Error deleting DeviceEthernetFiberDetail with ID %d: %v", id, err)
		return err
	}
	logger.InfoLogger.Printf("Deleted DeviceEthernetFiberDetail with ID %d successfully", id)
	return nil
}

func (db *DB) UpdateDeviceEthernetFiberDetail(id int, data *models.DeviceEthernetFiberDetail) error {
	query := `
        UPDATE device_ethernet_fiber
        SET serial_number = ?, device_make_model = ?, model = ?, device_type = ?, device_physical_port = ?, device_port_type = ?, device_port_mac_address_wwn = ?, connected_device_port = ?
        WHERE id = ?
    `

	logger.InfoLogger.Printf("Updating DeviceEthernetFiberDetail with ID %d", id)
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DevicePhysicalPort, data.DevicePortType, data.DevicePortMACWWN, data.ConnectedDevicePort, id)
	if err != nil {
		logger.ErrorLogger.Printf("Error updating DeviceEthernetFiberDetail: %v", err)
		return err
	}
	logger.InfoLogger.Printf("Updated DeviceEthernetFiberDetail with ID %d successfully", id)
	return nil
}
