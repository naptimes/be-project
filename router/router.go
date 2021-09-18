package router

import (
	"be-project/handler"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	router.GET("/", handler.LandingPage)
	router.GET("/dashboard", handler.GetDashboard)
	router.GET("/timesheet", handler.GetTimesheet)
	router.GET("/administration", handler.GetAdministration)

	router.Run()
}
