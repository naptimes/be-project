package router

import (
	"be-project/handler"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/", handler.Auth, handler.LandingPage)
	api.GET("/dashboard", handler.Auth, handler.GetDashboard)
	api.GET("/timesheet", handler.Auth, handler.GetTimesheet)
	api.GET("/administration", handler.Auth, handler.GetAdministration)
	api.POST("/administration/approve", handler.ApproveUser)
	api.POST("/checkin", handler.Auth, handler.PostCheckIn)
	api.POST("/checkout", handler.Auth, handler.PostCheckOut)
	api.POST("/auth/register", handler.Register)
	api.POST("/auth/login", handler.Login)

	router.Use(CORSMiddleware())
	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://ipe8-workerattendance.herokuapp.com/")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
