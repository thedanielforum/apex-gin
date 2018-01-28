package apex_gin

import (
	"time"

	"github.com/apex/log"
	"github.com/gin-gonic/gin"
)

// Handler gin middleware handler
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
				"latency":     latency,
				"client_ip":   clientIP,
				"method":      method,
				"path":        path,
			}).Info(message)
		}
	}
}
