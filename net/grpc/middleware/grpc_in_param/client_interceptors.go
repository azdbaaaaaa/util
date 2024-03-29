package grpc_in_param

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
		//logger.Debug("start in_param client interceptor", ClientField)
		data := make([]byte, 0)
		if v := ctx.Value(metadata2.ContextKeyInParam); v != nil {
			if inParam, ok := v.(metadata2.InParam); ok {
				var err error
				data, err = json.Marshal(inParam)
				if err != nil {
					logger.Error("in_param marshal error", ClientField, zap.String("key", metadata2.ContextKeyInParam))
				}
			}
		}
		if len(data) != 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, metadata2.ContextKeyInParam, string(data))
		}

		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

// StreamClientInterceptor returns a new streaming client interceptor that optionally logs the execution of external gRPC calls.
func StreamClientInterceptor(logger *zap.Logger) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		//if v := ctx.Value(metadata2.ContextKeyInParam); v != nil {
		//	if inParam, ok := v.(metadata2.InParam); ok {
		//		data, err := json.Marshal(inParam)
		//		if err != nil {
		//			logger.Error("in_param marshal error", ClientField, zap.String("key", metadata2.ContextKeyInParam))
		//		} else {
		//			md := metadata.New(map[string]string{})
		//			md.Set(metadata2.ContextKeyInParam, string(data))
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
