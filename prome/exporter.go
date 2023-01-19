package prome

import (
	"github.com/azdbaaaaaa/util/log"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"net/http"
)

type PrometheusConfig struct {
	Port int `json:"port"`
}

type PromeExporter struct {
	srv *http.Server
	up  prometheus.Gauge

	total       *prometheus.GaugeVec
	qps         *prometheus.GaugeVec
	avgCost     *prometheus.GaugeVec
	maxCost     *prometheus.GaugeVec
	costSummery *prometheus.SummaryVec
	counter     *prometheus.GaugeVec
	avgCounter  *prometheus.GaugeVec
}

var labels = []string{"status", "name"}

func NewExporter(ns string) *PromeExporter {
	return &PromeExporter{
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "up",
			Help:      "work success",
		}),
		total:      NewGaugeVec(ns, "total", "请求总数"),
		qps:        NewGaugeVec(ns, "qps", "QPS"),
		avgCost:    NewGaugeVec(ns, "avg_cost", "平均耗时"),
		maxCost:    NewGaugeVec(ns, "max_cost", "最大耗时"),
		counter:    NewGaugeVec(ns, "counter", "计数"),
		avgCounter: NewGaugeVec(ns, "avg_counter", "平均计数"),
	}
}

func NewGaugeVec(ns, name, help string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: ns,
		Name:      name,
		Help:      help,
	}, labels)
}

func (pe *PromeExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- pe.up.Desc()

	pe.total.Describe(ch)
	pe.qps.Describe(ch)
	pe.avgCost.Describe(ch)
	pe.maxCost.Describe(ch)
	pe.counter.Describe(ch)
	pe.avgCounter.Describe(ch)
}

func (pe *PromeExporter) reset() {
	pe.total.Reset()
	pe.qps.Reset()
	pe.avgCounter.Reset()
	pe.counter.Reset()
	pe.avgCost.Reset()
	pe.maxCost.Reset()
	//pe.costSummery.Reset()
}

func (pe *PromeExporter) Collect(ch chan<- prometheus.Metric) {
	defer func() {
		if err := recover(); err != nil {
			log.Infow("prometheus", zap.String("err", "collect"))
		}
	}()

	pe.reset()

	// Collect
	statInfos := promeStater.Collect()
	if len(statInfos) == 0 {
		pe.up.Set(0)
	} else {
		pe.up.Set(1)
		for _, info := range statInfos {
			labelDict := make(map[string]string, 8)
			labelDict["name"] = info.Name
			labelDict["status"] = info.Status
			labels := prometheus.Labels(labelDict)
			pe.total.With(labels).Add(float64(info.Total))
			pe.qps.With(labels).Add(float64(info.QPS))
			pe.maxCost.With(labels).Add(info.MaxCost)
			pe.counter.With(labels).Add(float64(info.Counter))
			pe.avgCost.With(labels).Add(info.AvgCost)
			pe.avgCounter.With(labels).Add(float64(info.AvgCounter))
		}
		pe.total.Collect(ch)
		pe.qps.Collect(ch)
		pe.maxCost.Collect(ch)
		pe.counter.Collect(ch)
		pe.avgCost.Collect(ch)
		pe.avgCounter.Collect(ch)
	}
	ch <- pe.up
	log.Infow("prometheus", zap.Bool("success", true))
}

//func (pe *PromeExporter) ServerStart(port int) error {
//	err := prometheus.Register(pe)
//	if err != nil {
//		return err
//	}
//	addr := fmt.Sprintf(":%d", port)
//	log.Infow("prometheus", zap.String("addr", addr))
//	pe.srv = &http.Server{Addr: addr, Handler: promhttp.Handler()}
//	return pe.srv.ListenAndServe()
//}
//
//func (pe *PromeExporter) ServerClose() error {
//	log.Infow("prometheus", zap.Bool("server_close", true))
//	return pe.srv.Close()
//}

//func (pe *PromeExporter) Register(router *gin.RouterGroup)  {
//	router.Handle()
//}

//func init() {
//	promeExport := NewExporter("prome")
//	go func() {
//		err := promeExport.ServerStart(4016)
//		if err != nil {
//			panic(err)
//		}
//	}()
//}
