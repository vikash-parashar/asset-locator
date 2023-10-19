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

func AddData(c *gin.Context) {

}
