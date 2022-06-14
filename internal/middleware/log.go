package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/where-it-go/internal/logger"
	"time"
)

func GinRequestLog(c *gin.Context) {
	// start time
	startTime := time.Now()

	// next
	c.Next()

	// end time
	endTime := time.Now()

	// execute time
	latencyTime := endTime.Sub(startTime)

	// request method
	reqMethod := c.Request.Method

	// request uri
	reqUri := c.Request.RequestURI

	// status code
	statusCode := c.Writer.Status()

	// request ip
	clientIP := c.ClientIP()

	// write log
	logger.Info.Println(fmt.Sprintf(
		"| %3d | %13v | %15s | %s | %s |",
		statusCode, latencyTime, clientIP, reqMethod, reqUri),
	)
}
