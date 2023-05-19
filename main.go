package main

import (
	"hello-world-http/m/v2/db"
	"hello-world-http/m/v2/handlers"
	"log"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ListenAddr = ":8080"
	DbAddr     = "localhost:6379"
)

func get_env(key string, def ...string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		val = def[0]
	}
	return val
}

func main() {

	var err error

	db.DbClient, err = db.NewDatabase(get_env("REDISURL", DbAddr))
	if err != nil {
		log.Printf("WARN: Failed to connect to redis: %s", err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	router.GET("/hello/:name", handlers.Hello)
	router.GET("/hostname", handlers.Hostname)

	router.Run(ListenAddr)
}
