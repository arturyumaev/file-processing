package middleware

import (
	"time"

	"github.com/arturyumaev/file-processing/api/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Logger() gin.HandlerFunc {
	logger := logger.GetLogger()

	return func(c *gin.Context) {
		timeStart := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		c.Next()

		timeAfter := time.Now()
		latency := timeAfter.Sub(timeStart)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		bodySize := c.Writer.Size()

		var logEvent *zerolog.Event
		if c.Writer.Status() >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}

		logEvent.
			Str("client_id", clientIP).
			Str("method", method).
			Int("status_code", statusCode).
			Int("body_size", bodySize).
			Str("path", path).
			Str("latency", latency.String()).
			Msg(errorMessage)
	}
}
