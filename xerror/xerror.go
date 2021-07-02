package xerror

import (
	"errors"
	"fmt"
)

type Error struct {
	Code     int32             `json:"code"`
	Message  string            `json:"message"`
	Reason   string            `json:"reason"`
	Metadata map[string]string `json:"metadata"`
}

// New returns an error object for the code, message.
func New(code int32, message, reason string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Reason:  reason,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s metadata = %v", e.Code, e.Reason, e.Message, e.Metadata)
}

func (e *Error) GetCode() int32 {
	return e.Code
}

// Is matches each error in the chain with the target value.
func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Code == e.Code
	}
	return false
}

//func (ec ErrorCode) IsSuccess() bool {
//	return ec == ErrCodeSuccess
//}
//
//func (ec ErrorCode) Equals(code int32) bool {
//	return ec.Code == code
//}
//

func (e *Error) WithReason(reason string) *Error {
	e.Reason = reason
	return e
}
