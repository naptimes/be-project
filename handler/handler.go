package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// template
const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

type User interface {
	UserID      int
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
	OfficeID    string
	RoleID      int
	Latitude    float32
	Longitude   float32
	CurrentDate string
	AuthToken   string
}

type test interface {
	Ini int
	Aku string 
}

type Timesheet struct {
}

func LandingPage(c *gin.Context) {
	c.Data(http.StatusOK, ContentTypeHTML, []byte("<h1>ini landing page</h1>"))
}

func GetDashboard(c *gin.Context) {
	user := &User{}
	c.JSON(http.StatusOK, ContentTypeJSON, user)
}

func GetTimesheet(c *gin.Context) {
	c.JSON(http.StatusOK, ContentTypeJSON)
}

func getAdministration(c *gin.Context) {
	c.JSON(http.StatusOK, ContentTypeJSON)
}
