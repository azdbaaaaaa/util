package util

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	code := CodeUnknown

	if c, ok := err.(Code); ok {
		code = c
	}

	if c, ok := err.(validator.ValidationErrors); ok {
		var buffer bytes.Buffer
		for i := range c {
			buffer.WriteString(c[i].Error() + ";")
		}
		code = ValidateErrorCode(buffer.String())
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
