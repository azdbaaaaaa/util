package gin

import (
	"github.com/azdbaaaaaa/util/log"
	gin_prometheus "github.com/azdbaaaaaa/util/net/http/gin/middleware/prometheus"
	"github.com/azdbaaaaaa/util/net/metadata"
	"github.com/azdbaaaaaa/util/xutil/xerror"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

var (
	metricResultTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   gin_prometheus.DefaultNameSpace,
			Subsystem:   gin_prometheus.DefaultSubSystem,
			Name:        "result_request_total",
			Help:        "all the server response result count",
			ConstLabels: gin_prometheus.ConstLabels,
		},
		[]string{"path", "method", "code", "ip", "result", "version"})
)

func init() {
	prometheus.MustRegister(metricResultTotal)
}

type ApiError struct {
	// 错误码 ex:999999
	Result int32 `json:"result"`
	// 错误消息
	Message string `json:"message"`
	// 数据
	Data interface{} `json:"data"`

	// error debug信息
	Reason string `json:"reason,omitempty"`
	// request id
	Rid string `json:"rid"`
	// ts now
	Ts int64 `json:"ts"`
}

type WrapperHandle func(c *gin.Context) (interface{}, error)

func ErrorWrapper(handle WrapperHandle) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := xerror.Success
		var (
			rid, version string
		)
		start := time.Now()
		data, err := handle(ctx)
		if err != nil {
			if ec, ok := err.(xerror.Error); ok {
				code = ec
			} else if err == gorm.ErrRecordNotFound {
				code = xerror.ErrNotFound.WithReason(err.Error())
			} else if err == mongo.ErrNoDocuments {
				code = xerror.ErrNotFound.WithReason(err.Error())
			} else {
				code = xerror.ErrUnknown.WithReason(err.Error())
			}
		}
		if reqId, exists := ctx.Get(metadata.ContextKeyReqID); exists {
			rid = reqId.(string)
		}
		if device, exists := ctx.Get(metadata.ContextKeyDevice); exists {
			version = device.(metadata.Device).VersionName
		}
		metricResultTotal.
			WithLabelValues([]string{
				ctx.FullPath(), ctx.Request.Method, strconv.Itoa(ctx.Writer.Status()),
				ctx.ClientIP(), strconv.Itoa(int(code.GetCode())), version,
			}...).
			Inc()
		resp := ApiError{
			Result:  code.GetCode(),
			Message: code.GetMessage(),
			Data:    data,
			Reason:  code.GetReason(),
			Rid:     rid,
			Ts:      time.Now().UnixNano() / 1e6,
		}
		log.Infow(ctx.FullPath(), "result", resp.Result, "message", resp.Message, "reason", resp.Reason, "rid", rid, "time", time.Since(start).String())
		ctx.JSON(http.StatusOK, resp)
	}
}
