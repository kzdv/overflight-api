package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "hawton.dev/log4g"
)

func Logger(c *gin.Context) {
	start := time.Now()
	c.Next()
	end := time.Now()
	latency := end.Sub(start)

	log.Category("web").Info(fmt.Sprintf("%s - %s %s - %d \"%s\" %s", c.ClientIP(), c.Request.Method, c.Request.URL.Path, c.Writer.Status(), c.Request.UserAgent(), latency))
}
