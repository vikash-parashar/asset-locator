package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/vikash-parashar/asset-locator/db"
	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/models"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
)

// GetPowerDetails handles the GET request to retrieve power details.
func GetPowerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Handling GET request for Power Details")

		data, err := db.GetAllDevicePowerDetail()
		if err != nil {
			logger.ErrorLogger.Println("Failed to retrieve power details:", err)
			return
		}
		c.HTML(http.StatusOK, "power_details.html", data)
	}
}

func CreateNewPowerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Handling POST request for creating new Power Details")

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
			Id:              id,
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
			logger.ErrorLogger.Println("Failed to create new Power Details entry:", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Entry Added Successfully"},
		)
	}
}

// DeleteDevicePowerDetailHandler deletes a DevicePowerDetail record based on its ID.
func DeleteDevicePowerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Handling DELETE request for Power Details")

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.DeleteDevicePowerDetail(id); err != nil {
			logger.ErrorLogger.Println("Failed to delete Power Details:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Power Details"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Power Details deleted successfully"})
	}
}

// UpdateDevicePowerDetail updates a DevicePowerDetail record based on its ID.
func UpdateDevicePowerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Handling PUT request for updating Power Details")

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.ErrorLogger.Println("Invalid ID:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		logger.InfoLogger.Printf("Fetching original data for ID: %d", id)

		// Fetch the original data from the database
		originalData, err := db.GetDevicePowerDetailById(id)
		if err != nil {

			logger.ErrorLogger.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		type RequestData struct {
			SerialNumber    string `json:"serial_number"`
			DeviceMakeModel string `json:"device_make_model"`
			Model           string `json:"model"`
			DeviceType      string `json:"device_type"`
			TotalPowerWatt  string `json:"total_power_watt"`
			TotalBTU        string `json:"total_btu"`
			TotalPowerCable string `json:"total_power_cable"`
			PowerSocketType string `json:"power_socket_type"`
		}

		var requestData RequestData
		if err := c.ShouldBindJSON(&requestData); err != nil {

			logger.ErrorLogger.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		tpw, _ := strconv.Atoi(requestData.TotalPowerWatt)
		tpc, _ := strconv.Atoi(requestData.TotalPowerCable)
		tbtu, _ := strconv.ParseFloat(requestData.TotalBTU, 64)

		// Update fields only if data is available in the request
		if requestData.SerialNumber != "" {
			originalData.SerialNumber = requestData.SerialNumber
		}
		if requestData.DeviceMakeModel != "" {
			originalData.DeviceMakeModel = requestData.DeviceMakeModel
		}
		if requestData.Model != "" {
			originalData.Model = requestData.Model
		}
		if requestData.DeviceType != "" {
			originalData.DeviceType = requestData.DeviceType
		}
		if requestData.TotalPowerWatt != "" {
			originalData.TotalPowerWatt = tpw
		}
		if requestData.TotalBTU != "" {
			originalData.TotalBTU = tbtu
		}
		if requestData.TotalPowerCable != "" {
			originalData.TotalPowerCable = tpc
		}
		if requestData.PowerSocketType != "" {
			originalData.PowerSocketType = requestData.PowerSocketType
		}

		logger.InfoLogger.Printf("Updating data for ID: %d", id)

		// Update the data in the database
		if err := db.UpdateDevicePowerDetail(id, originalData); err != nil {
			logger.ErrorLogger.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		logger.InfoLogger.Println("Power Details updated successfully.")
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Power Details updated successfully"})
	}
}

// func UpdateDevicePowerDetail(db *db.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		logger.InfoLogger.Println("Updating Power Details")

// 		idStr := c.Param("id")
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			logger.ErrorLogger.Println("Invalid ID:", err)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 			return
// 		}

// 		// Fetch the original data from the database
// 		originalData, err := db.GetDevicePowerDetailById(id)
// 		if err != nil {
// 			logger.ErrorLogger.Println("Failed to fetch original data:", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch original data"})
// 			return
// 		}

// 		type RequestData struct {
// 			SerialNumber    string `json:"serial_number"`
// 			DeviceMakeModel string `json:"device_make_model"`
// 			Model           string `json:"model"`
// 			DeviceType      string `json:"device_type"`
// 			TotalPowerWatt  string `json:"total_power_watt"`
// 			TotalBTU        string `json:"total_btu"`
// 			TotalPowerCable string `json:"total_power_cable"`
// 			PowerSocketType string `json:"power_socket_type"`
// 		}

// 		var requestData RequestData
// 		if err := c.ShouldBindJSON(&requestData); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
// 			return
// 		}

// 		tpw, _ := strconv.Atoi(requestData.TotalPowerWatt)
// 		tpc, _ := strconv.Atoi(requestData.TotalPowerCable)
// 		tbtu, _ := strconv.ParseFloat(requestData.TotalBTU, 64)

// 		// Update fields only if data is available in the request
// 		if requestData.SerialNumber != "" {
// 			originalData.SerialNumber = requestData.SerialNumber
// 		}
// 		if requestData.DeviceMakeModel != "" {
// 			originalData.DeviceMakeModel = requestData.DeviceMakeModel
// 		}
// 		if requestData.Model != "" {
// 			originalData.Model = requestData.Model
// 		}
// 		if requestData.DeviceType != "" {
// 			originalData.DeviceType = requestData.DeviceType
// 		}
// 		if requestData.TotalPowerWatt != "" {
// 			originalData.TotalPowerWatt = tpw
// 		}
// 		if requestData.TotalBTU != "" {
// 			originalData.TotalBTU = tbtu
// 		}
// 		if requestData.TotalPowerCable != "" {
// 			originalData.TotalPowerCable = tpc
// 		}
// 		if requestData.PowerSocketType != "" {
// 			originalData.PowerSocketType = requestData.PowerSocketType
// 		}

// 		// Update the data in the database
// 		if err := db.UpdateDevicePowerDetail(id, originalData); err != nil {
// 			logger.ErrorLogger.Println("Failed to update Power Details:", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Power Details"})
// 			return
// 		}

// 		logger.InfoLogger.Println("Power Details updated successfully.")
// 		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Power Details updated successfully"})
// 	}
// }

// func UpdateDevicePowerDetail(db *db.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		logger.InfoLogger.Println("Handling PUT request for updating Power Details")
// 		idStr := c.Param("id")
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 			return
// 		}

// 		type RequestData struct {
// 			SerialNumber    string `json:"serial_number"`
// 			DeviceMakeModel string `json:"device_make_model"`
// 			Model           string `json:"model"`
// 			DeviceType      string `json:"device_type"`
// 			TotalPowerWatt  string `json:"total_power_watt"`
// 			TotalBTU        string `json:"total_btu"`
// 			TotalPowerCable string `json:"total_power_cable"`
// 			PowerSocketType string `json:"power_socket_type"`
// 		}

// 		var requestData RequestData
// 		if err := c.ShouldBindJSON(&requestData); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
// 			return
// 		}

// 		tpw, _ := strconv.Atoi(requestData.TotalPowerWatt)
// 		tpc, _ := strconv.Atoi(requestData.TotalPowerCable)
// 		tbtu, _ := strconv.ParseFloat(requestData.TotalBTU, 64)

// 		updatedData := &models.DevicePowerDetail{
// 			Id:              id,
// 			SerialNumber:    requestData.SerialNumber,
// 			DeviceMakeModel: requestData.DeviceMakeModel,
// 			Model:           requestData.Model,
// 			DeviceType:      requestData.DeviceType,
// 			TotalPowerWatt:  tpw,
// 			TotalBTU:        tbtu,
// 			TotalPowerCable: tpc,
// 			PowerSocketType: requestData.PowerSocketType,
// 		}

// 		if err := db.UpdateDevicePowerDetail(id, updatedData); err != nil {
// 			logger.ErrorLogger.Println("Failed to update Power Details:", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Power Details"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Power Details updated successfully"})
// 	}
// }

// DeleteDevicePowerDetailHandler deletes a DevicePowerDetail record based on its ID.
func DownloadDevicePowerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Handling GET request for downloading Power Details in Excel format")

		rows, err := db.Query("SELECT * FROM device_power")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to query the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a new Excel file
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("DevicePowerDetails")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to create Excel sheet", http.StatusInternalServerError)
			return
		}

		// Add header row
		headerRow := sheet.AddRow()
		headerRow.AddCell().SetString("ID")
		headerRow.AddCell().SetString("Serial Number")
		headerRow.AddCell().SetString("Device Make Model")
		headerRow.AddCell().SetString("Model")
		headerRow.AddCell().SetString("Device Type")
		headerRow.AddCell().SetString("Total Power Watt")
		headerRow.AddCell().SetString("Total BTU")
		headerRow.AddCell().SetString("Total Power Cable")
		headerRow.AddCell().SetString("Power Socket Type")

		// Add data rows from the database
		for rows.Next() {
			var device models.DevicePowerDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.TotalPowerWatt, &device.TotalBTU, &device.TotalPowerCable, &device.PowerSocketType); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}
			dataRow := sheet.AddRow()
			dataRow.AddCell().SetInt(device.Id)
			dataRow.AddCell().SetString(device.SerialNumber)
			dataRow.AddCell().SetString(device.DeviceMakeModel)
			dataRow.AddCell().SetString(device.Model)
			dataRow.AddCell().SetString(device.DeviceType)
			dataRow.AddCell().SetInt(device.TotalPowerWatt)
			dataRow.AddCell().SetFloat(device.TotalBTU)
			dataRow.AddCell().SetInt(device.TotalPowerCable)
			dataRow.AddCell().SetString(device.PowerSocketType)
		}

		// Save the Excel file to the response
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=DevicePowerDetails.xlsx")
		err = file.Write(c.Writer)
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to write Excel file to response", http.StatusInternalServerError)
		}
		logger.InfoLogger.Println("Power Details downloaded successfully in Excel format.")
	}
}

