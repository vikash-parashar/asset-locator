package handlers

import (
	"log"
	"net/http"

	"github.com/vikash-parashar/asset-locator/utils"

	"github.com/gin-gonic/gin"
)

func FetchDisks(c *gin.Context) {
	log.Println("fetching disk's data")
	data, err := utils.FetchAndSendDisksInfo()
	if err != nil {
		log.Println("failed to fetch disk's data")
		log.Println(err)
	}
	log.Println("disk's data fetched successfully from external server.")
	log.Println("sending disk's data")

	c.JSON(http.StatusFound, gin.H{
		"message":      "disk's data fetched successfully",
		"disk's count": string(data),
	})
}
