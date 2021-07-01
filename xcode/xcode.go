package xcode

import (
	"fmt"
	"sync"
)

var _codes = &sync.Map{} // 注册Code信息

//func ValidateErrorCode() Code {
//	return NewCode(100000, "params error")
//}
//
//func RpcErrorCode() Code {
//	return NewCode(110000, "internal service error")
//}

type ErrorCode struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Reason string `json:"reason"`
}

func NewCodeError(code int, msg string) ErrorCode {
	//_codes.Store(code, msg)
	ec := ErrorCode{Code: code, Msg: msg}
	_, loaded := _codes.LoadOrStore(code, ec)
	if loaded {
		panic(fmt.Sprintf("code already exists:%d", code))
	}
	return ec
}

// GetCode return error code
func (ec ErrorCode) GetCode() int {
	return ec.Code
}

func (ec ErrorCode) Error() string {
	v, ok := _codes.Load(ec)
	if ok {
		return v.(ErrorCode).Msg
	}
	return ""
}

func (ec ErrorCode) IsSuccess() bool {
	return ec == ErrCodeSuccess
}

func (ec ErrorCode) Equals(ce2 ErrorCode) bool {
	return ec.Code == ce2.Code
}

func (ec ErrorCode) WithReason(reason string) {
	ec.Reason = reason
	return
}
