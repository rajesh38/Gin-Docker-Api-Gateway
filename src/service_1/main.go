package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ResponseFromSystem2 struct {
	message string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to service_1 app",
		})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, this is service_1 app",
		})
	})
	router.GET("/ping-system-2", func(c *gin.Context) {
		response, err := http.Get("http://service_2:8080/ping")

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
	router.GET("/download", func(c *gin.Context) {
		fileName := "file.png"
		targetPath := filepath.Join("./assets/", fileName)
		fmt.Println("targetPath = " + targetPath)
		// Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Type", "image/png")
		c.File(targetPath)
	})
	router.Run()
}
