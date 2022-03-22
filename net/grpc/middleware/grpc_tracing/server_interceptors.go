package grpc_tracing

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/opentracing/opentracing-go"
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
		// 1. 从ctx中解析出md
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			//如果对metadata进行修改，那么需要用拷贝的副本进行修改。（FromIncomingContext的注释）
			md = md.Copy()
		}

		// 2. 从md创建的carrier中解析出span
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
		}

		// 3. 创建一个新的span
		span := tracer.StartSpan(info.FullMethod, opentracing.ChildOf(spanContext))
		defer span.Finish()
		span.SetTag("span.kind", "server")
		span.SetTag("grpc.service", service)
		span.SetTag("grpc.method", methodName)
		span.SetTag("error", false)

		// 4. 把span放到ctx中
		ctx = opentracing.ContextWithSpan(ctx, span)
		resp, err := handler(ctx, req)
		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that adds zap.Logger to the context.
func StreamServerInterceptor(logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := stream.Context()
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			//如果对metadata进行修改，那么需要用拷贝的副本进行修改。（FromIncomingContext的注释）
			md = md.Copy()
		}
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = ctx
		err := handler(srv, wrapped)
		return err
	}
}
