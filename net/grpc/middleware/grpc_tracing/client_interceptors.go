package grpc_tracing

import (
	"context"
	"fmt"
	metadata2 "github.com/azdbaaaaaa/util/net/metadata"
	"github.com/opentracing/opentracing-go"
	uuid "github.com/satori/go.uuid"
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

		var parentContext opentracing.SpanContext
		//先从context中获取原始的span
		parentSpan := opentracing.SpanFromContext(ctx)
		if parentSpan != nil {
			parentContext = parentSpan.Context()
		}
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
			fmt.Println("tracer Inject err,", e)
		}
		//创建一个新的context，把metadata附带上
		ctx = metadata.NewOutgoingContext(ctx, md)

		//
		//var span opentracing.Span
		//opName := service + ":" + methodName
		//wireContext, err := opentracing.GlobalTracer().Extract(
		//	opentracing.TextMap,
		//	opentracing.TextMapCarrier(),
		//	//opentracing.HTTPHeadersCarrier(c.Request().Header),
		//)
		//if err != nil {
		//	span = opentracing.StartSpan(opName)
		//} else {
		//	log.Debug("opentracing span child")
		//	span = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
		//}
		//
		//defer span.Finish()

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
		reqID := ""
		v := ctx.Value(metadata2.ContextKeyReqID)
		if v == nil {
			reqID = uuid.NewV4().String()
			//logger.Debug("not found req_id generate one", ClientField, zap.String("uuid", reqID))
			ctx = context.WithValue(ctx, metadata2.ContextKeyReqID, reqID)
		} else {
			reqID = v.(string)
		}
		ctx = metadata.AppendToOutgoingContext(ctx, metadata2.ContextKeyReqID, reqID)
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		return clientStream, err
	}
}
