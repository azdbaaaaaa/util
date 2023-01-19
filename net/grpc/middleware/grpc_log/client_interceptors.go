package grpc_log

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"path"
)

var (
	// ClientField is used in every client-side log statement made through grpc_zap. Can be overwritten before initialization.
	ClientField = zap.String("span.kind", "client")
)

// UnaryClientInterceptor returns a new unary client interceptor that optionally logs the execution of external gRPC calls.
func UnaryClientInterceptor(logger *zap.Logger) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		logger.Debug("logging client interceptor", ClientField,
			zap.String("methodName", path.Base(method)),
			zap.Any("req", req),
			zap.Any("reply", reply))
		return err
	}
}

// StreamClientInterceptor returns a new streaming client interceptor that optionally logs the execution of external gRPC calls.
func StreamClientInterceptor(logger *zap.Logger) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		logger.Debug("logging client interceptor", ClientField,
			zap.String("methodName", path.Base(method)))
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		return clientStream, err
	}
}
