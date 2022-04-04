package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ResponseFromSystem2 struct {
	message string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to sample_1 app",
		})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, this is sample_1 app",
		})
	})
	router.GET("/ping-system-2", func(c *gin.Context) {
		response, err := http.Get("http://sample_2:8080/ping")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
		c.JSON(200, string(responseData))
	})
	router.Run()
}
