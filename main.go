package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"username": "iniTesting",
			"password": "iniPassword",
		})
	})

	router.GET("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"username": "iniTesting",
			"password": "iniPassword",
			"email":    "ini@email.com",
			"number":   "0812-xxxx-xxxx",
			"fname":    "ali",
			"lname":    "basyaeb",
		})
	})

	router.Run()
}
