package net

import (
	"context"
	"fmt"
	net_grpc "github.com/azdbaaaaaa/util/net/grpc"
	net_http "github.com/azdbaaaaaa/util/net/http"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/oklog/run"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.uber.org/zap"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewServer(httpConf net_http.ServerConfig, grpcConf net_grpc.ServerConfig, log *zap.Logger, register func()) () {
	var g run.Group
	{
		r := gin.Default()
		p := ginprometheus.NewPrometheus("gin")
		p.Use(r)
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
				httpConf.ShutdownTimeout = 10 // default 10s
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
			log.Error("Listen Grpc Error", zap.String("addr", grpcConf.Addr))
			os.Exit(1)
		}
		g.Add(func() error {
			log.Info("transport HTTP ", zap.String("addr", grpcConf.Addr))
			s := net_grpc.NewServer(grpcConf, log)
			register()
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
