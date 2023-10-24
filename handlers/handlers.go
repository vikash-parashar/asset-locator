package handlers

import (
	"fmt"
	"go-server/db"
	"go-server/models"
	"go-server/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "application is running"})
}

func RenderIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func RenderHomePage(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	}
}

func RenderLoginUser(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func RenderRegisterUser(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

// RegisterUser handles the registration of a new user.
func SignUp(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve form data
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		role := c.PostForm("role")
		username := c.PostForm("username")

		// Check if the user already exists with the provided username or email
		if _, err := db.GetUserByUsername(username); err == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "user already exists with this username"})
			return
		}
		if _, err := db.GetUserByEmailID(email); err == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "user already exists with this email"})
			return
		}

		// Hash the password
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "failed to hash password"})
			return
		}

		// Create a new user with the hashed password
		newUser := models.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Role:      role,
			Password:  hashedPassword, // Store the hashed password
			Username:  username,
		}

		// Register the user in the database
		if err := db.RegisterUser(&newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "you are now registered"})
	}
}

// LoginUserHandler handles the user login and returns a JWT token upon successful login.
func Login(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve form data
		username := c.PostForm("username")
		password := c.PostForm("password")
		log.Println(username)
		log.Println(password)
		// Check if the user exists in the database
		user, err := db.GetUserByUsername(username)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User not found"})
			return
		}

		// Verify the password
		if !utils.VerifyPassword(password, user.Password) {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Incorrect password"})
			return
		}

		// Generate a JWT token
		token, err := utils.GenerateJWTToken(user)
		if err != nil {
			log.Println("Failed to generate JWT token")
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to generate JWT token"})
			return
		}

		cookie := http.Cookie{
			Name:    "jwt-token",
			Value:   token,
			Expires: time.Now().Add(5 * time.Minute),
		}
		http.SetCookie(c.Writer, &cookie)
		c.JSON(http.StatusOK, gin.H{"success": true, "token": token, "message": "Login successful"})
	}
}

// LogoutUserHandler handles the user logout by clearing the JWT token cookie.
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Clear the JWT token cookie
		cookie := http.Cookie{
			Name:    "jwt-token",
			Value:   "",
			Expires: time.Unix(0, 0),
		}
		http.SetCookie(c.Writer, &cookie)
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Logout successful"})
	}
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

// UpdateDeviceLocationDetailHandler updates a DeviceLocationDetail record based on its ID.
func UpdateDeviceLocationDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var data models.DeviceLocationDetail
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		if err := db.UpdateDeviceLocationDetail(id, &data); err != nil {
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

// DeleteDevicePowerDetailHandler deletes a DevicePowerDetail record based on its ID.
func DownloadDevicePowerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
			if err := rows.Scan(&device.ID, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.TotalPowerWatt, &device.TotalBTU, &device.TotalPowerCable, &device.PowerSocketType); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}
			dataRow := sheet.AddRow()
			dataRow.AddCell().SetInt(device.ID)
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
	}
}

func DownloadDeviceEthernetFiberDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Query the database for DeviceEthernetFiberDetail data (similar to the DevicePowerDetail function)
		rows, err := db.Query("SELECT * FROM device_ethernet_fiber")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to query the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a new Excel file
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("DeviceEthernetFiberDetails")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to create Excel sheet", http.StatusInternalServerError)
			return
		}

		// Add header row (similar to the DevicePowerDetail function)
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

		// Add data rows from the database (similar to the DevicePowerDetail function)
		for rows.Next() {
			var device models.DeviceEthernetFiberDetail
			if err := rows.Scan(&device.ID, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.DevicePhysicalPort, &device.DevicePortType, &device.DevicePortMACWWN, &device.ConnectedDevicePort); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}
			dataRow := sheet.AddRow()
			dataRow.AddCell().SetInt(device.ID)
			dataRow.AddCell().SetString(device.SerialNumber)
			dataRow.AddCell().SetString(device.DeviceMakeModel)
			dataRow.AddCell().SetString(device.Model)
			dataRow.AddCell().SetString(device.DeviceType)
			dataRow.AddCell().SetString(device.DevicePhysicalPort)
			dataRow.AddCell().SetString(device.DevicePortType)
			dataRow.AddCell().SetString(device.DevicePortMACWWN)
			dataRow.AddCell().SetString(device.ConnectedDevicePort)
		}

		// Save the Excel file to the response (similar to the DevicePowerDetail function)
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=DeviceEthernetFiberDetails.xlsx")
		err = file.Write(c.Writer)
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to write Excel file to response", http.StatusInternalServerError)
		}
	}
}

func DownloadDeviceAMCOwnerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Query the database for DeviceAMCOwnerDetail data
		rows, err := db.Query("SELECT * FROM device_amc_owner")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to query the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a new Excel file
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("DeviceAMCOwnerDetails")
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
		headerRow.AddCell().SetString("PO Number")
		headerRow.AddCell().SetString("PO Order Date")
		headerRow.AddCell().SetString("EOSL Date")
		headerRow.AddCell().SetString("AMC Start Date")
		headerRow.AddCell().SetString("AMC End Date")
		headerRow.AddCell().SetString("Device Owner")

		// Add data rows from the database
		for rows.Next() {
			var device models.DeviceAMCOwnerDetail
			if err := rows.Scan(&device.ID, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.PONumber, &device.POOrderDate, &device.EOSLDate, &device.AMCStartDate, &device.AMCEndDate, &device.DeviceOwner); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}
			dataRow := sheet.AddRow()
			dataRow.AddCell().SetInt(device.ID)
			dataRow.AddCell().SetString(device.SerialNumber)
			dataRow.AddCell().SetString(device.DeviceMakeModel)
			dataRow.AddCell().SetString(device.Model)
			dataRow.AddCell().SetString(device.PONumber)
			dataRow.AddCell().SetDate(device.POOrderDate)
			dataRow.AddCell().SetDate(device.EOSLDate)
			dataRow.AddCell().SetDate(device.AMCStartDate)
			dataRow.AddCell().SetDate(device.AMCEndDate)
			dataRow.AddCell().SetString(device.DeviceOwner)
		}

		// Save the Excel file to the response
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=DeviceAMCOwnerDetails.xlsx")
		err = file.Write(c.Writer)
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to write Excel file to response", http.StatusInternalServerError)
		}
	}
}

func DownloadDeviceLocationDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
			if err := rows.Scan(&device.ID, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.DataCenter, &device.Region, &device.DCLocation, &device.DeviceLocation, &device.DeviceRowNumber, &device.DeviceRackNumber, &device.DeviceRUNumber); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}
			dataRow := sheet.AddRow()
			dataRow.AddCell().SetInt(device.ID)
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
			if err := rows.Scan(&device.ID, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.DataCenter, &device.Region, &device.DCLocation, &device.DeviceLocation, &device.DeviceRowNumber, &device.DeviceRackNumber, &device.DeviceRUNumber); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}

			data := []string{
				fmt.Sprint(device.ID),
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

// DownloadDevicePowerDetailAsPDF exports DevicePowerDetail data as a PDF file.
func DownloadDevicePowerDetailPDF(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
			if err := rows.Scan(&device.ID, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.TotalPowerWatt, &device.TotalBTU, &device.TotalPowerCable, &device.PowerSocketType); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}

			data := []string{
				fmt.Sprint(device.ID),
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
	}
}

func DownloadDeviceEthernetFiberDetailPDF(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Query the database for DeviceEthernetFiberDetail data
		rows, err := db.Query("SELECT * FROM device_ethernet_fiber")
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
		headers := []string{"ID", "Serial Number", "Device Make Model", "Model", "Device Type", "Device Physical Port", "Device Port Type", "Device Port MAC/WWN", "Connected Device Port"}
		for _, header := range headers {
			pdf.CellFormat(40, 10, header, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)

		// Add data rows from the database
		for rows.Next() {
			var device models.DeviceEthernetFiberDetail
			if err := rows.Scan(&device.ID, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.DeviceType, &device.DevicePhysicalPort, &device.DevicePortType, &device.DevicePortMACWWN, &device.ConnectedDevicePort); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}

			data := []string{
				fmt.Sprint(device.ID),
				device.SerialNumber,
				device.DeviceMakeModel,
				device.Model,
				device.DeviceType,
				device.DevicePhysicalPort,
				device.DevicePortType,
				device.DevicePortMACWWN,
				device.ConnectedDevicePort,
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
		c.Header("Content-Disposition", "attachment; filename=DeviceEthernetFiberDetails.pdf")
	}
}

func DownloadDeviceAMCOwnerDetailPDF(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Query the database for DeviceAMCOwnerDetail data
		rows, err := db.Query("SELECT * FROM device_amc_owner")
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
		headers := []string{"ID", "Serial Number", "Device Make Model", "Model", "PO Number", "PO Order Date", "EOSL Date", "AMC Start Date", "AMC End Date", "Device Owner"}
		for _, header := range headers {
			pdf.CellFormat(40, 10, header, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)

		// Add data rows from the database
		for rows.Next() {
			var device models.DeviceAMCOwnerDetail
			if err := rows.Scan(&device.ID, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.PONumber, &device.POOrderDate, &device.EOSLDate, &device.AMCStartDate, &device.AMCEndDate, &device.DeviceOwner); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}

			data := []string{
				fmt.Sprint(device.ID),
				device.SerialNumber,
				device.DeviceMakeModel,
				device.Model,
				device.PONumber,
				device.POOrderDate.Format("2006-01-02"),
				device.EOSLDate.Format("2006-01-02"),
				device.AMCStartDate.Format("2006-01-02"),
				device.AMCEndDate.Format("2006-01-02"),
				device.DeviceOwner,
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
		c.Header("Content-Disposition", "attachment; filename=DeviceAMCOwnerDetails.pdf")
	}
}