// DownloadDevicePowerDetailAsPDF exports DevicePowerDetail data as a PDF file.
func DownloadDevicePowerDetailPDF(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Handling GET request for downloading Power Details in PDF format")

		// Query the database for DevicePowerDetail data
		rows, err := db.Query("SELECT * FROM device_power")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to query the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a new PDF document
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

		// Set font and text size
		pdf.SetFont("Arial", "", 12)

		// Add table headers
		headers := []string{"ID", "Serial Number", "Device Make Model", "Model", "Device Type", "Total Power Watt", "Total BTU", "Total Power Cable", "Power Socket Type"}
		for _, header := range headers {
			pdf.CellFormat(40, 10, header, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)

		// Add data rows from the database
		for rows.Next() {
			var device models.DevicePowerDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.TotalPowerWatt, &device.TotalBTU, &device.TotalPowerCable, &device.PowerSocketType); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}

			data := []string{
				fmt.Sprint(device.Id),
				device.SerialNumber,
				device.DeviceMakeModel,
				device.Model,
				device.DeviceType,
				fmt.Sprint(device.TotalPowerWatt),
				fmt.Sprint(device.TotalBTU),
				fmt.Sprint(device.TotalPowerCable),
				device.PowerSocketType,
			}

			for _, str := range data {
				pdf.CellFormat(40, 10, str, "1", 0, "C", false, 0, "")
			}
			pdf.Ln(-1)
		}

		// Create the PDF file
		pdf.Output(c.Writer)

		// Set response headers
		c.Header("Content-Type", "application/pdf")
		c.Header("Content-Disposition", "attachment; filename=DevicePowerDetails.pdf")
		logger.InfoLogger.Println("Power Details downloaded successfully in PDF format.")
	}
}
