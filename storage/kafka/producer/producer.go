package producer

import (
	"github.com/Shopify/sarama"
	"github.com/azdbaaaaaa/util/log"
	"strings"
)

type Config struct {
	Broker string `json:"broker"` // 以,分割的列表
	Topic  string `json:"topic"`
}

func NewProducer(conf *Config) (producer sarama.SyncProducer, err error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err = sarama.NewSyncProducer(strings.Split(conf.Broker, ","), config)
	if err != nil {
		log.Errorf("error to create sync producer: %v", err)
		return
	}
	return
}
