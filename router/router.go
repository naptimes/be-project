package router

import (
	"be-project/handler"
	"time"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Run()
}
