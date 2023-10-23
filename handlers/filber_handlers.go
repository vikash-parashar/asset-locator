package handlers

import (
	"go-server/db"
	"go-server/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

// UpdateDeviceEthernetFiberDetailHandler updates a DeviceEthernetFiberDetail record based on its ID.
func UpdateDeviceEthernetFiberDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var data models.DeviceEthernetFiberDetail
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		if err := db.UpdateDeviceEthernetFiberDetail(id, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update DeviceEthernetFiberDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceEthernetFiberDetail updated successfully"})
	}
}

// DeleteDeviceEthernetFiberDetailHandler deletes a DeviceEthernetFiberDetail record based on its ID.
func DeleteDeviceEthernetFiberDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.DeleteDeviceEthernetFiberDetail(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete DeviceEthernetFiberDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceEthernetFiberDetail deleted successfully"})
	}
}
