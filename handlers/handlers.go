package handlers

import (
	"fmt"
	"hello-world-http/m/v2/db"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	zlog "github.com/rs/zerolog/log"
	"github.com/tamathecxder/randomail"
	"github.com/tjarratt/babble"
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
	// handles /hostname path
	host, _ := os.Hostname()
	c.JSON(http.StatusOK, gin.H{
		"message": host,
	})
}

func Mask(c *gin.Context) {
	// handles /mask path
	pattern := c.Param("pattern")

	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = 10

	switch pattern {
	case "email":
		log.Printf("INFO: log sample %s: %s", pattern, randomail.GenerateRandomEmail())
		zlog.Info().Msgf("log sample %s: %s", pattern, randomail.GenerateRandomEmail())
	default:
		log.Printf("INFO: %s", babbler.Babble())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": pattern,
	})
}
