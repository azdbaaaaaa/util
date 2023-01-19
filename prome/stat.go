package prome

import (
	"encoding/json"
	"github.com/azdbaaaaaa/util/log"
	"math"
	"time"

	"go.uber.org/zap"
)

func init() {
	go promeStater.run()
}

const (
	COLLECTINTERVAL = 30
	CHANNELBUFFSIZE = 20000
)

type Status int8

const (
	OK Status = iota
	ERR
	STATUSMAX
)

type StatUnit struct {
	name       string
	counter    int
	status     Status
	timeRecord int64
}

func (su *StatUnit) MarkOK() *StatUnit {
	su.status = OK
	return su
}

func (su *StatUnit) MarkERR() *StatUnit {
	su.status = ERR
	return su
}

func (su *StatUnit) SetCounter(counter int) *StatUnit {
	su.counter = counter
	return su
}
func (su *StatUnit) End() {
	su.timeRecord = time.Now().UnixNano() - su.timeRecord
	select {
	case promeStater.channel <- su:
	default:
		log.Warnw("prometheus", zap.String("channel", "full"))
	}
}

func NewStatUnit(name string) *StatUnit {
	return &StatUnit{
		name:       name,
		timeRecord: time.Now().UnixNano(),
	}
}

type StatDescInfo struct {
	Name       string  `json:"name"`
	Total      int     `json:"total"`
	QPS        float32 `json:"qps"`
	Counter    int     `json:"counter"`
	AvgCounter float32 `json:"avg_counter"`
	AvgCost    float64 `json:"avg_cost"`
	MaxCost    float64 `json:"max_cost"`
	Status     string  `json:"status"`
}

func (info *StatDescInfo) String() string {
	data, _ := json.Marshal(info)
	return string(data)
}

type UnitSummary struct {
	costT   int64
	total   int
	counter int
	maxCost float64
}

type PromeStat struct {
	DB                map[string][STATUSMAX]*UnitSummary
	channel           chan *StatUnit
	lastStatDescInfos []StatDescInfo
	lastCollectTime   int64
	lastGenTime       int64
}

func (ps *PromeStat) Collect() []StatDescInfo {
	ps.lastCollectTime = time.Now().UnixNano()
	return ps.lastStatDescInfos
}

func (ps *PromeStat) run() {
	ps.channel = make(chan *StatUnit, CHANNELBUFFSIZE)
	ticker := time.NewTicker(time.Second * time.Duration(COLLECTINTERVAL))
	defer ticker.Stop()
	ps.DB = make(map[string][STATUSMAX]*UnitSummary, 100)
	ps.lastGenTime = time.Now().UnixNano()
	for {
		select {
		case u := <-ps.channel:
			unitSummarys, ok := ps.DB[u.name]
			if !ok {
				unitSummarys = [STATUSMAX]*UnitSummary{}
				ps.DB[u.name] = unitSummarys
			}

			summery := unitSummarys[u.status]
			if summery == nil {
				summery = &UnitSummary{}
				unitSummarys[u.status] = summery
			}
			summery.counter += u.counter
			summery.total++
			summery.costT += u.timeRecord
			summery.maxCost = math.Max(summery.maxCost, float64(u.timeRecord))
			ps.DB[u.name] = unitSummarys

		case <-ticker.C:
			statDescInfos := make([]StatDescInfo, len(ps.DB)*int(STATUSMAX))
			i := 0
			now := time.Now().UnixNano()
			genInterval := now - ps.lastGenTime
			for k, v := range ps.DB {
				for j := 0; j < int(STATUSMAX); j++ {
					summery := v[j]
					if summery != nil && summery.total > 0 {
						switch Status(j) {
						case OK:
							statDescInfos[i].Status = "OK"
						case ERR:
							statDescInfos[i].Status = "ERR"
						}
						statDescInfos[i].Name = k
						statDescInfos[i].Total = summery.total
						statDescInfos[i].Counter = summery.counter
						statDescInfos[i].AvgCounter = float32(summery.counter) / float32(summery.total)
						statDescInfos[i].MaxCost = summery.maxCost / 1e6                                 // ms
						statDescInfos[i].AvgCost = float64(summery.costT) / 1e6 / float64(summery.total) // ms
						statDescInfos[i].QPS = float32(summery.total) / float32(genInterval/1e9)
						i++
					}
				}
			}
			ps.lastStatDescInfos = statDescInfos[:i]
			ps.lastGenTime = now

			if now-ps.lastCollectTime > 3*COLLECTINTERVAL*int64(time.Second) {
				for _, v := range statDescInfos[:i] {
					log.Infow("prometheus", zap.String("stat", v.String()))
				}
			}
			ps.DB = make(map[string][STATUSMAX]*UnitSummary, CHANNELBUFFSIZE)
		}
	}
}

var promeStater PromeStat
