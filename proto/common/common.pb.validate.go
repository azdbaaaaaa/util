// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common.proto

package common

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on OutParam with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OutParam) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for SubCode

	// no validation rules for Message

	// no validation rules for Reason

	// no validation rules for Metadata

	return nil
}

// OutParamValidationError is the validation error returned by
// OutParam.Validate if the designated constraints aren't met.
type OutParamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutParamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutParamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutParamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutParamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutParamValidationError) ErrorName() string { return "OutParamValidationError" }

// Error satisfies the builtin error interface
func (e OutParamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutParam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutParamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutParamValidationError{}

// Validate checks the field values on CommonResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CommonResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for SubCode

	// no validation rules for Message

	// no validation rules for Reason

	// no validation rules for Metadata

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CommonResponseValidationError is the validation error returned by
// CommonResponse.Validate if the designated constraints aren't met.
type CommonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonResponseValidationError) ErrorName() string { return "CommonResponseValidationError" }

// Error satisfies the builtin error interface
func (e CommonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonResponseValidationError{}
