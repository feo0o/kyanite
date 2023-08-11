package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DelayBySecond() gin.HandlerFunc {
	return func(c *gin.Context) {
		delay := c.DefaultQuery("delay", "0")
		d, err := strconv.Atoi(delay)
		if err != nil || d < 0 {
			c.String(http.StatusBadRequest, fmt.Sprintf("invalid query value: %s, %s\n", delay, err))
			return
		}

		if d > 0 {
			time.Sleep(time.Duration(d) * time.Second)
		}

		c.Next()
	}
}
