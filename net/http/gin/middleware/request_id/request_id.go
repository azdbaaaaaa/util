package request_id

import (
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestID(c *gin.Context) {
	reqID := uuid.NewV4().String()
	c.Set(metadata2.ContextKeyReqID, reqID)
	c.Next()
}
