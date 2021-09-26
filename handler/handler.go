package handler

import (
	"be-project/database"
	"be-project/models"
	"log"
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
	// connect to db
	db := database.ConnectDB()
	var dashboard models.Dashboard

	// query for collecting /dashboard data
	db.Raw("SELECT full_name, role_description, office_longitude, office_latitude FROM users AS a JOIN roles AS b ON a.role_id = b.role_id JOIN offices AS c ON a.office_id = c.office_id JOIN attendances AS d ON a.user_id = d.user_id WHERE a.user_id = 1 ORDER BY d.dates DESC LIMIT 1;").Scan(&dashboard)

	data := &models.Respon{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    dashboard,
	}
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

func PostCheckIn(c *gin.Context) {
	// connect to db
	db := database.ConnectDB()

	var body models.Checking

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Respon{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    body,
		})
		return
	}

	// change template for allowing gorm to work
	temp := &models.Attendances{
		UserId:        body.UserId,
		UserLatitude:  body.UserLatitude,
		UserLongitude: body.UserLongitude,
		CurrentDate:   body.CurrentDate,
		CheckIn:       body.CurrentDate,
		CheckOut:      "23:59:59",
	}

	// db config for insert data
	if err := db.Create(&temp).Error; err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNotAcceptable, models.Respon{
			Status:  http.StatusNotAcceptable,
			Message: http.StatusText(http.StatusNotAcceptable),
			Data:    temp,
		})
		return
	}

	c.JSON(http.StatusAccepted, models.Respon{
		Status:  http.StatusAccepted,
		Message: http.StatusText(http.StatusAccepted),
		Data:    temp,
	})
}

func PostCheckOut(c *gin.Context) {
	// connect to db
	db := database.ConnectDB()

	var body models.Checking

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Respon{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    body,
		})
		return
	}

	// db config for update column
	if err := db.Exec("UPDATE attendances SET checkout = '" + body.CurrentDate + "' WHERE user_id = 1 AND dates = (SELECT MAX(dates) FROM attendances)").Error; err != nil {
		c.JSON(http.StatusNotAcceptable, models.Respon{
			Status:  http.StatusNotAcceptable,
			Message: http.StatusText(http.StatusNotAcceptable),
			Data:    body,
		})
		return
	}

	c.JSON(http.StatusAccepted, models.Respon{
		Status:  http.StatusAccepted,
		Message: http.StatusText(http.StatusAccepted),
		Data:    body,
	})
}
