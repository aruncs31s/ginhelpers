package helpers

import (
	"github.com/aruncs31s/ginhelpers/middlewares"
	"github.com/gin-gonic/gin"
)

func GetRequestID(c *gin.Context) string {
	if id, ok := c.Get(middlewares.RequestIDKey); ok {
		if s, ok := id.(string); ok {
			return s
		}
	}
	return ""
}
