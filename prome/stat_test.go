package prome

import (
	"reflect"
	"testing"
)

func TestStatUnit_MarkOK(t *testing.T) {
	type fields struct {
		name       string
		counter    int
		status     Status
		timeRecord int64
	}
	tests := []struct {
		name   string
		fields fields
		want   *StatUnit
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &StatUnit{
				name:       tt.fields.name,
				counter:    tt.fields.counter,
				status:     tt.fields.status,
				timeRecord: tt.fields.timeRecord,
			}
			if got := su.MarkOK(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatUnit.MarkOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatUnit_MarkERR(t *testing.T) {
	type fields struct {
		name       string
		counter    int
		status     Status
		timeRecord int64
	}
	tests := []struct {
		name   string
		fields fields
		want   *StatUnit
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &StatUnit{
				name:       tt.fields.name,
				counter:    tt.fields.counter,
				status:     tt.fields.status,
				timeRecord: tt.fields.timeRecord,
			}
			if got := su.MarkERR(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatUnit.MarkERR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatUnit_SetCounter(t *testing.T) {
	type fields struct {
		name       string
		counter    int
		status     Status
		timeRecord int64
	}
	type args struct {
		counter int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StatUnit
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := &StatUnit{
				name:       tt.fields.name,
				counter:    tt.fields.counter,
				status:     tt.fields.status,
				timeRecord: tt.fields.timeRecord,
			}
			if got := su.SetCounter(tt.args.counter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatUnit.SetCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStatUnit(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *StatUnit
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatUnit(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatDescInfo_String(t *testing.T) {
	type fields struct {
		Name       string
		Total      int
		QPS        float32
		Counter    int
		AvgCounter float32
		AvgCost    float64
		MaxCost    float64
		Status     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := &StatDescInfo{
				Name:       tt.fields.Name,
				Total:      tt.fields.Total,
				QPS:        tt.fields.QPS,
				Counter:    tt.fields.Counter,
				AvgCounter: tt.fields.AvgCounter,
				AvgCost:    tt.fields.AvgCost,
				MaxCost:    tt.fields.MaxCost,
				Status:     tt.fields.Status,
			}
			if got := info.String(); got != tt.want {
				t.Errorf("StatDescInfo.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromeStat_Collect(t *testing.T) {
	type fields struct {
		DB                map[string][STATUSMAX]*UnitSummary
		channel           chan *StatUnit
		lastStatDescInfos []StatDescInfo
		lastCollectTime   int64
		lastGenTime       int64
	}
	tests := []struct {
		name   string
		fields fields
		want   []StatDescInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PromeStat{
				DB:                tt.fields.DB,
				channel:           tt.fields.channel,
				lastStatDescInfos: tt.fields.lastStatDescInfos,
				lastCollectTime:   tt.fields.lastCollectTime,
				lastGenTime:       tt.fields.lastGenTime,
			}
			if got := ps.Collect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PromeStat.Collect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromeStat_run(t *testing.T) {
	type fields struct {
		DB                map[string][STATUSMAX]*UnitSummary
		channel           chan *StatUnit
		lastStatDescInfos []StatDescInfo
		lastCollectTime   int64
		lastGenTime       int64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PromeStat{
				DB:                tt.fields.DB,
				channel:           tt.fields.channel,
				lastStatDescInfos: tt.fields.lastStatDescInfos,
				lastCollectTime:   tt.fields.lastCollectTime,
				lastGenTime:       tt.fields.lastGenTime,
			}
			ps.run()
		})
	}
}
