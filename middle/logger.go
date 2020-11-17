package middle

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)


func Logger() gin.HandlerFunc {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return func(c *gin.Context) {
		sTime := time.Now()
		c.Next()
		aTime := time.Now().Sub(sTime)

		method := c.Request.Method
		path := c.Request.RequestURI
		clientIP := c.ClientIP()
		status := c.Writer.Status()

		logger.Infof("|%19v|%15s | %3d | %s | %s| %v",
			sTime.Format("2006-01-02 15:04:05"),
			clientIP,
			status,
			method,
			path,
			aTime,
			)
	}
}
