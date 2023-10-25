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
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.TotalPowerWatt, &data.TotalBTU, &data.TotalPowerCable, &data.PowerSocketType)
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
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DevicePhysicalPort, &data.DevicePortType, &data.DevicePortMACWWN, &data.ConnectedDevicePort)
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
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.PONumber, &data.POOrderDate, &data.EOSLDate, &data.AMCStartDate, &data.AMCEndDate, &data.DeviceOwner)
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
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DataCenter, &data.Region, &data.DCLocation, &data.DeviceLocation, &data.DeviceRowNumber, &data.DeviceRackNumber, &data.DeviceRUNumber)
		if err != nil {
			log.Printf("Error scanning DeviceLocationDetail: %v", err)
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
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
	_, err := db.Exec(query, id, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DevicePhysicalPort, data.DevicePortType, data.DevicePortMACWWN, data.ConnectedDevicePort)
	if err != nil {
		log.Printf("Error updating DeviceEthernetFiberDetail: %v", err)
		return err
	}
	return nil
}

// UpdateDeviceLocationDetail updates an existing record in the device_location table based on the ID.
func (db *DB) UpdateDeviceLocationDetail(id int, data *models.DeviceLocationDetail) error {
	query := `
        UPDATE device_location
        SET serial_number = $2, device_make_model = $3, model = $4, device_type = $5, data_center = $6, region = $7, dc_location = $8, device_location = $9, device_row_number = $10, device_rack_number = $11, device_ru_number = $12
        WHERE id = $1
    `
	_, err := db.Exec(query, id, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.DataCenter, data.Region, data.DCLocation, data.DeviceLocation, data.DeviceRowNumber, data.DeviceRackNumber, data.DeviceRUNumber)
	if err != nil {
		log.Printf("Error updating DeviceLocationDetail: %v", err)
		return err
	}
	return nil
}

// DeleteDeviceLocationDetail deletes a record from the device_location table based on the ID.
func (db *DB) DeleteDeviceLocationDetail(id int) error {
	query := "DELETE FROM device_location WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting DeviceLocationDetail: %v", err)
		return err
	}
	return nil
}

// UpdateDeviceAMCOwnerDetail updates an existing record in the device_amc_owner table based on the ID.
func (db *DB) UpdateDeviceAMCOwnerDetail(id int, data *models.DeviceAMCOwnerDetail) error {
	query := `
        UPDATE device_amc_owner
        SET serial_number = $2, device_make_model = $3, model = $4, po_number = $5, po_order_date = $6, eosl_date = $7, amc_start_date = $8, amc_end_date = $9, device_owner = $10
        WHERE id = $1
    `
	_, err := db.Exec(query, id, data.SerialNumber, data.DeviceMakeModel, data.Model, data.PONumber, data.POOrderDate, data.EOSLDate, data.AMCStartDate, data.AMCEndDate, data.DeviceOwner)
	if err != nil {
		log.Printf("Error updating DeviceAMCOwnerDetail: %v", err)
		return err
	}
	return nil
}

// DeleteDeviceAMCOwnerDetail deletes a record from the device_amc_owner table based on the ID.
func (db *DB) DeleteDeviceAMCOwnerDetail(id int) error {
	query := "DELETE FROM device_amc_owner WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting DeviceAMCOwnerDetail: %v", err)
		return err
	}
	return nil
}

// UpdateDevicePowerDetail updates an existing record in the device_power table based on the ID.
func (db *DB) UpdateDevicePowerDetail(id int, data *models.DevicePowerDetail) error {
	query := `
        UPDATE device_power
        SET serial_number = $2, device_make_model = $3, model = $4, device_type = $5, total_power_watt = $6, total_btu = $7, total_power_cable = $8, power_socket_type = $9
        WHERE id = $1
    `
	_, err := db.Exec(query, id, data.SerialNumber, data.DeviceMakeModel, data.Model, data.DeviceType, data.TotalPowerWatt, data.TotalBTU, data.TotalPowerCable, data.PowerSocketType)
	if err != nil {
		log.Printf("Error updating DevicePowerDetail: %v", err)
		return err
	}
	return nil
}

// DeleteDevicePowerDetail deletes a record from the device_power table based on the ID.
func (db *DB) DeleteDevicePowerDetail(id int) error {
	query := "DELETE FROM device_power WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting DevicePowerDetail: %v", err)
		return err
	}
	return nil
}

