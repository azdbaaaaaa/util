package grpc

import (
	"context"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_device"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_in_param"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_request_id"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

const (
	DefaultClientDialTimeoutInSec = 10
)

type ServerConfig struct {
	Addr string `json:"addr"`
}

type ClientConfig struct {
	Addr        string `json:"addr"`
	DialTimeout int    `json:"dial_timeout" mapstructure:"dial_timeout"`
}

func NewClientConn(conf ClientConfig, logger *zap.Logger) (conn *grpc.ClientConn, err error) {
	if conf.DialTimeout == 0 {
		conf.DialTimeout = DefaultClientDialTimeoutInSec
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(conf.DialTimeout)*time.Second)
	defer cancel()
	conn, err = grpc.DialContext(ctx, conf.Addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(grpc_request_id.UnaryClientInterceptor(logger)),
		grpc.WithStreamInterceptor(grpc_request_id.StreamClientInterceptor(logger)),
		grpc.WithUnaryInterceptor(grpc_device.UnaryClientInterceptor(logger)),
		grpc.WithStreamInterceptor(grpc_device.StreamClientInterceptor(logger)),
		grpc.WithUnaryInterceptor(grpc_in_param.UnaryClientInterceptor(logger)),
		grpc.WithStreamInterceptor(grpc_in_param.StreamClientInterceptor(logger)),
	)
	if err != nil {
		logger.Error("did not connect", zap.Error(err))
	}
	return
}
