package grpc

import (
	"context"
	"fmt"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_device"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_error"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_in_param"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_log"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_request_id"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_tracing"
	net_http "github.com/azdbaaaaaa/util/net/http"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/oklog/run"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const DefaultJaegerHost = "zipkin.istio-system"

func NewServer(conf ServerConfig, logger *zap.Logger) (s *grpc.Server) {
	s = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor([]grpc_recovery.Option{
				grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
					return status.Errorf(codes.Unknown, "panic triggered: %v", p)
				}),
			}...),
			grpc_tracing.StreamServerInterceptor(logger),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			grpc_log.StreamServerInterceptor(logger),
			grpc_validator.StreamServerInterceptor(),
			grpc_request_id.StreamServerInterceptor(logger),
			grpc_device.StreamServerInterceptor(logger),
			grpc_in_param.StreamServerInterceptor(logger),
			grpc_error.StreamServerInterceptor(logger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor([]grpc_recovery.Option{
				grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
					return status.Errorf(codes.Unknown, "panic triggered: %v", p)
				}),
			}...),
			grpc_tracing.UnaryServerInterceptor(logger),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_log.UnaryServerInterceptor(logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_request_id.UnaryServerInterceptor(logger),
			grpc_device.UnaryServerInterceptor(logger),
			grpc_in_param.UnaryServerInterceptor(logger),
			grpc_error.UnaryServerInterceptor(logger),
		)),
	)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s)
	return s
}

func RunServer(httpConf net_http.ServerConfig, grpcConf ServerConfig, log *zap.Logger, serviceName string, registerHttp func(r *gin.Engine), registerGrpc func(s *grpc.Server)) {
	if serviceName != "" {
		_, _, err := grpc_tracing.InitJaeger(serviceName, DefaultJaegerHost)
		if err != nil {
			log.Panic("init jaeger error")
			return
		}
	}

	var g run.Group
	{
		r := gin.Default()
		p := ginprometheus.NewPrometheus("gin")
		p.Use(r)
		registerHttp(r)
		if httpConf.Pprof {
			log.Info("http pprof on")
			pprof.Register(r)
		}
		server := &http.Server{
			Addr:    httpConf.Addr,
			Handler: r,
		}
		g.Add(func() error {
			log.Info("transport HTTP ", zap.String("addr", httpConf.Addr))
			return server.ListenAndServe()
		}, func(error) {
			log.Info("http shutdown running")
			if httpConf.ShutdownTimeout == 0 {
				httpConf.ShutdownTimeout = net_http.DefaultClientDialTimeoutInSec
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(httpConf.ShutdownTimeout)*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				log.Panic("http listen panic", zap.Error(err))
			}
			log.Info("http shutdown done")
		})
	}
	{
		grpcListener, err := net.Listen("tcp", grpcConf.Addr)
		if err != nil {
			log.Error("Listen GRPC Error", zap.String("addr", grpcConf.Addr))
			os.Exit(1)
		}
		g.Add(func() error {
			log.Info("transport GRPC ", zap.String("addr", grpcConf.Addr))
			s := NewServer(grpcConf, log)
			registerGrpc(s)
			return s.Serve(grpcListener)
		}, func(error) {
			log.Info("gRPC Listener closing")
			err = grpcListener.Close()
			log.Info("gRPC Listener closed", zap.Error(err))
		})
	}
	{
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		g.Add(func() error {
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			}
		}, func(error) {
			c <- syscall.SIGINT
		})
	}
	log.Info("exit:", zap.Error(g.Run()))
}
