package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/azdbaaaaaa/util/log"
	"strings"
)

func NewSyncProducer(conf *ProducerConfig) (producer sarama.SyncProducer, err error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err = sarama.NewSyncProducer(strings.Split(conf.Broker, ","), config)
	if err != nil {
		log.Errorw("error to create sync producer", "err", err)
		return
	}
	return
}
