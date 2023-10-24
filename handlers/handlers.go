package handlers

import (
	"go-server/db"
	"go-server/models"
	"go-server/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