// GetUserByEmailID retrieves a user by their email address.
func (db *DB) GetUserByEmailID(email string) (*models.User, error) {
	query := `
		SELECT * FROM users
		WHERE email = $1
	`
	user := &models.User{}
	err := db.QueryRow(query, email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Phone, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Printf("Error fetching user by email: %v", err)
		return nil, err
	}
	return user, nil
}

// GetUserByUsername retrieves a user by their username.
func (db *DB) GetUserByUsername(username string) (*models.User, error) {
	query := `
		SELECT * FROM users
		WHERE username = $1
	`
	user := &models.User{}
	err := db.QueryRow(query, username).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Phone, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Printf("Error fetching user by username: %v", err)
		return nil, err
	}
	return user, nil
}

// RegisterUser creates a new user record in the database.
func (db *DB) RegisterUser(user *models.User) error {
	query := `
		INSERT INTO users (first_name, last_name, username, email, password, role)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err := db.QueryRow(query, user.FirstName, user.LastName, user.Phone, user.Email, user.Password, user.Role).Scan(&user.Id)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return err
	}
	return nil
}

// UpdateUserPassword updates a user's password by their ID.
func (db *DB) UpdateUserPassword(userID int, newPassword string) error {
	query := `
		UPDATE users
		SET password = $2, updated_at = NOW()
		WHERE id = $1
	`
	_, err := db.Exec(query, userID, newPassword)
	if err != nil {
		log.Printf("Error updating user password: %v", err)
		return err
	}
	return nil
}

// GetAllUsers retrieves all active user records.
func (db *DB) GetAllUsers() ([]*models.User, error) {
	query := `
		SELECT * FROM users
		WHERE deleted_at IS NULL
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching all users: %v", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Phone, &user.Email, &user.Password, &user.Role)
		if err != nil {
			log.Printf("Error scanning user rows: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over user rows: %v", err)
		return nil, err
	}
	return users, nil
}

// FetchDataFromTable1 retrieves data from table 1.
func (db *DB) FetchDataFromDeviceOwner() ([]*models.DeviceAMCOwnerDetail, error) {
	query := "SELECT * FROM device_amc_owner"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching data from table 1: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []*models.DeviceAMCOwnerDetail
	for rows.Next() {
		var data models.DeviceAMCOwnerDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.PONumber, &data.POOrderDate, &data.EOSLDate, &data.AMCStartDate, &data.AMCEndDate, &data.DeviceOwner)
		if err != nil {
			log.Printf("Error scanning data from table 1: %v", err)
			return nil, err
		}
		results = append(results, &data)
	}

	return results, nil
}

// FetchDataFromTable2 retrieves data from table 2.
func (db *DB) FetchDataFromDeviceFiber() ([]*models.DeviceEthernetFiberDetail, error) {
	query := "SELECT * FROM device_ethernet_fiber"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching data from table 2: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []*models.DeviceEthernetFiberDetail
	for rows.Next() {
		var data models.DeviceEthernetFiberDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DevicePhysicalPort, &data.DevicePortType, &data.DevicePortMACWWN, &data.ConnectedDevicePort)
		if err != nil {
			log.Printf("Error scanning data from table 2: %v", err)
			return nil, err
		}
		results = append(results, &data)
	}

	return results, nil
}

// FetchDataFromTable3 retrieves data from table 3.
func (db *DB) FetchDataFromDeviceLocation() ([]*models.DeviceLocationDetail, error) {
	query := "SELECT * FROM device_location"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching data from table 3: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []*models.DeviceLocationDetail
	for rows.Next() {
		var data models.DeviceLocationDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.DataCenter, &data.Region, &data.DCLocation, &data.DeviceLocation, &data.DeviceRowNumber, &data.DeviceRackNumber, &data.DeviceRUNumber)
		if err != nil {
			log.Printf("Error scanning data from table 3: %v", err)
			return nil, err
		}
		results = append(results, &data)
	}

	return results, nil
}

// FetchDataFromTable4 retrieves data from table 4.
func (db *DB) FetchDataFromDevicePower() ([]*models.DevicePowerDetail, error) {
	query := "SELECT * FROM device_power"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching data from table 4: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []*models.DevicePowerDetail
	for rows.Next() {
		var data models.DevicePowerDetail
		err := rows.Scan(&data.Id, &data.SerialNumber, &data.DeviceMakeModel, &data.Model, &data.DeviceType, &data.TotalPowerWatt, &data.TotalBTU, &data.TotalPowerCable, &data.PowerSocketType)
		if err != nil {
			log.Printf("Error scanning data from table 4: %v", err)
			return nil, err
		}
		results = append(results, &data)
	}

	return results, nil
}
