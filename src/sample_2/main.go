package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to sample_2 app",
		})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, this is sample_2 app",
		})
	})
	router.Run()
}
