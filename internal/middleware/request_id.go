package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	ContextKeyRequestID = "requestId"
	RequestIdHeaderName = "X-Request-ID"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.New().String()
		c.Set(ContextKeyRequestID, requestId)
		c.Writer.Header().Set(RequestIdHeaderName, requestId)

		c.Next()
	}
}
