package handlers

import (
	"go-server/db"
	"go-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route for generating and serving Fiber data as Excel
func GenerateFiberDataExcel(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.GetAllDeviceEthernetFiberDetail()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
			return
		}
		GenerateAndServeExcelData(data, "fiber.xlsx", c)
	}
}

// Route for generating and serving Power data as Excel
func GeneratePowerDataExcel(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.GetAllDevicePowerDetail()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
			return
		}

		GenerateAndServeExcelData(data, "power.xlsx", c)
	}
}

// Route for generating and serving Owner data as Excel
func GenerateOwnerDataExcel(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.GetAllDeviceAMCOwnerDetail()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
			return
		}

		GenerateAndServeExcelData(data, "owner.xlsx", c)
	}
}

// Route for generating and serving Location data as Excel
func GenerateLocationDataExcel(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.GetAllDeviceLocationDetail()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
			return
		}

		GenerateAndServeExcelData(data, "location.xlsx", c)
	}
}

// Function to generate and serve Excel data
func GenerateAndServeExcelData(data interface{}, filename string, c *gin.Context) {
	// Generate the Excel file
	err := utils.GenerateExcelFile(data, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
		return
	}

	// Serve the Excel file for download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File(filename)
}

// func GenerateLocationDataPDF(db *db.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		data, err := db.GetAllDeviceLocationDetail()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		if err = utils.GenerateExcelFile(data, "location.xlsx"); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// func GenerateFiberDataPDF(db *db.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		data, err := db.GetAllDeviceEthernetFiberDetail()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// func GeneratePowerDataPDF(db *db.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		data, err := db.GetAllDevicePowerDetail()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// func GenerateOwnerDataPDF(db *db.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		data, err := db.GetAllDeviceAMCOwnerDetail()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }
