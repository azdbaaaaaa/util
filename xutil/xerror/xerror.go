package xerror

import (
	"fmt"
)

type Error interface {
	Error() string
	GetCode() int32
	GetSubCode() int32
	GetMessage() string
	GetReason() string
	WithReason(reason string) Error
}

type err struct {
	Code     int32             `json:"code"`
	SubCode  int32             `json:"sub_code"`
	Message  string            `json:"message"`
	Reason   string            `json:"reason"`
	Metadata map[string]string `json:"metadata"`
}

// New returns an error object for the code, message.
func New(code, subCode int32, message string) Error {
	return &err{
		Code:    code,
		SubCode: subCode,
		Message: message,
	}
}

func NewError(code int32, message string) Error {
	return &err{
		Code:    code,
		Message: message,
	}
}

type ErrorCodeMessage interface {
	String() string
}

func NewProtoError(code interface{}) Error {
	e := &err{Code: code.(int32)}
	e.Message = code.(ErrorCodeMessage).String()
	return e
}

func (e *err) WithSubCode(subCode int32) Error {
	e.SubCode = subCode
	return e
}

func (e *err) WithMessage(message string) Error {
	e.Message = message
	return e
}

func (e *err) WithReason(reason string) Error {
	e.Reason = reason
	return e
}

func (e *err) Error() string {
	return fmt.Sprintf("error: code = %d subCode = %d reason = %s message = %s metadata = %v", e.Code, e.SubCode, e.Reason, e.Message, e.Metadata)
}

func (e *err) GetCode() int32 {
	return e.Code
}

func (e *err) GetSubCode() int32 {
	return e.SubCode
}

// Is matches each error in the chain with the target value.
//func (e *err) Is(err error) bool {
//	if se := new(err); errors.As(err, &se) {
//		return se.Code == e.Code && se.SubCode == e.SubCode
//	}
//	return false
//}

func (e *err) GetMessage() string {
	return e.Message
}

func (e *err) GetReason() string {
	return e.Reason
}
