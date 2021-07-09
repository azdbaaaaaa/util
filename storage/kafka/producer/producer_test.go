package producer

import (
	"github.com/Shopify/sarama"
	"github.com/azdbaaaaaa/util/log"
	"testing"
)

func TestNewProducer(t *testing.T) {
	cfg := &Config{
		Broker: "b-2.dev-ficool.oei4cx.c6.kafka.eu-west-1.amazonaws.com:9092,b-3.dev-ficool.oei4cx.c6.kafka.eu-west-1.amazonaws.com:9092,b-1.dev-ficool.oei4cx.c6.kafka.eu-west-1.amazonaws.com:9092",
		Topic:  "ficool-balance",
	}
	producer, err := NewProducer(cfg)
	if err != nil {
		log.Errorw("NewProducer", "error", err)
		return
	}
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic:     cfg.Topic,
		Value:     sarama.StringEncoder("test producer"),
	})
	if err != nil {
		log.Errorf("SendMessage", "error", err)
		return
	}
}
