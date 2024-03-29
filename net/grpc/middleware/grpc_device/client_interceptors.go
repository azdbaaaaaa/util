package grpc_device

import (
	"context"
	"encoding/json"
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
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
		data := make([]byte, 0)
		if v := ctx.Value(metadata2.ContextKeyDevice); v != nil {
			if d, ok := v.(metadata2.Device); ok {
				var err error
				data, err = json.Marshal(d)
				if err != nil {
					logger.Error("device marshal error", ClientField, zap.String("key", metadata2.ContextKeyDevice))
				}
			}
		}
		if len(data) != 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, metadata2.ContextKeyDevice, string(data))
		}

		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

// StreamClientInterceptor returns a new streaming client interceptor that optionally logs the execution of external gRPC calls.
func StreamClientInterceptor(logger *zap.Logger) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		//if v := ctx.Value(metadata2.ContextKeyDevice); v != nil {
		//	if d, ok := v.(metadata2.Device); ok {
		//		data, err := json.Marshal(d)
		//		if err != nil {
		//			logger.Error("device marshal error", ClientField, zap.String("key", metadata2.ContextKeyDevice))
		//		} else {
		//			md := metadata.New(map[string]string{})
		//			md.Set(metadata2.ContextKeyDevice, string(data))
		//			ctx = metadata.NewOutgoingContext(ctx, md)
		//		}
		//	}
		//}
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		//newCtx := ctxzap.ToContext(ctx, logger.With(fields...))
		//logFinalClientLine(newCtx, o, startTime, err, "finished client streaming call")
		return clientStream, err
	}
}
