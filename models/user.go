package models

const (
	UserRoleAdmin   = "admin"
	UserRoleGeneral = "general"
)
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

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}
