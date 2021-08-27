package grpc

import (
	"context"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_device"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_in_param"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_request_id"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func NewClientConn(conf ClientConfig, logger *zap.Logger) (conn *grpc.ClientConn, err error) {
	if conf.DialTimeout == 0 {
		conf.DialTimeout = DefaultClientDialTimeoutInSec
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(conf.DialTimeout)*time.Second)
	defer cancel()
	conn, err = grpc.DialContext(ctx, conf.Addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_prometheus.UnaryClientInterceptor,
			grpc_request_id.UnaryClientInterceptor(logger),
			grpc_device.UnaryClientInterceptor(logger),
			grpc_in_param.UnaryClientInterceptor(logger),
			grpc_zap.UnaryClientInterceptor(logger, []grpc_zap.Option{grpc_zap.WithDurationField(grpc_zap.DurationToDurationField)}...),
		)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_prometheus.StreamClientInterceptor,
			grpc_request_id.StreamClientInterceptor(logger),
			grpc_device.StreamClientInterceptor(logger),
			grpc_in_param.StreamClientInterceptor(logger),
			grpc_zap.StreamClientInterceptor(logger, []grpc_zap.Option{grpc_zap.WithDurationField(grpc_zap.DurationToDurationField)}...),
		)),
	)
	if err != nil {
		logger.Error("did not connect", zap.Error(err))
	}
	return
}
