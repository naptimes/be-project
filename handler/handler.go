package handler

import (
	"be-project/database"
	"be-project/models"
	"log"
	"math"
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
	var dashboard models.Dashboard

	// query for collecting /dashboard data
	db.Raw("SELECT full_name, role_description, office_longitude, office_latitude FROM users AS a JOIN roles AS b ON a.role_id = b.role_id JOIN offices AS c ON a.office_id = c.office_id JOIN attendances AS d ON a.user_id = d.user_id WHERE a.user_id = 1 ORDER BY d.dates DESC LIMIT 1;").Scan(&dashboard)

	officeLat := dashboard.OfficeLatitude
	officeLong := dashboard.OfficeLongitude

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
		UserId:        1,
		UserLatitude:  body.UserLatitude,
		UserLongitude: body.UserLongitude,
		CurrentDate:   body.CurrentDate,
		CheckIn:       body.CurrentDate,
		CheckOut:      "23:59:59",
	}

	distance := CalculateDistance(temp.UserLatitude, temp.UserLongitude, officeLat, officeLong)
	if distance > 100 {
		c.JSON(http.StatusNotAcceptable, models.Respon{
			Status:  http.StatusNotAcceptable,
			Message: http.StatusText(http.StatusNotAcceptable),
			Data:    temp,
		})
		return
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
	var dashboard models.Dashboard

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Respon{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    body,
		})
		return
	}

	temp := &models.Checking{
		UserId:        1,
		UserLatitude:  body.UserLatitude,
		UserLongitude: body.UserLongitude,
		CurrentDate:   body.CurrentDate,
	}

	// query for collecting data
	db.Raw("SELECT full_name, role_description, office_longitude, office_latitude FROM users AS a JOIN roles AS b ON a.role_id = b.role_id JOIN offices AS c ON a.office_id = c.office_id JOIN attendances AS d ON a.user_id = d.user_id WHERE a.user_id = 1 ORDER BY d.dates DESC LIMIT 1;").Scan(&dashboard)

	officeLat := dashboard.OfficeLatitude
	officeLong := dashboard.OfficeLongitude

	// For logging the coordinates
	// fmt.Println("office lat:", officeLat)
	// fmt.Println("office long:", officeLong)
	// fmt.Println("user lat:", temp.UserLatitude)
	// fmt.Println("user long:", temp.UserLongitude)

	distance := CalculateDistance(temp.UserLatitude, temp.UserLongitude, officeLat, officeLong)
	if distance > 100 {
		c.JSON(http.StatusNotAcceptable, models.Respon{
			Status:  http.StatusNotAcceptable,
			Message: http.StatusText(http.StatusNotAcceptable),
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

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func CalculateDistance(userLat, userLong, officeLat, officeLong float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = userLat * math.Pi / 180
	lo1 = userLong * math.Pi / 180
	la2 = officeLat * math.Pi / 180
	lo2 = officeLong * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)
	d := 2 * r * math.Asin(math.Sqrt(h))

	//Logging the Distance
	//fmt.Println("Distance:", d)

	return d
}
