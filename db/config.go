package db

import (
	"database/sql"
	"fmt"
	"log"

	"go-server/models"

	_ "github.com/lib/pq"
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
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// Close closes the database connection.
func (db *DB) Close() {
	db.DB.Close()
}

// CreateDevicePowerDetail creates a new record in the DevicePowerDetail table.
func (db *DB) CreateDevicePowerDetail(data *models.DevicePowerDetail) error {
	query := `
		INSERT INTO device_power (serial_number, device_make_model, model, device_type, total_power_watt, total_btu, total_power_cable, power_socket_type)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.TotalPowerWatt, data.TotalBTU, data.TotalPowerCable, data.PowerSocketType)
	if err != nil {
		log.Printf("Error creating DevicePowerDetail: %v", err)
		return err
	}
	return nil
}

// GetAllDevicePowerDetail retrieves all records from the DevicePowerDetail table.
func (db *DB) GetAllDevicePowerDetail() ([]models.DevicePowerDetail, error) {
	query := "SELECT * FROM device_power"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying DevicePowerDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DevicePowerDetail
	for rows.Next() {
		var data models.DevicePowerDetail
		err := rows.Scan(&data.ID, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.TotalPowerWatt, &data.TotalBTU, &data.TotalPowerCable, &data.PowerSocketType)
		if err != nil {
			log.Printf("Error scanning DevicePowerDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}

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
		err := rows.Scan(&data.ID, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DevicePhysicalPort, &data.DevicePortType, &data.DevicePortMACWWN, &data.ConnectedDevicePort)
		if err != nil {
			log.Printf("Error scanning DeviceEthernetFiberDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}

// CreateDeviceAMCOwnerDetail creates a new record in the DeviceAMCOwnerDetail table.
func (db *DB) CreateDeviceAMCOwnerDetail(data *models.DeviceAMCOwnerDetail) error {
	query := `
		INSERT INTO device_amc_owner (serial_number, device_make_model, model, po_number, po_order_date, eosl_date, amc_start_date, amc_end_date, device_owner)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.PONumber, data.POOrderDate, data.EOSLDate, data.AMCStartDate, data.AMCEndDate, data.DeviceOwner)
	if err != nil {
		log.Printf("Error creating DeviceAMCOwnerDetail: %v", err)
		return err
	}
	return nil
}

// GetAllDeviceAMCOwnerDetail retrieves all records from the DeviceAMCOwnerDetail table.
func (db *DB) GetAllDeviceAMCOwnerDetail() ([]models.DeviceAMCOwnerDetail, error) {
	query := "SELECT * FROM device_amc_owner"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying DeviceAMCOwnerDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DeviceAMCOwnerDetail
	for rows.Next() {
		var data models.DeviceAMCOwnerDetail
		err := rows.Scan(&data.ID, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.PONumber, &data.POOrderDate, &data.EOSLDate, &data.AMCStartDate, &data.AMCEndDate, &data.DeviceOwner)
		if err != nil {
			log.Printf("Error scanning DeviceAMCOwnerDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}

// CreateDeviceLocationDetail creates a new record in the DeviceLocationDetail table.
func (db *DB) CreateDeviceLocationDetail(data *models.DeviceLocationDetail) error {
	query := `
		INSERT INTO device_location (serial_number, device_make_model, model, device_type, data_center, region, dc_location, device_location, device_row_number, device_rack_number, device_ru_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := db.Exec(query, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DataCenter, data.Region, data.DCLocation, data.DeviceLocation, data.DeviceRowNumber, data.DeviceRackNumber, data.DeviceRUNumber)
	if err != nil {
		log.Printf("Error creating DeviceLocationDetail: %v", err)
		return err
	}
	return nil
}

// GetAllDeviceLocationDetail retrieves all records from the DeviceLocationDetail table.
func (db *DB) GetAllDeviceLocationDetail() ([]models.DeviceLocationDetail, error) {
	query := "SELECT * FROM device_location"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying DeviceLocationDetail: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.DeviceLocationDetail
	for rows.Next() {
		var data models.DeviceLocationDetail
		err := rows.Scan(&data.ID, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DataCenter, &data.Region, &data.DCLocation, &data.DeviceLocation, &data.DeviceRowNumber, &data.DeviceRackNumber, &data.DeviceRUNumber)
		if err != nil {
			log.Printf("Error scanning DeviceLocationDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}
