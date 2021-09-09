package grpc_error

import (
	"github.com/azdbaaaaaa/util/proto/common"
	"github.com/azdbaaaaaa/util/xutil/xerror"
	"github.com/qiniu/qmgo"
	"gorm.io/gorm"
)

func ErrorWrapper(param *common.OutParam, err error) {
	if err != nil {
		if err == gorm.ErrRecordNotFound || err == qmgo.ErrNoSuchDocuments {
			param.Code = xerror.ErrNotFound.GetCode()
			param.SubCode = xerror.ErrNotFound.GetSubCode()
			param.Message = xerror.ErrNotFound.GetMessage()
			param.Reason = xerror.ErrNotFound.GetReason()
		}
		if v, ok := err.(xerror.Error); ok {
			param.Code = v.GetCode()
			param.SubCode = v.GetSubCode()
			param.Message = v.GetMessage()
			param.Reason = v.GetReason()
		}
	}
}
