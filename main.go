package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func handle_hello(c *gin.Context) {
	// handles /hello/:name path
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("hello %s", c.Param("name")),
	})

}

func handle_hostname(c *gin.Context) {
	// handles /netinfo path
	host, _ := os.Hostname()
	c.JSON(http.StatusOK, gin.H{
		"message": host,
	})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	router.GET("/hello/:name", handle_hello)
	router.GET("/hostname", handle_hostname)

	router.Run(":8080")
}
