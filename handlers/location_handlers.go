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

func GetLocationDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Fetching location details from the database")
		data, err := db.GetAllDeviceLocationDetail()
		if err != nil {
			logger.ErrorLogger.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
			return
		}
		logger.InfoLogger.Println("Location details fetched successfully.")
		c.HTML(http.StatusOK, "location_details.html", data)
	}
}

func CreateNewLocationDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.DeviceLocationDetail

		logger.InfoLogger.Println("Creating new location details")

		// Retrieve form data
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

		deviceRowNumber, err := strconv.Atoi(deviceRowNumberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device_row_number"})
			return
		}
		deviceRackNumber, err := strconv.Atoi(deviceRackNumberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device_rack_number"})
			return
		}

		// Assign the values to the DeviceLocationDetail struct
		data = models.DeviceLocationDetail{
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
			logger.ErrorLogger.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create DeviceLocationDetail"})
			return
		}

		logger.InfoLogger.Println("New location details created successfully.")
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Entry Added Successfully"})
	}
}

// UpdateDeviceLocationDetail updates a DeviceLocationDetail record based on its ID.
func UpdateDeviceLocationDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		type RequestData struct {
			SerialNumber     string `json:"serial_number"`
			DeviceMakeModel  string `json:"device_make_model"`
			Model            string `json:"model"`
			DeviceType       string `json:"device_type"`
			DataCenter       string `json:"data_center"`
			Region           string `json:"region"`
			DCLocation       string `json:"dc_location"`
			DeviceLocation   string `json:"device_location"`
			DeviceRowNumber  string `json:"device_row_number"`  //int
			DeviceRackNumber string `json:"device_rack_number"` //int
			DeviceRUNumber   string `json:"device_ru_number"`
		}

		var requestData RequestData
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		row, _ := strconv.Atoi(requestData.DeviceRowNumber)
		rack, _ := strconv.Atoi(requestData.DeviceRackNumber)

		updatedData := &models.DeviceLocationDetail{
			Id:               id,
			SerialNumber:     requestData.SerialNumber,
			DeviceMakeModel:  requestData.DeviceMakeModel,
			Model:            requestData.Model,
			DeviceType:       requestData.DeviceType,
			DataCenter:       requestData.DataCenter,
			Region:           requestData.Region,
			DCLocation:       requestData.DCLocation,
			DeviceLocation:   requestData.DeviceLocation,
			DeviceRowNumber:  row,
			DeviceRackNumber: rack,
			DeviceRUNumber:   requestData.DeviceRUNumber,
		}

		if err := db.UpdateDeviceLocationDetail(id, updatedData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update DeviceLocationDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceLocationDetail updated successfully"})
	}
}

// DeleteDeviceLocationDetailHandler deletes a DeviceLocationDetail record based on its ID.
func DeleteDeviceLocationDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.DeleteDeviceLocationDetail(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete DeviceLocationDetail"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceLocationDetail deleted successfully"})
	}
}

func DownloadDeviceLocationDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Downloading DeviceLocationDetails as Excel file")
		// Query the database for DeviceLocationDetail data
		rows, err := db.Query("SELECT * FROM device_location")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to query the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a new Excel file
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("DeviceLocationDetails")
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
		headerRow.AddCell().SetString("Data Center")
		headerRow.AddCell().SetString("Region")
		headerRow.AddCell().SetString("DC Location")
		headerRow.AddCell().SetString("Device Location")
		headerRow.AddCell().SetString("Device Row Number")
		headerRow.AddCell().SetString("Device Rack Number")
		headerRow.AddCell().SetString("Device RU Number")

		// Add data rows from the database
		for rows.Next() {
			var device models.DeviceLocationDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.DataCenter, &device.Region, &device.DCLocation, &device.DeviceLocation, &device.DeviceRowNumber, &device.DeviceRackNumber, &device.DeviceRUNumber); err != nil {
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
			dataRow.AddCell().SetString(device.DataCenter)
			dataRow.AddCell().SetString(device.Region)
			dataRow.AddCell().SetString(device.DCLocation)
			dataRow.AddCell().SetString(device.DeviceLocation)
			dataRow.AddCell().SetInt(device.DeviceRowNumber)
			dataRow.AddCell().SetInt(device.DeviceRackNumber)
			dataRow.AddCell().SetString(device.DeviceRUNumber)
		}

		// Save the Excel file to the response
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=DeviceLocationDetails.xlsx")
		err = file.Write(c.Writer)
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to write Excel file to response", http.StatusInternalServerError)
		}
	}
}

func DownloadDeviceLocationDetailPDF(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Downloading DeviceLocationDetails as PDF")
		// Query the database for DeviceLocationDetail data
		rows, err := db.Query("SELECT * FROM device_location")
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
		headers := []string{"ID", "Serial Number", "Device Make Model", "Model", "Device Type", "Data Center", "Region", "DC Location", "Device Location", "Device Row Number", "Device Rack Number", "Device RU Number"}
		for _, header := range headers {
			pdf.CellFormat(40, 10, header, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)

		// Add data rows from the database
		for rows.Next() {
			var device models.DeviceLocationDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.DataCenter, &device.Region, &device.DCLocation, &device.DeviceLocation, &device.DeviceRowNumber, &device.DeviceRackNumber, &device.DeviceRUNumber); err != nil {
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
				device.DataCenter,
				device.Region,
				device.DCLocation,
				device.DeviceLocation,
				fmt.Sprint(device.DeviceRowNumber),
				fmt.Sprint(device.DeviceRackNumber),
				device.DeviceRUNumber,
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
		c.Header("Content-Disposition", "attachment; filename=DeviceLocationDetails.pdf")
	}
}
