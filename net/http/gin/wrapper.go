package gin

import (
	"github.com/azdbaaaaaa/util/net/metadata"
	 "github.com/azdbaaaaaa/util/xutil/xerror"
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
		var code xerror.Error
		var rid string
		data, err := handle(c)
		if err != nil {
			if ec, ok := err.(xerror.Error); ok {
				code = ec
			} else {
				code = xerror.ErrUnknown
			}
		} else {
			code = xerror.Success
		}
		reqId, exists := c.Get(metadata.ContextKeyReqID)
		if exists {
			if v, ok := reqId.(string); ok {
				rid = v
			}
		}
		c.JSON(http.StatusOK, ApiError{
			Result:  code.GetCode(),
			Message: code.GetMessage(),
			Data:    data,
			Reason:  code.GetReason(),
			Rid:     rid,
		})
	}
}
