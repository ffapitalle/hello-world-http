package handlers

import (
	"fmt"
	"hello-world-http/m/v2/db"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	// handles /hello/:name path
	visitor := c.Param("name")

	msg := fmt.Sprintf("hello %s", visitor)

	if db.DbClient != nil {
		seen, _ := db.DbClient.AddVisit(visitor)
		switch seen {
		case 1:
			msg = msg + ", nice to meet you"
		default:
			msg = msg + fmt.Sprintf(", glad to have you back (%d times)", seen)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})

}

func Hostname(c *gin.Context) {
	// handles /netinfo path
	host, _ := os.Hostname()
	c.JSON(http.StatusOK, gin.H{
		"message": host,
	})
}
