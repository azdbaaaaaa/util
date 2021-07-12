package http

import (
	"github.com/azdbaaaaaa/util/net/http/gin/middleware/health"
	"github.com/azdbaaaaaa/util/net/http/gin/middleware/pprof"
	"github.com/azdbaaaaaa/util/net/middleware/device"
	"github.com/azdbaaaaaa/util/net/middleware/in_param"
	"github.com/azdbaaaaaa/util/net/middleware/request_id"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.uber.org/zap"
)

func NewServer(conf ServerConfig, logger *zap.Logger) {
	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	health.Register(r)
	pprof.Register(conf.Pprof, r)
	r.Use(request_id.RequestID, device.SetDevice, in_param.SetInParam)

	//api.WithHandler(r, svc)
	//server := &http.Server{
	//	Addr:    conf.Addr,
	//	Handler: r,
	//}
	//g.Add(func() error {
	//	logger.Info("transport HTTP addr:%s",zap.String("addr": conf.HTTP.Addr) )
	//	return server.ListenAndServe()
	//}, func(error) {
	//	logger.Infow("http shutdown running")
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.HTTP.ShutdownTimeout)*time.Second)
	//	defer cancel()
	//	if err := server.Shutdown(ctx); err != nil {
	//		logger.Panicf("listen: %s", err)
	//	}
	//	logger.Infow("http shutdown done")
	//})
}
