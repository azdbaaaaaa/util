package grpc

import (
	"context"
	device2 "github.com/azdbaaaaaa/util/net/grpc/middleware/device"
	in_param2 "github.com/azdbaaaaaa/util/net/grpc/middleware/in_param"
	request_id2 "github.com/azdbaaaaaa/util/net/grpc/middleware/request_id"
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
		grpc.WithUnaryInterceptor(request_id2.UnaryClientInterceptor(logger)),
		grpc.WithStreamInterceptor(request_id2.StreamClientInterceptor(logger)),
		grpc.WithUnaryInterceptor(device2.UnaryClientInterceptor(logger)),
		grpc.WithStreamInterceptor(device2.StreamClientInterceptor(logger)),
		grpc.WithUnaryInterceptor(in_param2.UnaryClientInterceptor(logger)),
		grpc.WithStreamInterceptor(in_param2.StreamClientInterceptor(logger)),
	)
	if err != nil {
		logger.Error("did not connect", zap.Error(err))
	}
	return
}
