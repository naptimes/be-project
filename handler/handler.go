package handler

import (
	"be-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

func LandingPage(c *gin.Context) {
	c.Data(http.StatusOK, ContentTypeHTML, []byte("<h1>ini landing page</h1>"))
}

func GetDashboard(c *gin.Context) {
	// collect from db
	data := &models.Respon{} // not yet

	c.JSON(http.StatusOK, data)
}

func GetTimesheet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"testing": 123,
	})
}

func GetAdministration(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"aku": "orang",
	})
}
