package router

import (
	"be-project/handler"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/", CORSMiddleware(), handler.LandingPage)
	api.GET("/dashboard", CORSMiddleware(), handler.GetDashboard)
	api.GET("/timesheet", CORSMiddleware(), handler.GetTimesheet)
	api.GET("/administration", CORSMiddleware(), handler.GetAdministration)
	api.POST("/checkin", CORSMiddleware(), handler.PostCheckIn)
	api.POST("/checkout", CORSMiddleware(), handler.PostCheckOut)
	api.POST("/auth/register", CORSMiddleware(), handler.Register)
	api.POST("/auth/login", CORSMiddleware(), handler.Login)

	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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
