package handler

import (
	"be-project/database"
	"be-project/models"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
	jwtSecretKey      = "sangat-Rahasia!@#987"
)

var passClaims *jwt.StandardClaims

func LandingPage(c *gin.Context) {
	c.Data(http.StatusOK, ContentTypeHTML, []byte("<h1>ini landing page</h1>"))
}

func GetDashboard(c *gin.Context) {
	// connect to db
	db := database.ConnectDB()
	var dashboard models.Dashboard
	cookie := passClaims

	// query for collecting /dashboard data
	if err := db.Raw("SELECT full_name, role_description, office_longitude, office_latitude FROM users AS a JOIN roles AS b ON a.role_id = b.role_id JOIN offices AS c ON a.office_id = c.office_id WHERE a.user_id = " + cookie.Issuer + " LIMIT 1;").Scan(&dashboard).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Respon{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    dashboard,
		})
		return
	}

	data := &models.Respon{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    dashboard,
	}
	c.JSON(http.StatusOK, data)
}

func GetTimesheet(c *gin.Context) {
	// connect to db
	db := database.ConnectDB()
	var timesheet []models.Timesheet
	cookie := passClaims

	if err := db.Raw("SELECT a.dates, MIN(a.checkin) AS checkin, MAX(a.checkout) AS checkout, b.working_hours, b.attendance_status FROM attendances a JOIN work_hour b ON a.attendance_id = b.attendance_id WHERE a.user_id = " + cookie.Issuer + " GROUP BY a.dates;").Scan(&timesheet).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Respon{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    timesheet,
		})
		return
	}

	data := &models.Respon{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    timesheet,
	}
	c.JSON(http.StatusOK, data)
}

func GetAdministration(c *gin.Context) {
	db := database.ConnectDB()
	var userlist []models.Users

	if err := db.Select("user_id", "full_name", "email", "approval_status").Where("role_id = 2").Find(&userlist).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Respon{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    userlist,
		})
		return
	}

	data := &models.Respon{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    userlist,
	}
	c.JSON(http.StatusOK, data)
}

func ApproveUser(c *gin.Context) {
	db := database.ConnectDB()
	var body models.Users

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Respon{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    body,
		})
		return
	}

	//for updating the user email activation
	if err := db.Exec("UPDATE users SET approval_status = '1' WHERE email = '" + body.Email + "'").Error; err != nil {
		c.JSON(http.StatusNotAcceptable, models.Respon{
			Status:  http.StatusNotFound,
			Message: err.Error(),
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

func PostCheckIn(c *gin.Context) {
	// connect to db
	db := database.ConnectDB()

	var body models.Checking
	var dashboard models.Dashboard
	cookie := passClaims

	// query for collecting /dashboard data
	if err := db.Raw("SELECT full_name, role_description, office_longitude, office_latitude FROM users AS a JOIN roles AS b ON a.role_id = b.role_id JOIN offices AS c ON a.office_id = c.office_id WHERE a.user_id = " + cookie.Issuer + " LIMIT 1;").Scan(&dashboard).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Respon{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    dashboard,
		})
		return
	}

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

	id, _ := strconv.Atoi(cookie.Issuer)

	// change template for allowing gorm to work
	temp := &models.Attendances{
		UserId:        id,
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
		// log.Fatal(err)
		c.JSON(http.StatusNotAcceptable, models.Respon{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
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
	cookie := passClaims

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Respon{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    body,
		})
		return
	}

	id, _ := strconv.Atoi(cookie.Issuer)

	temp := &models.Checking{
		UserId:        id,
		UserLatitude:  body.UserLatitude,
		UserLongitude: body.UserLongitude,
		CurrentDate:   body.CurrentDate,
	}

	// query for collecting data
	if err := db.Raw("SELECT full_name, role_description, office_longitude, office_latitude FROM users AS a JOIN roles AS b ON a.role_id = b.role_id JOIN offices AS c ON a.office_id = c.office_id WHERE a.user_id = " + cookie.Issuer + " LIMIT 1;").Scan(&dashboard).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Respon{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    body,
		})
		return
	}

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
	if err := db.Exec("UPDATE attendances SET checkout = '" + body.CurrentDate + "' WHERE user_id = " + cookie.Issuer + " AND dates = (SELECT MAX(dates) FROM attendances)").Error; err != nil {
		c.JSON(http.StatusNotAcceptable, models.Respon{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
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

func Register(c *gin.Context) {
	// connect to db
	db := database.ConnectDB()
	var body models.Register

	// JWT token func here

	// collect data from body
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Respon{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    body,
		})
		return
	}

	// password hashing here
	password, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	// change template
	temp := &models.User{
		FullName:       body.FullName,
		Email:          body.Email,
		Password:       password,
		RoleID:         2,
		OfficeID:       "IT001",
		ApprovalStatus: false,
	}

	// insert new user information to db
	if err := db.Create(&temp).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, models.Respon{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
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

func Login(c *gin.Context) {
	// connect to db
	db := database.ConnectDB()
	var body models.Login
	// var catch models.User
	var user models.User

	// collect data from body
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Respon{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    body,
		})
		return
	}

	// get auth token from header here

	// check user information using email, password
	if err := db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Respon{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    body,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, models.Respon{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    body,
		})
		return
	}

	if user.ApprovalStatus == false {
		c.JSON(http.StatusNotFound, models.Respon{
			Status:  http.StatusNotFound,
			Message: "This email is not approved yet!",
			Data:    body,
		})
		return
	}

	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.UserID),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // a week
	})

	token, err := sign.SignedString([]byte(jwtSecretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Respon{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    body,
		})
		return
	}

	c.SetCookie("jwt", token, 86400, "", "", false, false)

	c.JSON(http.StatusOK, models.Respon{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}

// func for authenticate user
func Auth(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Respon{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})

	if token != nil && err == nil {
		claims := token.Claims.(*jwt.StandardClaims)
		passClaims = claims
	} else {
		c.JSON(http.StatusUnauthorized, models.Respon{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
		c.Abort()
		return
	}
}
