package handlers

import (
	"net/http" // Update with your actual import path

	"github.com/gin-gonic/gin"
	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/utils"
)

func FetchDisks(c *gin.Context) {
	logger.InfoLogger.Println("Fetching disk's data")

	data, err := utils.FetchAndSendDisksInfo()
	if err != nil {
		logger.ErrorLogger.Println("Failed to fetch disk's data")
		logger.ErrorLogger.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch disk's data",
		})
		return
	}

	logger.InfoLogger.Println("Disk's data fetched successfully from external server.")
	logger.InfoLogger.Println("Sending disk's data")

	c.JSON(http.StatusOK, gin.H{
		"message":      "Disk's data fetched successfully",
		"disk's count": string(data),
	})
}
