package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiError struct {
	// 错误码 ex:999999
	Result int `json:"result"`
	// 错误消息
	Message string `json:"message"`
	// 数据
	Data string `json:"data"`

	// error debug信息
	Err string `json:"err,omitempty"`
	// request id
	Rid string `json:"rid"`
}

func (err ApiError) Error() string {
	return err.Message
}

type WrapperHandle func(c *gin.Context) (interface{}, error)

func ErrorWrapper(handle WrapperHandle) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apiError)
			return
		}
		apiError := data.(ApiError)
		c.JSON(http.StatusOK, apiError)
	}
}
