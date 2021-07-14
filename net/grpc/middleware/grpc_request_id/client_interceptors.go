package grpc_request_id

import (
	"context"
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	// ClientField is used in every client-side log statement made through grpc_zap. Can be overwritten before initialization.
	ClientField = zap.String("span.kind", "client")
)

// UnaryClientInterceptor returns a new unary client interceptor that optionally logs the execution of external gRPC calls.
func UnaryClientInterceptor(logger *zap.Logger) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		reqID := ""
		v := ctx.Value(metadata2.ContextKeyReqID)
		if v == nil {
			reqID = uuid.NewV4().String()
			logger.Debug("not found req_id generate one", ClientField, zap.String("uuid", reqID))
			ctx = context.WithValue(ctx, metadata2.ContextKeyReqID, reqID)
		} else {
			reqID = v.(string)
		}
		md := metadata.New(map[string]string{})
		md.Set(metadata2.ContextKeyReqID, reqID)
		ctx = metadata.NewOutgoingContext(ctx, md)
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

// StreamClientInterceptor returns a new streaming client interceptor that optionally logs the execution of external gRPC calls.
func StreamClientInterceptor(logger *zap.Logger) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		reqID := ""
		v := ctx.Value(metadata2.ContextKeyReqID)
		if v == nil {
			reqID = uuid.NewV4().String()
			logger.Debug("not found req_id generate one", ClientField, zap.String("uuid", reqID))
			ctx = context.WithValue(ctx, metadata2.ContextKeyReqID, reqID)
		} else {
			reqID = v.(string)
		}
		md := metadata.New(map[string]string{})
		md.Set(metadata2.ContextKeyReqID, reqID)
		ctx = metadata.NewOutgoingContext(ctx, md)
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		//newCtx := ctxzap.ToContext(ctx, logger.With(fields...))
		//logFinalClientLine(newCtx, o, startTime, err, "finished client streaming call")
		return clientStream, err
	}
}
