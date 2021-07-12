package gin

import (
	request_id2 "github.com/azdbaaaaaa/util/net/middleware/request_id"
	xerror2 "github.com/azdbaaaaaa/util/xutil/xerror"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiError struct {
	// 错误码 ex:999999
	Result int32 `json:"result"`
	// 错误消息
	Message string `json:"message"`
	// 数据
	Data interface{} `json:"data"`

	// error debug信息
	Reason string `json:"reason,omitempty"`
	// request id
	Rid string `json:"rid"`
}

type WrapperHandle func(c *gin.Context) (interface{}, error)

func ErrorWrapper(handle WrapperHandle) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code *xerror2.Error
		var rid string
		data, err := handle(c)
		if err != nil {
			if ec, ok := err.(*xerror2.Error); ok {
				code = ec
			} else {
				code = xerror2.ErrUnknown
			}
		} else {
			code = xerror2.Success
		}
		reqId, exists := c.Get(request_id2.ContextKeyReqID)
		if exists {
			if v, ok := reqId.(string); ok {
				rid = v
			}
		}
		c.JSON(http.StatusOK, ApiError{
			Result:  code.Code,
			Message: code.Message,
			Data:    data,
			Reason:  code.Reason,
			Rid:     rid,
		})
	}
}
