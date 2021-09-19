package handler

import (
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

type User struct {
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

type Respawns struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    Dashboard `json:"data"`
}

type Dashboard struct {
	FullName        string  `json:"full_name"`
	Role            string  `json:"role"`
	OfficeLongitude float32 `json:"office_longitude"`
	OfficeLatitude  float32 `json:"office_latitude"`
	CurrentDate     string  `json:"current_date"`
	UserLongitude   float32 `json:"user_longitude"`
	UserLatitude    float32 `json:"user_latitude"`
}

type Timesheet struct {
}

func LandingPage(c *gin.Context) {
	c.Data(http.StatusOK, ContentTypeHTML, []byte("<h1>ini landing page</h1>"))
}

func GetDashboard(c *gin.Context) {
	// collect from db
	data := &Respawns{} // not yet

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
