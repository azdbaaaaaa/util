package grpc_request_id

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestID(c *gin.Context) {
	reqID := uuid.NewV4().String()
	c.Set(ContextKeyReqID, reqID)
	c.Next()
}
