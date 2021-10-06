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
	api.GET("/administration", handler.Auth, handler.GetAdministration)
	api.POST("/administration/approve", handler.ApproveUser)
	api.POST("/checkin", handler.Auth, handler.PostCheckIn)
	api.POST("/checkout", handler.Auth, handler.PostCheckOut)
	api.POST("/auth/register", handler.Register)
	api.POST("/auth/login", handler.Login)

	router.Use(cors.Default())
	router.Run()
}
