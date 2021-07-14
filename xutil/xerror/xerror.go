package xerror

import (
	"fmt"
)

type Error interface {
	Error() string
	GetCode() int32
	GetSubCode() int32
}

type err struct {
	Code     int32             `json:"code"`
	SubCode  int32             `json:"sub_code"`
	Message  string            `json:"message"`
	Reason   string            `json:"reason"`
	Metadata map[string]string `json:"metadata"`
}

// New returns an error object for the code, message.
func New(code, subCode int32, message, reason string) Error {
	return &err{
		Code:    code,
		SubCode: subCode,
		Message: message,
		Reason:  reason,
	}
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

//func (ec ErrorCode) IsSuccess() bool {
//	return ec == ErrCodeSuccess
//}
//
//func (ec ErrorCode) Equals(code int32) bool {
//	return ec.Code == code
//}
//

func (e *err) WithReason(reason string) Error {
	e.Reason = reason
	return e
}
