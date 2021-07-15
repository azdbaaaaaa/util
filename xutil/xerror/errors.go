package xerror

import (
	"github.com/azdbaaaaaa/util/proto/common"
)

var (
	Success         = New(int32(common.ErrCode_success), int32(common.ErrCode_unknown_error), "success", "")
	ErrUnknown      = New(int32(common.ErrCode_unknown_error), int32(common.ErrCode_unknown_error), common.ErrCode_unknown_error.String(), "")
	ErrParamInvalid = New(int32(common.ErrCode_server_param_invalid), int32(common.ErrCode_unknown_error), common.ErrCode_server_param_invalid.String(), "")
)
