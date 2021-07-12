package grpc

import (
	"context"
	"github.com/azdbaaaaaa/util/net/middleware/device"
	"github.com/azdbaaaaaa/util/net/middleware/in_param"
	"github.com/azdbaaaaaa/util/net/middleware/request_id"
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
		grpc.WithUnaryInterceptor(request_id.UnaryClientInterceptor(logger)),
		grpc.WithStreamInterceptor(request_id.StreamClientInterceptor(logger)),
		grpc.WithUnaryInterceptor(device.UnaryClientInterceptor(logger)),
		grpc.WithStreamInterceptor(device.StreamClientInterceptor(logger)),
		grpc.WithUnaryInterceptor(in_param.UnaryClientInterceptor(logger)),
		grpc.WithStreamInterceptor(in_param.StreamClientInterceptor(logger)),
	)
	if err != nil {
		logger.Error("did not connect", zap.Error(err))
	}
	return
}
