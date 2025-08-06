package producer

import (
	"github.com/Shopify/sarama"
	"github.com/azdbaaaaaa/util/log"
	"testing"
)

func TestNewProducer(t *testing.T) {
	//broker := "b-2.nsmspdev.04141j.c3.kafka.us-west-2.amazonaws.com:9096,b-1.nsmspdev.04141j.c3.kafka.us-west-2.amazonaws.com:9096"
	cfg := &Config{
		Broker: "127.0.0.1:9092",
		Topic:  "dev_test_01",
	}
	producer, err := NewProducer(cfg)
	if err != nil {
		log.Errorw("NewProducer", "error", err)
		return
	}
	defer producer.Close()
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: cfg.Topic,
		Value: sarama.StringEncoder("test producer"),
	})
	if err != nil {
		log.Errorf("SendMessage", "error", err)
		return
	}
}
