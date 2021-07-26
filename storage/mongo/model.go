package mongo

import (
	"time"
)

type Model struct {
	Id         int64 `bson:"_id"`
	CreateTime int64 `json:"create_time" bson:"create_time"`
	UpdateTime int64 `json:"update_time" bson:"update_time"`
}

func (m *Model) DefaultUpdateTime() {
	m.UpdateTime = time.Now().UnixNano() / 1e6
}

func (m *Model) DefaultCreateTime() {
	nowTime := time.Now().UnixNano() / 1e6
	m.CreateTime = nowTime
	m.UpdateTime = nowTime
}
