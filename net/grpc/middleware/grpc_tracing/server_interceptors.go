package grpc_tracing

import (
	"context"
	"fmt"
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/opentracing/opentracing-go"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"path"
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
		service := path.Dir(info.FullMethod)[1:]
		methodName := path.Base(info.FullMethod)
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			//如果对metadata进行修改，那么需要用拷贝的副本进行修改。（FromIncomingContext的注释）
			md = md.Copy()
		}
		carrier := TextMapReader{md}
		_ = carrier.ForeachKey(func(key, val string) error {
			logger.Debug("incoming md", zap.String(key, val))
			return nil
		})
		tracer := opentracing.GlobalTracer()
		spanContext, e := tracer.Extract(opentracing.TextMap, carrier)
		if e != nil {
			logger.Debug("extract error",
				zap.String("grpc.service", service),
				zap.String("grpc.method", methodName),
				ServerField,
				SystemField,
				zap.Error(e),
			)
			fmt.Println("Extract err:", e)
		}

		span := tracer.StartSpan(info.FullMethod, opentracing.ChildOf(spanContext))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)

		//ctx = context.WithValue(ctx, metadata2.ContextKeyReqID, reqID)
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
