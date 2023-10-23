package handlers

import (
	"go-server/db"
	"go-server/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

// UpdateDevicePowerDetailHandler updates a DevicePowerDetail record based on its ID.
func UpdateDevicePowerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var data models.DevicePowerDetail
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		if err := db.UpdateDevicePowerDetail(id, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update DevicePowerDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DevicePowerDetail updated successfully"})
	}
}

// DeleteDevicePowerDetailHandler deletes a DevicePowerDetail record based on its ID.
func DeleteDevicePowerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.DeleteDevicePowerDetail(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete DevicePowerDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DevicePowerDetail deleted successfully"})
	}
}
