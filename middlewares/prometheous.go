package middlewares

import (
	prometheous "github.com/aruncs31s/gometrics/middlewares"
	"github.com/gin-gonic/gin"
)

func Metrics() gin.HandlerFunc {
	return prometheous.PrometheusMiddleware()
}
