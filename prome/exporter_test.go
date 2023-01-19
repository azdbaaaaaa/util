package prome

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestNewExporter(t *testing.T) {
	type args struct {
		ns string
	}
	tests := []struct {
		name string
		args args
		want *PromeExporter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExporter(tt.args.ns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExporter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGaugeVec(t *testing.T) {
	type args struct {
		ns   string
		name string
		help string
	}
	tests := []struct {
		name string
		args args
		want *prometheus.GaugeVec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGaugeVec(tt.args.ns, tt.args.name, tt.args.help); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGaugeVec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromeExporter_Describe(t *testing.T) {
	type fields struct {
		srv         *http.Server
		up          prometheus.Gauge
		total       *prometheus.GaugeVec
		qps         *prometheus.GaugeVec
		avgCost     *prometheus.GaugeVec
		maxCost     *prometheus.GaugeVec
		costSummery *prometheus.SummaryVec
		counter     *prometheus.GaugeVec
		avgCounter  *prometheus.GaugeVec
	}
	type args struct {
		ch chan<- *prometheus.Desc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pe := &PromeExporter{
				srv:         tt.fields.srv,
				up:          tt.fields.up,
				total:       tt.fields.total,
				qps:         tt.fields.qps,
				avgCost:     tt.fields.avgCost,
				maxCost:     tt.fields.maxCost,
				costSummery: tt.fields.costSummery,
				counter:     tt.fields.counter,
				avgCounter:  tt.fields.avgCounter,
			}
			pe.Describe(tt.args.ch)
		})
	}
}

func TestPromeExporter_reset(t *testing.T) {
	type fields struct {
		srv         *http.Server
		up          prometheus.Gauge
		total       *prometheus.GaugeVec
		qps         *prometheus.GaugeVec
		avgCost     *prometheus.GaugeVec
		maxCost     *prometheus.GaugeVec
		costSummery *prometheus.SummaryVec
		counter     *prometheus.GaugeVec
		avgCounter  *prometheus.GaugeVec
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pe := &PromeExporter{
				srv:         tt.fields.srv,
				up:          tt.fields.up,
				total:       tt.fields.total,
				qps:         tt.fields.qps,
				avgCost:     tt.fields.avgCost,
				maxCost:     tt.fields.maxCost,
				costSummery: tt.fields.costSummery,
				counter:     tt.fields.counter,
				avgCounter:  tt.fields.avgCounter,
			}
			pe.reset()
		})
	}
}

func TestPromeExporter_Collect(t *testing.T) {
	type fields struct {
		srv         *http.Server
		up          prometheus.Gauge
		total       *prometheus.GaugeVec
		qps         *prometheus.GaugeVec
		avgCost     *prometheus.GaugeVec
		maxCost     *prometheus.GaugeVec
		costSummery *prometheus.SummaryVec
		counter     *prometheus.GaugeVec
		avgCounter  *prometheus.GaugeVec
	}
	type args struct {
		ch chan<- prometheus.Metric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pe := &PromeExporter{
				srv:         tt.fields.srv,
				up:          tt.fields.up,
				total:       tt.fields.total,
				qps:         tt.fields.qps,
				avgCost:     tt.fields.avgCost,
				maxCost:     tt.fields.maxCost,
				costSummery: tt.fields.costSummery,
				counter:     tt.fields.counter,
				avgCounter:  tt.fields.avgCounter,
			}
			pe.Collect(tt.args.ch)
		})
	}
}

func TestPromeExporter_ServerStart(t *testing.T) {
	type fields struct {
		srv         *http.Server
		up          prometheus.Gauge
		total       *prometheus.GaugeVec
		qps         *prometheus.GaugeVec
		avgCost     *prometheus.GaugeVec
		maxCost     *prometheus.GaugeVec
		costSummery *prometheus.SummaryVec
		counter     *prometheus.GaugeVec
		avgCounter  *prometheus.GaugeVec
	}
	type args struct {
		port int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pe := &PromeExporter{
				srv:         tt.fields.srv,
				up:          tt.fields.up,
				total:       tt.fields.total,
				qps:         tt.fields.qps,
				avgCost:     tt.fields.avgCost,
				maxCost:     tt.fields.maxCost,
				costSummery: tt.fields.costSummery,
				counter:     tt.fields.counter,
				avgCounter:  tt.fields.avgCounter,
			}
			if err := pe.ServerStart(tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("PromeExporter.ServerStart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPromeExporter_ServerClose(t *testing.T) {
	type fields struct {
		srv         *http.Server
		up          prometheus.Gauge
		total       *prometheus.GaugeVec
		qps         *prometheus.GaugeVec
		avgCost     *prometheus.GaugeVec
		maxCost     *prometheus.GaugeVec
		costSummery *prometheus.SummaryVec
		counter     *prometheus.GaugeVec
		avgCounter  *prometheus.GaugeVec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pe := &PromeExporter{
				srv:         tt.fields.srv,
				up:          tt.fields.up,
				total:       tt.fields.total,
				qps:         tt.fields.qps,
				avgCost:     tt.fields.avgCost,
				maxCost:     tt.fields.maxCost,
				costSummery: tt.fields.costSummery,
				counter:     tt.fields.counter,
				avgCounter:  tt.fields.avgCounter,
			}
			if err := pe.ServerClose(); (err != nil) != tt.wantErr {
				t.Errorf("PromeExporter.ServerClose() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
