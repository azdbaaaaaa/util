package request_id

import (
	"context"
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
		v := ctx.Value(ContextKeyReqID)
		if v == nil {
			u := uuid.NewV4().String()
			logger.Debug("not found req_id generate one", zap.String("uuid", u))
			ctx = context.WithValue(ctx, ContextKeyReqID, u)
		}
		md := metadata.New(map[string]string{
			ContextKeyReqID: v.(string),
		})
		ctx = metadata.NewOutgoingContext(ctx, md)
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

// StreamClientInterceptor returns a new streaming client interceptor that optionally logs the execution of external gRPC calls.
func StreamClientInterceptor(logger *zap.Logger) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		v := ctx.Value(ContextKeyReqID)
		if v == nil {
			u := uuid.NewV4().String()
			logger.Debug("not found req_id generate one", zap.String("uuid", u))
			ctx = context.WithValue(ctx, ContextKeyReqID, u)
		}
		md := metadata.New(map[string]string{
			ContextKeyReqID: v.(string),
		})
		ctx = metadata.NewOutgoingContext(ctx, md)
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		//newCtx := ctxzap.ToContext(ctx, logger.With(fields...))
		//logFinalClientLine(newCtx, o, startTime, err, "finished client streaming call")
		return clientStream, err
	}
}
