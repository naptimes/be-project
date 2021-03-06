package router

import (
	"be-project/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/", handler.Auth, handler.LandingPage)
	api.GET("/dashboard", handler.Auth, handler.GetDashboard)
	api.GET("/timesheet", handler.Auth, handler.GetTimesheet)
	api.GET("/administration", handler.GetAdministration)
	api.GET("/administration/attendance", handler.GetUsersAttendances)
	api.POST("/administration/approve", handler.ApproveUser)
	api.POST("/checkin", handler.Auth, handler.PostCheckIn)
	api.POST("/checkout", handler.Auth, handler.PostCheckOut)
	api.POST("/auth/register", handler.Register)
	api.POST("/auth/login", handler.Login)

	// router.Use(cors.New(cors.Config{
	// 	//AllowOrigins:     []string{"*"},
	// 	AllowAllOrigins: true,
	// 	AllowMethods:    []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTION"},
	// 	AllowHeaders:    []string{"Accept, Accept-Language, Content-Type, YourOwnHeader"},
	// 	// AllowHeaders:     []string{"Access-Control-Allow-Headers ,Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
	// 	ExposeHeaders:    []string{"*"},
	// 	AllowCredentials: false,
	// 	MaxAge:           12 * time.Hour,
	// }))

	router.Use(cors.Default())
	router.Run()
}
