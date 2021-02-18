package util

import (
	"strconv"
	"sync"
)

var _codes = &sync.Map{} // 注册Code信息

var (
	CodeSuccess        = NewCode(8000, "成功")
	CodeUnknown        = NewCode(8001, "未知错误")
	CodeErrorParam     = NewCode(8010, "参数不正确")
	CodeTokenError     = NewCode(9000, "无效令牌，请重新获取令牌")
	CodeTokenExpired   = NewCode(9001, "令牌已过期，请重新获取令牌")
	CodeCtxKeyNotFound = NewCode(9100, "请求上下文中键不存在")
	CodeCtxKeyInvalid  = NewCode(9101, "请求上下文中键不合法")
)

type Code int

func NewCode(code Code, msg string) Code {
	_codes.Store(code, msg)
	return code
}

func (c Code) Error() string {
	return strconv.FormatInt(int64(c), 10)
}

// Code return error code
func (c Code) Code() int { return int(c) }

// Message return error message
func (c Code) Message() string {
	v, ok := _codes.Load(c)
	if ok {
		return v.(string)
	}
	return c.Error()
}

func (c Code) IsSuccess() bool {
	return c == CodeSuccess
}
