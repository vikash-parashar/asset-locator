package handlers

import (
	"go-server/db"
	"go-server/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
}

// GetLocationDetails handles the GET request to retrieve location details.
func GetLocationDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.GetAllDeviceLocationDetail()
		if err != nil {
			log.Println(err)
			return
		}
		c.HTML(http.StatusOK, "location_details.html", data)
	}
}

// GetOwnerDetails handles the GET request to retrieve owner details.
func GetOwnerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.GetAllDeviceAMCOwnerDetail()
		if err != nil {
			log.Println(err)
			return
		}
		c.HTML(http.StatusOK, "owner_details.html", data)
	}
}

// GetPowerDetails handles the GET request to retrieve power details.
func GetPowerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.GetAllDevicePowerDetail()
		if err != nil {
			log.Println(err)
			return
		}
		c.HTML(http.StatusOK, "power_details.html", data)
	}
}

// GetFiberDetails handles the GET request to retrieve fiber details.
func GetFiberDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.GetAllDeviceEthernetFiberDetail()
		if err != nil {
			log.Println(err)
			return
		}
		c.HTML(http.StatusOK, "fiber_details.html", data)
	}
}

func CreateNewLocationDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.DeviceLocationDetail

		// Retrieve form data
		idStr := c.PostForm("id")
		serialNumber := c.PostForm("serial_number")
		deviceMakeModel := c.PostForm("device_make_model")
		model := c.PostForm("model")
		deviceType := c.PostForm("device_type")
		dataCenter := c.PostForm("data_center")
		region := c.PostForm("region")
		dcLocation := c.PostForm("dc_location")
		deviceLocation := c.PostForm("device_location")
		deviceRowNumberStr := c.PostForm("device_row_number")
		deviceRackNumberStr := c.PostForm("device_rack_number")
		deviceRUNumber := c.PostForm("device_ru_number")

		// Parse and cast the string values to their respective types
		id, _ := strconv.Atoi(idStr)
		deviceRowNumber, _ := strconv.Atoi(deviceRowNumberStr)
		deviceRackNumber, _ := strconv.Atoi(deviceRackNumberStr)

		// Assign the values to the DeviceLocationDetail struct
		data = models.DeviceLocationDetail{
			ID:               id,
			SerialNumber:     serialNumber,
			DeviceMakeModel:  deviceMakeModel,
			Model:            model,
			DeviceType:       deviceType,
			DataCenter:       dataCenter,
			Region:           region,
			DCLocation:       dcLocation,
			DeviceLocation:   deviceLocation,
			DeviceRowNumber:  deviceRowNumber,
			DeviceRackNumber: deviceRackNumber,
			DeviceRUNumber:   deviceRUNumber,
		}

		if err := db.CreateDeviceLocationDetail(&data); err != nil {
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Entry Added Successfully"},
		)
	}
}

func CreateNewOwnerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.DeviceAMCOwnerDetail

		// Retrieve form data
		idStr := c.PostForm("id")
		serialNumber := c.PostForm("serial_number")
		deviceMakeModel := c.PostForm("device_make_model")
		model := c.PostForm("model")
		poNumber := c.PostForm("po_number")
		poOrderDateStr := c.PostForm("po_order_date")
		eoslDateStr := c.PostForm("eosl_date")
		amcStartDateStr := c.PostForm("amc_start_date")
		amcEndDateStr := c.PostForm("amc_end_date")
		deviceOwner := c.PostForm("device_owner")

		// Parse the date strings into time.Time values
		poOrderDate, _ := time.Parse("2006-01-02", poOrderDateStr)
		eoslDate, _ := time.Parse("2006-01-02", eoslDateStr)
		amcStartDate, _ := time.Parse("2006-01-02", amcStartDateStr)
		amcEndDate, _ := time.Parse("2006-01-02", amcEndDateStr)

		// Convert ID to int
		id, _ := strconv.Atoi(idStr)

		// Assign the values to the DeviceAMCOwnerDetail struct
		data = models.DeviceAMCOwnerDetail{
			ID:              id,
			SerialNumber:    serialNumber,
			DeviceMakeModel: deviceMakeModel,
			Model:           model,
			PONumber:        poNumber,
			POOrderDate:     poOrderDate,
			EOSLDate:        eoslDate,
			AMCStartDate:    amcStartDate,
			AMCEndDate:      amcEndDate,
			DeviceOwner:     deviceOwner,
		}

		if err := db.CreateDeviceAMCOwnerDetail(&data); err != nil {
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Entry Added Successfully"},
		)
	}
}

func CreateNewPowerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.DevicePowerDetail

		// Retrieve form data
		idStr := c.PostForm("id")
		serialNumber := c.PostForm("serial_number")
		deviceMakeModel := c.PostForm("device_make_model")
		model := c.PostForm("model")
		deviceType := c.PostForm("device_type")
		totalPowerWattStr := c.PostForm("total_power_watt")
		totalBTUStr := c.PostForm("total_btu")
		totalPowerCableStr := c.PostForm("total_power_cable")
		powerSocketType := c.PostForm("power_socket_type")

		// Parse and cast the string values to their respective types
		id, _ := strconv.Atoi(idStr)
		totalPowerWatt, _ := strconv.Atoi(totalPowerWattStr)
		totalBTU, _ := strconv.ParseFloat(totalBTUStr, 64)
		totalPowerCable, _ := strconv.Atoi(totalPowerCableStr)

		// Assign the values to the DevicePowerDetail struct
		data = models.DevicePowerDetail{
			ID:              id,
			SerialNumber:    serialNumber,
			DeviceMakeModel: deviceMakeModel,
			Model:           model,
			DeviceType:      deviceType,
			TotalPowerWatt:  totalPowerWatt,
			TotalBTU:        totalBTU,
			TotalPowerCable: totalPowerCable,
			PowerSocketType: powerSocketType,
		}

		if err := db.CreateDevicePowerDetail(&data); err != nil {
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Entry Added Successfully"},
		)
	}
}

func CreateNewFiberDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.DeviceEthernetFiberDetail

		// Retrieve form data in a similar manner as your original code
		idStr := c.PostForm("id")
		serialNumber := c.PostForm("serial_number")
		deviceMakeModel := c.PostForm("device_make_model")
		model := c.PostForm("model")
		deviceType := c.PostForm("device_type")
		devicePhysicalPort := c.PostForm("device_physical_port")
		devicePortType := c.PostForm("device_port_type")
		devicePortMACWWN := c.PostForm("device_port_macwwn")
		connectedDevicePort := c.PostForm("connected_device_port")

		// Parse and cast the string values to their respective types
		id, _ := strconv.Atoi(idStr)

		// Assign the values to the DeviceEthernetFiberDetail struct
		data = models.DeviceEthernetFiberDetail{
			ID:                  id,
			SerialNumber:        serialNumber,
			DeviceMakeModel:     deviceMakeModel,
			Model:               model,
			DeviceType:          deviceType,
			DevicePhysicalPort:  devicePhysicalPort,
			DevicePortType:      devicePortType,
			DevicePortMACWWN:    devicePortMACWWN,
			ConnectedDevicePort: connectedDevicePort,
		}

		if err := db.CreateDeviceEthernetFiberDetail(&data); err != nil {
			log.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Entry Added Successfully",
		})
	}
}
