package grpc

import (
	"context"
	"github.com/azdbaaaaaa/util/net/middleware/device"
	"github.com/azdbaaaaaa/util/net/middleware/in_param"
	"github.com/azdbaaaaaa/util/net/middleware/request_id"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"time"
)

const (
	DefailtClientDialTimeoutInSec = 10
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
		conf.DialTimeout = DefailtClientDialTimeoutInSec
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

func NewServer(conf ServerConfig, logger *zap.Logger) (s *grpc.Server) {
	s = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			grpc_validator.StreamServerInterceptor(),
			request_id.StreamServerInterceptor(logger),
			device.StreamServerInterceptor(logger),
			in_param.StreamServerInterceptor(logger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_validator.UnaryServerInterceptor(),
			request_id.UnaryServerInterceptor(logger),
			device.UnaryServerInterceptor(logger),
			in_param.UnaryServerInterceptor(logger),
		)),
	)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s)
	return s
}
