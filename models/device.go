package models

import "time"

const (
	DeviceTypeServer        = "Server"
	DeviceTypeObjectStorage = "Object Storage"
	DeviceTypeSwitch        = "Switch"
	DeviceTypeSANSwitch     = "SAN Switch"
)
const (
	DeviceMakeOracle  = "Oracle"
	DeviceMakeHitachi = "Hitachi"
	DeviceMakeCisco   = "Cisco"
	DeviceMakeBrocade = "Brocade"
)

const (
	DeviceModelT84        = "T8-4"
	DeviceModelHCPG10     = "HCP-G10"
	DeviceModelNEXUS93108 = "NEXUS-93108"
	DeviceModel6520       = "6520"
)

type DevicePowerDetail struct {
	ID              int     `json:"id"`
	SerialNumber    string  `json:"serial_number"`
	DeviceMakeModel string  `json:"device_make_model"`
	Model           string  `json:"model"`
	DeviceType      string  `json:"device_type"`
	TotalPowerWatt  int     `json:"total_power_watt"`
	TotalBTU        float64 `json:"total_btu"`
	TotalPowerCable int     `json:"total_power_cable"`
	PowerSocketType string  `json:"power_socket_type"`
	Created_At      string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	DeletedAt       string  `json:"deleted_at"`
}

type DeviceEthernetFiberDetail struct {
	ID                  int    `json:"id"`
	SerialNumber        string `json:"serial_number"`
	DeviceMakeModel     string `json:"device_make_model"`
	Model               string `json:"model"`
	DeviceType          string `json:"device_type"`
	DevicePhysicalPort  string `json:"device_physical_port"`
	DevicePortType      string `json:"device_port_type"`
	DevicePortMACWWN    string `json:"device_port_macwwn"`
	ConnectedDevicePort string `json:"connected_device_port"`
	Created_At          string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	DeletedAt           string `json:"deleted_at"`
}

type DeviceAMCOwnerDetail struct {
	ID              int       `json:"id"`
	SerialNumber    string    `json:"serial_number"`
	DeviceMakeModel string    `json:"device_make_model"`
	Model           string    `json:"model"`
	PONumber        string    `json:"po_number"`
	POOrderDate     time.Time `json:"po_order_date"`
	EOSLDate        time.Time `json:"eosl_date"`
	AMCStartDate    time.Time `json:"amc_start_date"`
	AMCEndDate      time.Time `json:"amc_end_date"`
	DeviceOwner     string    `json:"device_owner"`
	Created_At      string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
	DeletedAt       string    `json:"deleted_at"`
}

type DeviceLocationDetail struct {
	ID               int    `json:"id"`
	SerialNumber     string `json:"serial_number"`
	DeviceMakeModel  string `json:"device_make_model"`
	Model            string `json:"model"`
	DeviceType       string `json:"device_type"`
	DataCenter       string `json:"data_center"`
	Region           string `json:"region"`
	DCLocation       string `json:"dc_location"`
	DeviceLocation   string `json:"device_location"`
	DeviceRowNumber  int    `json:"device_row_number"`
	DeviceRackNumber int    `json:"device_rack_number"`
	DeviceRUNumber   string `json:"device_ru_number"`
	Created_At       string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
}
