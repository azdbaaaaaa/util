package xerror

import "github.com/azdbaaaaaa/util/proto"

var (
	Success         = New(int32(proto.ErrCode_success), int32(proto.ErrCode_unknown_error), "success", "")
	ErrUnknown      = New(int32(proto.ErrCode_unknown_error), int32(proto.ErrCode_unknown_error), proto.ErrCode_unknown_error.String(), "")
	ErrParamInvalid = New(int32(proto.ErrCode_server_param_invalid), int32(proto.ErrCode_unknown_error), proto.ErrCode_server_param_invalid.String(), "")
)
