package apex_gin

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/apex/log"
)

func Handler(message string) gin.HandlerFunc {
	var skip map[string]struct{}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			// Stop timer
			end := time.Now()
			latency := end.Sub(start)

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()

			log.WithFields(log.Fields{
				"status_code": statusCode,
				"latency": latency,
				"client_ip": clientIP,
				"method": method,
				"path": path,
			}).Info(message)
		}
	}
}
