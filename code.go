package util

import (
	"strconv"
	"sync"
)

var _codes = &sync.Map{} // 注册Code信息

var (
	CodeSuccess    = NewCode(8000, "成功")
	CodeErrorParam = NewCode(8010, "参数不正确")
	CodeErrorToken = NewCode(9001, "令牌已过期，请重新获取令牌")
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
