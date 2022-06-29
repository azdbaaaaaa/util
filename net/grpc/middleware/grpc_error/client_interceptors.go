package grpc_error

import (
	"context"
	"github.com/azdbaaaaaa/util/proto/common"
	"github.com/azdbaaaaaa/util/xutil/xerror"
	"github.com/qiniu/qmgo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	// ClientField is used in every client-side log statement made through grpc_zap. Can be overwritten before initialization.
	ClientField = zap.String("span.kind", "client")
)

func UnaryClientInterceptor(logger *zap.Logger) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if res, ok := reply.(*common.CommonResponse); ok {
			if res.Code != int32(common.ErrCode_ERR_CODE_UNSPECIFIED) {
				logger.Error("CommonResponse error",
					ClientField,
					zap.Int32("code", res.Code),
					zap.Int32("subCode", res.SubCode),
					zap.String("message", res.Message),
					zap.String("reason", res.Reason),
				)
				return xerror.New(res.Code, res.SubCode, res.Message).WithReason(res.Reason)
			}
		}
		switch err {
		case qmgo.ErrNoSuchDocuments, gorm.ErrRecordNotFound:
			return xerror.ErrNotFound
		default:
		}
		return err
	}
}

func StreamClientInterceptor(logger *zap.Logger) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		return clientStream, err
	}
}
