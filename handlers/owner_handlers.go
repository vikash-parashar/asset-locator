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

// UpdateDeviceAMCOwnerDetailHandler updates a DeviceAMCOwnerDetail record based on its ID.
func UpdateDeviceAMCOwnerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var data models.DeviceAMCOwnerDetail
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		if err := db.UpdateDeviceAMCOwnerDetail(id, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update DeviceAMCOwnerDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceAMCOwnerDetail updated successfully"})
	}
}

// DeleteDeviceAMCOwnerDetailHandler deletes a DeviceAMCOwnerDetail record based on its ID.
func DeleteDeviceAMCOwnerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.DeleteDeviceAMCOwnerDetail(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete DeviceAMCOwnerDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceAMCOwnerDetail deleted successfully"})
	}
}
