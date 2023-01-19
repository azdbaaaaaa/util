package consumer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/azdbaaaaaa/util/log"
	"testing"
)

type TestService struct {

}

func (ts *TestService) ReceiveMessages(msg *sarama.ConsumerMessage) (err error) {
	log.Infof("msg: %s", string(msg.Value))
	return
}

func TestRun(t *testing.T) {
	cfg := &Config{
		Broker: "b-2.dev-ficool.oei4cx.c6.kafka.eu-west-1.amazonaws.com:9092,b-3.dev-ficool.oei4cx.c6.kafka.eu-west-1.amazonaws.com:9092,b-1.dev-ficool.oei4cx.c6.kafka.eu-west-1.amazonaws.com:9092",
		Topic:  "ficool-balance",
		Group:  "ficool.balance.test",
		Offset: "latest",
	}
	ts := new(TestService)
	Run(context.Background(), cfg, ts)
}
