package handlers

import (
	"go-server/db"
	"go-server/models"
	"log"
	"net/http"

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
		log.Println(data)
		// c.JSON(http.StatusFound, gin.H{"status": "data fetched successfully", "data": data})
		c.HTML(http.StatusOK, "location_details.html", data)
		// render.RenderTemplate(c.Writer, "location_details", gin.H{"data": data})
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
		log.Println(data)
		// c.JSON(http.StatusFound, gin.H{"status": "data fetched successfully", "data": data})
		c.HTML(http.StatusOK, "owner_details.html", data)
		// render.RenderTemplate(c.Writer, "owner_details", gin.H{"data": data})
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
		log.Println(data)
		// c.JSON(http.StatusFound, gin.H{"status": "data fetched successfully", "data": data})
		c.HTML(http.StatusOK, "power_details.html", data)
		// render.RenderTemplate(c.Writer, "power_details", gin.H{"data": data})
	}
}

// GetFiberDetails handles the GET request to retrieve fiber details.
func GetFiberDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement your logic to fetch fiber details from the database and return them as a JSON response
		// Example: data, err := db.GetAllDeviceEthernetFiberDetail()
		// Handle errors and send a response
		data, err := db.GetAllDeviceEthernetFiberDetail()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(data)
		// c.JSON(http.StatusFound, gin.H{"status": "data fetched successfully", "data": data})
		c.HTML(http.StatusOK, "fiber_details.html", data)
		// render.RenderTemplate(c.Writer, "fiber_details", gin.H{"data": data})
	}
}

// CreateNewLocationDetails handles the POST request to create new location details.
func CreateNewLocationDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse JSON request payload and create a new location detail record in the database
		var data models.DeviceLocationDetail
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Implement your logic to create the new record
		// Example: err := db.CreateDeviceLocationDetail(&data)
		// Handle errors and send a response
	}
}

// CreateNewOwnerDetails handles the POST request to create new owner details.
func CreateNewOwnerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse JSON request payload and create a new owner detail record in the database
		var data models.DeviceAMCOwnerDetail
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Implement your logic to create the new record
		// Example: err := db.CreateDeviceAMCOwnerDetail(&data)
		// Handle errors and send a response
	}
}

// CreateNewPowerDetails handles the POST request to create new power details.
func CreateNewPowerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse JSON request payload and create a new power detail record in the database
		var data models.DevicePowerDetail
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Implement your logic to create the new record
		// Example: err := db.CreateDevicePowerDetail(&data)
		// Handle errors and send a response
	}
}

// CreateNewFiberDetails handles the POST request to create new fiber details.
func CreateNewFiberDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse JSON request payload and create a new fiber detail record in the database
		var data models.DeviceEthernetFiberDetail
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Implement your logic to create the new record
		// Example: err := db.CreateDeviceEthernetFiberDetail(&data)
		// Handle errors and send a response
	}
}
