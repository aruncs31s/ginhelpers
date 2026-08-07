package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	RequestIDHeader = "X-Request-ID"
	RequestIDKey    = "req_id"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(RequestIDHeader)

		// Generate one if the client didn't provide it.
		if requestID == "" {
			requestID = uuid.NewString()
		}

		// Make it available everywhere.
		c.Set(RequestIDKey, requestID)

		// Return it to the client.
		c.Writer.Header().Set(RequestIDHeader, requestID)

		c.Next()
	}
}
