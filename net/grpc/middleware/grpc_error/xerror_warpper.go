package grpc_error

import (
	"github.com/azdbaaaaaa/util/proto/common"
	"github.com/azdbaaaaaa/util/xutil/xerror"
)

func ErrorWrapper(param *common.OutParam, err error) {
	if err != nil {
		if v, ok := err.(xerror.Error); ok {
			param.Code = v.GetCode()
			param.SubCode = v.GetSubCode()
			param.Message = v.GetMessage()
			param.Reason = v.GetReason()
		}
	}
}
