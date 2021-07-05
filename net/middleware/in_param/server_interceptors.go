package in_param

import (
	"context"
	"encoding/json"
	"github.com/azdbaaaaaa/util/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
		inParamStr := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if md[ContextKeyInParam] != nil && len(md[ContextKeyInParam]) > 0 {
				inParamStr = md[ContextKeyInParam][0]
			}
		}
		if inParamStr != "" {
			inParam := &proto.InParam{}
			err := json.Unmarshal([]byte(inParamStr), inParam)
			if err != nil {
				logger.Error("in_param unmarshal", zap.String("key", ContextKeyInParam), SystemField, ServerField)
			} else {
				ctx = context.WithValue(ctx, ContextKeyInParam, inParam)
			}
		}
		resp, err := handler(ctx, req)
		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that adds zap.Logger to the context.
func StreamServerInterceptor(logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := stream.Context()

		inParamStr := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if md[ContextKeyInParam] != nil && len(md[ContextKeyInParam]) > 0 {
				inParamStr = md[ContextKeyInParam][0]
			}
		}
		if inParamStr != "" {
			inParam := &proto.InParam{}
			err := json.Unmarshal([]byte(inParamStr), inParam)
			if err != nil {
				logger.Error("in_param unmarshal", zap.String("key", ContextKeyInParam), SystemField, ServerField)
			} else {
				ctx = context.WithValue(ctx, ContextKeyInParam, inParam)
			}
		}
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = ctx
		err := handler(srv, wrapped)
		return err
	}
}