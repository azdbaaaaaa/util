package request_id

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.NewV4().String()
		c.Set(ContextKeyReqID, reqID)
		c.Next()
	}
}
