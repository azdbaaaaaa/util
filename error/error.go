package error

import (
	"sync"
)

var _codes = &sync.Map{} // 注册Code信息

var (
	CodeSuccess = NewCode(0, "成功")
	CodeUnknown = NewCode(999999, "未知错误")
)

func ValidateErrorCode() Code {
	return NewCode(100000, "参数错误")
}

func RpcErrorCode() Code {
	return NewCode(110000, "转发到服务错误")
}

type Code int

func NewCode(code Code, msg string) Code {
	_codes.Store(code, msg)
	return code
}

// Code return error code
func (c Code) Code() int {
	return int(c)
}

func (c Code) Error(kvs ...interface{}) string {
	v, ok := _codes.Load(c)
	if ok {
		return v.(string)
	}
	return c.Code()
}

func (c Code) IsSuccess() bool {
	return c == CodeSuccess
}
