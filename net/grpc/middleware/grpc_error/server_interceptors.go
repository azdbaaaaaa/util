package error

import (
	"context"
	"github.com/azdbaaaaaa/util/xutil/xerror"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var (
	// SystemField is used in every log statement made through grpc_zap. Can be overwritten before any initialization code.
	SystemField = zap.String("system", "grpc")

	// ServerField is used in every server-side log statement made through grpc_zap.Can be overwritten before initialization.
	ServerField = zap.String("span.kind", "server")
)

// UnaryServerInterceptor returns a new unary server interceptors that adds zap.Logger to the context.
func UnaryServerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, status.Errorf(codes.NotFound, err.Error())
			}
			switch err.(type) {
			case xerror.Error:
				err = nil
			default:
				logger.Error("unknown error", zap.Errors("err", []error{err}), SystemField, ServerField)
				return nil, status.Errorf(codes.Internal, err.Error())
			}
		}
		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that adds zap.Logger to the context.
func StreamServerInterceptor(logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := stream.Context()
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = ctx
		err := handler(srv, wrapped)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return status.Errorf(codes.NotFound, err.Error())
			}
			switch err.(type) {
			case xerror.Error:
				err = nil
			default:
				logger.Error("unknown error", zap.Errors("err", []error{err}), SystemField, ServerField)
				return status.Errorf(codes.Internal, err.Error())
			}
		}
		return err
	}
}
