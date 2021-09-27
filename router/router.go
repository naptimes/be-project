package router

import (
	"be-project/handler"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/", handler.LandingPage)
	api.GET("/dashboard", handler.GetDashboard)
	api.GET("/timesheet", handler.GetTimesheet)
	api.GET("/administration", handler.GetAdministration)
	api.POST("/checkin", handler.PostCheckIn)
	api.POST("/checkout", handler.PostCheckOut)
	api.POST("/auth/register", handler.Register)
	api.POST("/auth/login", handler.Login)

	router.Run()
}
