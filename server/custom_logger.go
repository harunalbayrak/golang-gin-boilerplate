package server

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Custom logger
func CustomLogger(ctx *gin.Context) {
	// Starting time request
	startTime := time.Now()

	// Processing request
	ctx.Next()

	// End Time request
	endTime := time.Now()

	// execution time
	latencyTime := endTime.Sub(startTime)

	// Request method
	reqMethod := ctx.Request.Method

	// Request route
	reqUri := ctx.Request.RequestURI

	// status code
	statusCode := ctx.Writer.Status()

	// Request IP
	clientIP := ctx.ClientIP()

	logger.Infow("HTTP Request",
		"method", reqMethod,
		"uri", reqUri,
		"status", statusCode,
		"latency", latencyTime*time.Millisecond,
		"ip", clientIP,
	)
}
