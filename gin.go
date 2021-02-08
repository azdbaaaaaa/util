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
		Msg:     CodeSuccess.Message(),
		Success: CodeSuccess.IsSuccess(),
		Result:  data,
		Tid:     "",
		Ext:     ext,
	})
}

func Failure(c *gin.Context, code Code, ext map[string]interface{}) {
	c.JSON(http.StatusOK, CommonResponse{
		Code:    code,
		Msg:     code.Message(),
		Success: code.IsSuccess(),
		Result:  nil,
		Tid:     "",
		Ext:     ext,
	})
}
