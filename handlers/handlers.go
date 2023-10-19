package handlers

import (
	"go-server/render"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {

	// display dashboard
	render.RenderTemplate(c.Writer, "index", nil)
}

func GetData(c *gin.Context) {

}
func GetLocationDetails(c *gin.Context) {

}
func GetOwnerDetails(c *gin.Context) {

}
func GetPowerDetails(c *gin.Context) {

}
func GetFiberDetails(c *gin.Context) {

}
func CreateNewLocationDetails(c *gin.Context) {

}
func CreateNewOwnerDetails(c *gin.Context) {

}
func CreateNewPowerDetails(c *gin.Context) {

}
func CreateNewFiberDetails(c *gin.Context) {

}
