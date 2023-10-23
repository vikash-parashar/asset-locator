package handlers

import (
	"go-server/db"
	"net/http"

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
