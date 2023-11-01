package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vikash-parashar/asset-locator/models"
)

// CreateDeviceEthernetFiberDetail creates a new record in the DeviceEthernetFiberDetail table.
func (db *DB) CreateDeviceEthernetFiberDetail(data *models.DeviceEthernetFiberDetail) error {
	query := `
		INSERT INTO device_ethernet_fiber (serial_number, device_make_model, model, device_type, device_physical_port, device_port_type, device_port_mac_address_wwn, connected_device_port)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DevicePhysicalPort, data.DevicePortType, data.DevicePortMACWWN, data.ConnectedDevicePort)
	if err != nil {
		log.Printf("Error creating DeviceEthernetFiberDetail: %v", err)
		return err
	}
	return nil
}

// GetAllDeviceEthernetFiberDetail retrieves all records from the DeviceEthernetFiberDetail table.
func (db *DB) GetAllDeviceEthernetFiberDetail() ([]models.DeviceEthernetFiberDetail, error) {
	query := "SELECT * FROM device_ethernet_fiber"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying DeviceEthernetFiberDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DeviceEthernetFiberDetail
	for rows.Next() {
		var data models.DeviceEthernetFiberDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DevicePhysicalPort, &data.DevicePortType, &data.DevicePortMACWWN, &data.ConnectedDevicePort)
		if err != nil {
			log.Printf("Error scanning DeviceEthernetFiberDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}

// Your GetFiberDetailByID function
func (db *DB) GetFiberDetailByID(id int) (models.DeviceEthernetFiberDetail, error) {
	var fiberDetail models.DeviceEthernetFiberDetail
	query := `
        SELECT * FROM device_ethernet_fiber WHERE id = $1
    `
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
			// Return a custom error when the fiber detail is not found
			return fiberDetail, fmt.Errorf("fiber detail with ID %d not found", id)
		}
		log.Printf("Error retrieving FiberDetail: %v", err)
		return fiberDetail, err
	}
	return fiberDetail, nil
}

func (db *DB) DeleteDeviceEthernetFiberDetail(id int) error {
	query := "DELETE FROM device_ethernet_fiber WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting DeviceEthernetFiberDetail: %v", err)
		return err
	}
	return nil
}

func (db *DB) UpdateDeviceEthernetFiberDetail(id int, data *models.DeviceEthernetFiberDetail) error {
	query := `
        UPDATE device_ethernet_fiber
        SET serial_number = $2, device_make_model = $3, model = $4, device_type = $5, device_physical_port = $6, device_port_type = $7, device_port_mac_address_wwn = $8, connected_device_port = $9
        WHERE id = $1
    `

	log.Println("data sent to db to update::")
	log.Println(data)
	_, err := db.Exec(query, &id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DevicePhysicalPort, &data.DevicePortType, &data.DevicePortMACWWN, &data.ConnectedDevicePort)
	if err != nil {
		log.Printf("Error updating DeviceEthernetFiberDetail: %v", err)
		return err
	}
	return nil
}
