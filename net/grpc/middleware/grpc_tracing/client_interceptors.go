package grpc_tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"path"
)

var (
	// ClientField is used in every client-side log statement made through grpc_zap. Can be overwritten before initialization.
	ClientField = zap.String("span.kind", "client")
)

// UnaryClientInterceptor returns a new unary client interceptor that optionally logs the execution of external gRPC calls.
func UnaryClientInterceptor(logger *zap.Logger) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		//logger.Info("start req_id client interceptor", ClientField)
		service := path.Dir(method)[1:]
		methodName := path.Base(method)

		// 1. 先从context中获取原始的span
		var parentContext opentracing.SpanContext
		parentSpan := opentracing.SpanFromContext(ctx)
		if parentSpan != nil {
			parentContext = parentSpan.Context()
		}

		// 2. 创建一个新的span
		tracer := opentracing.GlobalTracer()
		span := tracer.StartSpan(method, opentracing.ChildOf(parentContext))
		defer span.Finish()

		//从context中获取metadata。md.(type) == map[string][]string
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			//如果对metadata进行修改，那么需要用拷贝的副本进行修改。（FromIncomingContext的注释）
			md = md.Copy()
		}
		//定义一个carrier，下面的Inject注入数据需要用到。carrier.(type) == map[string]string
		//carrier := opentracing.TextMapCarrier{}
		carrier := TextMapWriter{md}
		//将span的context信息注入到carrier中
		e := tracer.Inject(span.Context(), opentracing.TextMap, carrier)
		if e != nil {
			logger.Debug("inject error",
				zap.String("grpc.service", service),
				zap.String("grpc.method", methodName),
				ClientField,
				SystemField,
				zap.Error(e),
			)
		}
		//创建一个新的context，把metadata附带上
		ctx = metadata.NewOutgoingContext(ctx, carrier.MD)

		span.SetTag("span.kind", "server")
		span.SetTag("grpc.service", service)
		span.SetTag("grpc.method", methodName)
		span.SetTag("error", false)
		//span.SetTag("grpc.status_code", c.Response().Status)
		//ctx = metadata.AppendToOutgoingContext(ctx, metadata2.ContextKeyTracing, md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// StreamClientInterceptor returns a new streaming client interceptor that optionally logs the execution of external gRPC calls.
func StreamClientInterceptor(logger *zap.Logger) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		//md, ok := metadata.FromIncomingContext(ctx)
		//if !ok {
		//	md = metadata.New(nil)
		//} else {
		//	//如果对metadata进行修改，那么需要用拷贝的副本进行修改。（FromIncomingContext的注释）
		//	md = md.Copy()
		//}
		//ctx = metadata.NewOutgoingContext(ctx, md)
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		return clientStream, err
	}
}
