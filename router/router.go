package router

import (
	"be-project/handler"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	router.GET("/", handler.LandingPage)

	router.Run()
}
