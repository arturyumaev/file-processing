package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/arturyumaev/file-processing/internal/file_info"
)

func RequestTimeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// set context with timeout to request instead of default git context
		c.Request = c.Request.WithContext(ctx)
		done := make(chan bool)

		go func() {
			c.Next()
			done <- true
		}()

		select {
		case <-ctx.Done(): // timeout reached
			c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{
				"error": file_info.ErrRequestTimeoutReached.Error(),
			})
			return
		case <-done:
			return
		}
	}
}
