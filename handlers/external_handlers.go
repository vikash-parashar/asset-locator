package handlers

import (
	"fmt"
	"net/http" // Update with your actual import path

	"github.com/gin-gonic/gin"
	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/utils"
)

// func FetchDisks(c *gin.Context) {
// 	logger.InfoLogger.Println("Fetching disk's data")

// 	data, err := utils.FetchAndSendDisksInfo()
// 	if err != nil {
// 		logger.ErrorLogger.Println("Failed to fetch disk's data")
// 		logger.ErrorLogger.Println(err)

// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to fetch disk's data",
// 		})
// 		return
// 	}

// 	logger.InfoLogger.Println("Disk's data fetched successfully from external server.")
// 	logger.InfoLogger.Println("Sending disk's data")

// 	c.JSON(http.StatusOK, gin.H{
// 		"message":      "Disk's data fetched successfully",
// 		"disk's count": string(data),
// 	})
// }

func FetchDisks(c *gin.Context) {
	logger.InfoLogger.Println("Fetching disk's data from external server")

	// Replace the URL with the actual URL of your external server and endpoint
	externalServerURL := "http://localhost:8090/get-disk-count"

	resp, err := http.Get(externalServerURL)
	if err != nil {
		logger.ErrorLogger.Println("Failed to fetch disk's data from external server")
		logger.ErrorLogger.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch disk's data from external server",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.ErrorLogger.Printf("Failed to fetch disk's data from external server. Status code: %d\n", resp.StatusCode)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch disk's data from external server. Status code: %d", resp.StatusCode),
		})
		return
	}

	var responseData map[string]interface{}
	err = utils.ParseJSONResponse(resp.Body, &responseData)
	if err != nil {
		logger.ErrorLogger.Println("Failed to parse disk count from external server response")
		logger.ErrorLogger.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse disk count from external server response",
		})
		return
	}

	// Assuming the key in the response is "disk_count"
	diskCount, ok := responseData["disk_count"].(float64)
	if !ok {
		logger.ErrorLogger.Println("Failed to extract disk count from external server response")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to extract disk count from external server response",
		})
		return
	}

	logger.InfoLogger.Printf("Disk count fetched successfully from external server: %d\n", int(diskCount))

	c.JSON(http.StatusOK, gin.H{
		"message":      "Disk's data fetched successfully",
		"disk's count": int(diskCount),
	})
}
