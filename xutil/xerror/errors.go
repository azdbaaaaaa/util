package xerror

import (
	"github.com/azdbaaaaaa/util/proto/common"
)

var (
	Success         = New(int32(common.ErrCode_success), int32(common.ErrSubCode_err_sub_code_unknown), "success")
	ErrUnknown      = New(int32(common.ErrCode_unknown_error), int32(common.ErrSubCode_err_sub_code_unknown), common.ErrCode_unknown_error.String())
	ErrParamInvalid = New(int32(common.ErrCode_server_param_invalid), int32(common.ErrSubCode_err_sub_code_unknown), common.ErrCode_server_param_invalid.String())

	ErrNoDeviceError = New(int32(common.ErrCode_no_device_error), int32(common.ErrSubCode_err_sub_code_unknown), common.ErrCode_no_device_error.String())
	ErrNoInParamError = New(int32(common.ErrCode_no_in_param_error), int32(common.ErrSubCode_err_sub_code_unknown), common.ErrCode_no_in_param_error.String())
)
