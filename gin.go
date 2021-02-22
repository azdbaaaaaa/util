package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Code    Code                   `json:"code"`
	Msg     string                 `json:"msg"`
	Success bool                   `json:"success"`
	Result  interface{}            `json:"result"`
	Tid     string                 `json:"tid"`
	Ext     map[string]interface{} `json:"ext,omitempty"`
}

func Success(c *gin.Context, data interface{}, ext map[string]interface{}) {
	c.JSON(http.StatusOK, CommonResponse{
		Code:    CodeSuccess,
		Msg:     CodeSuccess.Error(),
		Success: CodeSuccess.IsSuccess(),
		Result:  data,
		Tid:     "",
		Ext:     ext,
	})
}

func Failure(c *gin.Context, err error, ext map[string]interface{}) {
	var code Code
	if c, ok := err.(Code); ok {
		code = c
	} else {
		code = CodeUnknown
	}
	c.JSON(http.StatusOK, CommonResponse{
		Code:    code,
		Msg:     code.Error(),
		Success: code.IsSuccess(),
		Result:  nil,
		Tid:     "",
		Ext:     ext,
	})
}
