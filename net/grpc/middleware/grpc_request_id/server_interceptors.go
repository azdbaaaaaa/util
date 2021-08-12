package grpc_request_id

import (
	"context"
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	uuid "github.com/satori/go.uuid"
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
		reqID := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if md[metadata2.ContextKeyReqID] != nil && len(md[metadata2.ContextKeyReqID]) > 0 {
				reqID = md[metadata2.ContextKeyReqID][0]
			}
		}
		if reqID == "" {
			reqID = uuid.NewV4().String()
			//logger.Debug("not found req_id generate one", SystemField, ServerField, zap.String("uuid", reqID))
		}
		ctx = context.WithValue(ctx, metadata2.ContextKeyReqID, reqID)
		resp, err := handler(ctx, req)
		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that adds zap.Logger to the context.
func StreamServerInterceptor(logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := stream.Context()
		reqID := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if md[metadata2.ContextKeyReqID] != nil && len(md[metadata2.ContextKeyReqID]) > 0 {
				reqID = md[metadata2.ContextKeyReqID][0]
			}
		}
		if reqID == "" {
			reqID = uuid.NewV4().String()
			//logger.Debug("not found req_id generate one", SystemField, ServerField, zap.String("uuid", reqID))
		}
		ctx = context.WithValue(ctx, metadata2.ContextKeyReqID, reqID)

		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = ctx
		err := handler(srv, wrapped)
		return err
	}
}
