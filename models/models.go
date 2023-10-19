package models

type DevicePowerDetail struct {
	SerialNumber    string
	DeviceMakeModel string
	Model           string
	DeviceType      string
	TotalPowerWatt  int
	TotalBTU        float64
	TotalPowerCable int
	PowerSocketType string
}

type DeviceEthernetFiberDetail struct {
	SerialNumber        string
	DeviceMakeModel     string
	Model               string
	DeviceType          string
	DevicePhysicalPort  string
	DevicePortType      string
	DevicePortMACWWN    string
	ConnectedDevicePort string
}

type DeviceAMCOwnerDetail struct {
	SerialNumber    string
	DeviceMakeModel string
	Model           string
	PONumber        string
	POOrderDate     string
	EOSLDate        string
	AMCStartDate    string
	AMCEndDate      string
	DeviceOwner     string
}

type DeviceLocationDetail struct {
	SerialNumber     string
	DeviceMakeModel  string
	Model            string
	DeviceType       string
	DataCenter       string
	Region           string
	DCLocation       string
	DeviceLocation   string
	DeviceRowNumber  int
	DeviceRackNumber int
	DeviceRUNumber   string
}
