package gin_prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
	"time"
)

const (
	DefaultNameSpace       = "ficool"
	DefaultSubSystem       = "gin"
	DefaultMetricsPath     = "/metrics"
	DefaultSlowRequestTime = 5 * time.Second
)

var (
	ConstLabels = make(map[string]string, 0)
)

var (
	metricRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   DefaultNameSpace,
			Subsystem:   DefaultSubSystem,
			Name:        "request_total",
			Help:        "all the server received requests",
			ConstLabels: ConstLabels,
		},
		[]string{"path", "method", "code", "ip"})
	metricRequestBody = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   DefaultNameSpace,
			Subsystem:   DefaultSubSystem,
			Name:        "request_body_total",
			Help:        "the server received request body size, unit byte",
			ConstLabels: ConstLabels,
		},
		[]string{"path", "method", "code", "ip"})
	metricResponseBody = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   DefaultNameSpace,
			Subsystem:   DefaultSubSystem,
			Name:        "response_body_total",
			Help:        "the server received response body size, unit byte",
			ConstLabels: ConstLabels,
		},
		[]string{"path", "method", "code", "ip"})
	metricRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   DefaultNameSpace,
			Subsystem:   DefaultSubSystem,
			Name:        "request_duration",
			Help:        "the time server took to handle the request",
			ConstLabels: ConstLabels,
			Buckets:     []float64{0.001, 0.005, 0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09, 0.1, 0.2, 0.3, 0.4, 0.5, 1, 2, 3, 5, 10, 30, 60},
		},
		[]string{"path", "method", "code", "ip"})

	metricSlowRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   DefaultNameSpace,
			Subsystem:   DefaultSubSystem,
			Name:        "slow_request_total",
			Help:        "the server handled slow requests counter",
			ConstLabels: ConstLabels,
		},
		[]string{"path", "method", "code", "ip"})
)

func init() {
	prometheus.MustRegister(metricRequestTotal, metricRequestBody, metricResponseBody, metricRequestDuration, metricSlowRequest)
}

func Collector() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == DefaultMetricsPath {
			ctx.Next()
			return
		}
		startTime := time.Now()
		ctx.Next()
		metricsHandler(ctx, startTime)
		return
	}
}

func Register(r *gin.Engine) {
	r.GET(DefaultMetricsPath, func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})
}

func metricsHandler(ctx *gin.Context, start time.Time) {
	// set request total
	metricRequestTotal.
		WithLabelValues([]string{ctx.FullPath(), ctx.Request.Method, strconv.Itoa(ctx.Writer.Status()), ctx.ClientIP()}...).
		Inc()

	// set request body size
	//since r.ContentLength can be negative (in some occasions) guard the operation
	if ctx.Request.ContentLength >= 0 {
		metricRequestBody.
			WithLabelValues([]string{ctx.FullPath(), ctx.Request.Method, strconv.Itoa(ctx.Writer.Status()), ctx.ClientIP()}...).
			Add(float64(ctx.Request.ContentLength))
	}

	// set slow request
	latency := time.Since(start)
	if latency > DefaultSlowRequestTime {
		metricSlowRequest.
			WithLabelValues([]string{ctx.FullPath(), ctx.Request.Method, strconv.Itoa(ctx.Writer.Status()), ctx.ClientIP()}...).
			Inc()
	}

	// set request duration
	metricRequestDuration.
		WithLabelValues([]string{ctx.FullPath(), ctx.Request.Method, strconv.Itoa(ctx.Writer.Status()), ctx.ClientIP()}...).
		Observe(latency.Seconds())

	// set response size
	if ctx.Writer.Size() > 0 {
		metricResponseBody.
			WithLabelValues([]string{ctx.FullPath(), ctx.Request.Method, strconv.Itoa(ctx.Writer.Status()), ctx.ClientIP()}...).
			Add(float64(ctx.Writer.Size()))
	}
}
