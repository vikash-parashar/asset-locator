package handlers

import (
	"net/http"
	"strconv" // Update with your actual import path

	"github.com/vikash-parashar/asset-locator/db"
	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/models"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
)

func GetFiberDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Fetching fiber details from the database")
		data, err := db.GetAllDeviceEthernetFiberDetail()
		if err != nil {
			logger.ErrorLogger.Println(err)
			c.HTML(http.StatusOK, "fiber_details.html", gin.H{"data": nil})
			return
		}
		logger.InfoLogger.Println("Fiber details fetched successfully.")
		c.HTML(http.StatusOK, "fiber_details.html", gin.H{"data": data})
	}
}

func GetFiberDetailByID(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.ErrorLogger.Println("Invalid ID:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		logger.InfoLogger.Println("Fetching fiber detail by ID:", id)
		fiberDetail, err := db.GetFiberDetailByID(id)
		if err != nil {
			logger.ErrorLogger.Println("Fiber detail not found:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Fiber detail not found"})
			return
		}

		logger.InfoLogger.Println("Fiber detail fetched successfully.")
		c.JSON(http.StatusOK, fiberDetail)
	}
}

func CreateNewFiberDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.DeviceEthernetFiberDetail

		logger.InfoLogger.Println("Creating new fiber details")

		// Retrieve form data
		serialNumber := c.PostForm("serial_number")
		deviceMakeModel := c.PostForm("device_make_model")
		model := c.PostForm("model")
		deviceType := c.PostForm("device_type")
		devicePhysicalPort := c.PostForm("device_physical_port")
		devicePortType := c.PostForm("device_port_type")
		devicePortMACWWN := c.PostForm("device_port_macwwn")
		connectedDevicePort := c.PostForm("connected_device_port")

		data = models.DeviceEthernetFiberDetail{
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
			logger.ErrorLogger.Println(err)
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "Failed to create entry"})
			return
		}

		logger.InfoLogger.Println("New fiber details created successfully.")
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Entry Added Successfully"})
	}
}

// func UpdateDeviceEthernetFiberDetail(db *db.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		type DeviceEthernetFiberDetail struct {
// 			Id                  string `json:"id"`
// 			SerialNumber        string `json:"serial_number"`
// 			DeviceMakeModel     string `json:"device_make_model"`
// 			Model               string `json:"model"`
// 			DeviceType          string `json:"device_type"`
// 			DevicePhysicalPort  string `json:"device_physical_port"`
// 			DevicePortType      string `json:"device_port_type"`
// 			DevicePortMACWWN    string `json:"device_port_macwwn"`
// 			ConnectedDevicePort string `json:"connected_device_port"`
// 		}
// 		var r DeviceEthernetFiberDetail
// 		if err := c.BindJSON(&r); err != nil {
// 			logger.ErrorLogger.Println("Invalid JSON data:", err)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
// 			return
// 		}

// 		nid, _ := strconv.Atoi(r.Id)

// 		updatedData := &models.DeviceEthernetFiberDetail{
// 			Id:                  nid,
// 			SerialNumber:        r.SerialNumber,
// 			DeviceMakeModel:     r.DeviceMakeModel,
// 			Model:               r.Model,
// 			DeviceType:          r.DeviceType,
// 			DevicePhysicalPort:  r.DevicePhysicalPort,
// 			DevicePortType:      r.DevicePortType,
// 			DevicePortMACWWN:    r.DevicePortMACWWN,
// 			ConnectedDevicePort: r.ConnectedDevicePort,
// 		}

// 		if err := db.UpdateDeviceEthernetFiberDetail(nid, updatedData); err != nil {
// 			logger.ErrorLogger.Println("Failed to update DeviceEthernetFiberDetail:", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update DeviceEthernetFiberDetail"})
// 			return
// 		}

// 		logger.InfoLogger.Println("DeviceEthernetFiberDetail updated successfully.")
// 		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceEthernetFiberDetail updated successfully"})
// 	}
// }

func UpdateDeviceEthernetFiberDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			logger.ErrorLogger.Println(err)
			return
		}

		type UpdateEthernetFiberDetailRequest struct {
			SerialNumber        string `json:"serial_number"`
			DeviceMakeModel     string `json:"device_make_model"`
			Model               string `json:"model"`
			DeviceType          string `json:"device_type"`
			DevicePhysicalPort  string `json:"device_physical_port"`
			DevicePortType      string `json:"device_port_type"`
			DevicePortMACWWN    string `json:"device_port_macwwn"`
			ConnectedDevicePort string `json:"connected_device_port"`
		}

		var request UpdateEthernetFiberDetailRequest
		if err := c.BindJSON(&request); err != nil {
			logger.ErrorLogger.Println("Invalid JSON data:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			return
		}

		// Retrieve existing data
		existingData, err := db.GetFiberDetailByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch existing DeviceEthernetFiberDetail"})
			logger.ErrorLogger.Println(err)
			return
		}

		// Update only non-null fields
		if request.SerialNumber != "" {
			existingData.SerialNumber = request.SerialNumber
		}
		if request.DeviceMakeModel != "" {
			existingData.DeviceMakeModel = request.DeviceMakeModel
		}
		if request.Model != "" {
			existingData.Model = request.Model
		}
		if request.DeviceType != "" {
			existingData.DeviceType = request.DeviceType
		}
		if request.DevicePhysicalPort != "" {
			existingData.DevicePhysicalPort = request.DevicePhysicalPort
		}
		if request.DevicePortType != "" {
			existingData.DevicePortType = request.DevicePortType
		}
		if request.DevicePortMACWWN != "" {
			existingData.DevicePortMACWWN = request.DevicePortMACWWN
		}
		if request.ConnectedDevicePort != "" {
			existingData.ConnectedDevicePort = request.ConnectedDevicePort
		}

		// You can add additional validation or handle errors here

		if err := db.UpdateDeviceEthernetFiberDetail(id, &existingData); err != nil {
			logger.ErrorLogger.Println("Failed to update DeviceEthernetFiberDetail:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update DeviceEthernetFiberDetail"})
			return
		}

		logger.InfoLogger.Println("DeviceEthernetFiberDetail updated successfully.")
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceEthernetFiberDetail updated successfully"})
	}
}

func DeleteDeviceEthernetFiberDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.ErrorLogger.Println("Invalid ID:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.DeleteDeviceEthernetFiberDetail(id); err != nil {
			logger.ErrorLogger.Println("Failed to delete DeviceEthernetFiberDetail:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete DeviceEthernetFiberDetail"})
			return
		}

		logger.InfoLogger.Println("DeviceEthernetFiberDetail deleted successfully.")
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "DeviceEthernetFiberDetail deleted successfully"})
	}
}

func DownloadDeviceEthernetFiberDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Downloading DeviceEthernetFiberDetails as Excel file")
		rows, err := db.Query("SELECT * FROM device_ethernet_fiber")
		if err != nil {
			logger.ErrorLogger.Println("Failed to query the database:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database"})
			return
		}
		defer rows.Close()

		file := xlsx.NewFile()
		sheet, err := file.AddSheet("DeviceEthernetFiberDetails")
		if err != nil {
			logger.ErrorLogger.Println("Failed to create Excel sheet:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Excel sheet"})
			return
		}

		headerRow := sheet.AddRow()
		headerRow.AddCell().SetString("ID")
		headerRow.AddCell().SetString("Serial Number")
		headerRow.AddCell().SetString("Device Make Model")
		headerRow.AddCell().SetString("Model")
		headerRow.AddCell().SetString("Device Type")
		headerRow.AddCell().SetString("Device Physical Port")
		headerRow.AddCell().SetString("Device Port Type")
		headerRow.AddCell().SetString("Device Port MAC/WWN")
		headerRow.AddCell().SetString("Connected Device Port")

		for rows.Next() {
			var device models.DeviceEthernetFiberDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.DevicePhysicalPort, &device.DevicePortType, &device.DevicePortMACWWN, &device.ConnectedDevicePort); err != nil {
				logger.ErrorLogger.Println("Failed to scan database row:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan database row"})
				return
			}
			dataRow := sheet.AddRow()
			dataRow.AddCell().SetInt(device.Id)
			dataRow.AddCell().SetString(device.SerialNumber)
			dataRow.AddCell().SetString(device.DeviceMakeModel)
			dataRow.AddCell().SetString(device.Model)
			dataRow.AddCell().SetString(device.DeviceType)
			dataRow.AddCell().SetString(device.DevicePhysicalPort)
			dataRow.AddCell().SetString(device.DevicePortType)
			dataRow.AddCell().SetString(device.DevicePortMACWWN)
			dataRow.AddCell().SetString(device.ConnectedDevicePort)
		}

		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Header("Content-Disposition", "attachment; filename=DeviceEthernetFiberDetails.xlsx")
		err = file.Write(c.Writer)
		if err != nil {
			logger.ErrorLogger.Println("Failed to write Excel file to response:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write Excel file to response"})
		}
	}
}

func DownloadDeviceEthernetFiberDetailPDF(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Downloading DeviceEthernetFiberDetails as PDF")
		rows, err := db.Query("SELECT * FROM device_ethernet_fiber")
		if err != nil {
			logger.ErrorLogger.Println("Failed to query the database:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database"})
			return
		}
		defer rows.Close()

		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

		pdf.SetFont("Arial", "", 12)

		headers := []string{"ID", "Serial Number", "Device Make Model", "Model", "Device Type", "Device Physical Port", "Device Port Type", "Device Port MAC/WWN", "Connected Device Port"}
		colWidths := []float64{10, 30, 50, 20, 30, 30, 30, 50, 50}

		for i, header := range headers {
			pdf.CellFormat(colWidths[i], 10, header, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)

		for rows.Next() {
			var device models.DeviceEthernetFiberDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.DevicePhysicalPort, &device.DevicePortType, &device.DevicePortMACWWN, &device.ConnectedDevicePort); err != nil {
				logger.ErrorLogger.Println("Failed to scan database row:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan database row"})
				return
			}

			data := []string{
				strconv.Itoa(device.Id),
				device.SerialNumber,
				device.DeviceMakeModel,
				device.Model,
				device.DeviceType,
				device.DevicePhysicalPort,
				device.DevicePortType,
				device.DevicePortMACWWN,
				device.ConnectedDevicePort,
			}

			for i, str := range data {
				pdf.CellFormat(colWidths[i], 10, str, "1", 0, "", false, 0, "")
			}
			pdf.Ln(-1)
		}

		c.Header("Content-Type", "application/pdf")
		c.Header("Content-Disposition", "attachment; filename=DeviceEthernetFiberDetails.pdf")
		err = pdf.Output(c.Writer)
		if err != nil {
			logger.ErrorLogger.Println("Failed to write PDF file to response:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write PDF file to response"})
		}
	}
}
