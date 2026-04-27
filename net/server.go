package net

import (
	"context"
	"fmt"
	net_grpc "github.com/azdbaaaaaa/util/net/grpc"
	"github.com/azdbaaaaaa/util/net/grpc/middleware/grpc_tracing"
	net_http "github.com/azdbaaaaaa/util/net/http"
	"github.com/azdbaaaaaa/util/net/http/gin/middleware/health"
	gin_prometheus "github.com/azdbaaaaaa/util/net/http/gin/middleware/prometheus"
	"github.com/azdbaaaaaa/util/net/http/gin/middleware/request_id"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/oklog/run"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const DefaultJaegerHost = "zipkin.istio-system"

func RunServer(httpConf net_http.ServerConfig, grpcConf net_grpc.ServerConfig, log *zap.Logger, serviceName string, registerHttp func(r *gin.Engine), registerGrpc func(s *grpc.Server)) {
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
		health.Register(r)
		gin_prometheus.Register(r)
		registerHttp(r)
		r.Use(
			gin_prometheus.Collector(),
			request_id.RequestID,
			gzip.Gzip(gzip.DefaultCompression),
		)
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
			s := net_grpc.NewServer(grpcConf, log)
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
