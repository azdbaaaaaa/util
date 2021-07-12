package middleware

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Gzip() gin.HandlerFunc {
	return func(c *gin.Context) {
		gzip.Gzip(gzip.DefaultCompression)
	}
}
