package ginhelpers

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	dbKey = "DB"
	once  sync.Once
)

func GetDB[T any](c *gin.Context) T {
	return c.MustGet("DB").(T)
}

func SetDBKey(key string) {
	once.Do(func() {
		dbKey = key
	})
}
