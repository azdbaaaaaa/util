package grpc

import (
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_device"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_in_param"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_request_id"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewServer(conf ServerConfig, logger *zap.Logger) (s *grpc.Server) {
	s = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			grpc_validator.StreamServerInterceptor(),
			grpc_request_id.StreamServerInterceptor(logger),
			grpc_device.StreamServerInterceptor(logger),
			grpc_in_param.StreamServerInterceptor(logger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_request_id.UnaryServerInterceptor(logger),
			grpc_device.UnaryServerInterceptor(logger),
			grpc_in_param.UnaryServerInterceptor(logger),
		)),
	)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s)
	return s
}
