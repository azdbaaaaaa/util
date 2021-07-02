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
	Code   int32  `json:"code"`
	Msg    string `json:"msg"`
	Reason string `json:"reason"`
}

func NewCodeError(code int32, msg string) ErrorCode {
	//_codes.Store(code, msg)
	ec := ErrorCode{Code: code, Msg: msg}
	_, loaded := _codes.LoadOrStore(code, ec)
	if loaded {
		panic(fmt.Sprintf("code already exists:%d", code))
	}
	return ec
}

// GetCode return error code
func (ec ErrorCode) GetCode() int32 {
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

func (ec ErrorCode) Equals(code int32) bool {
	return ec.Code == code
}

func (ec ErrorCode) WithReason(reason string) ErrorCode {
	ec.Reason = reason
	return ec
}
