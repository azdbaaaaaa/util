package consumer

import (
	"context"
	"github.com/azdbaaaaaa/util/log"
	"github.com/azdbaaaaaa/util/prome"
	"time"

	"github.com/Shopify/sarama"
)

type ConsumerGroup struct {
	group sarama.ConsumerGroup
}

func Run(ctx context.Context, conf *Config, cs ConsumerService) {
	cg := New(conf)
	go cg.Consume([]string{conf.Topic}, ctx, cs)
	return
}

func New(conf *Config) *ConsumerGroup {
	return &ConsumerGroup{group: NewConsumerGroup(conf)}
}

func (CG *ConsumerGroup) Consume(topics []string, ctx context.Context, cs ConsumerService) {
	var consumer = ConsumerGroupHandler{cs: cs}
	defer CG.group.Close()
	for {
		if err := CG.group.Consume(ctx, topics, &consumer); err != nil {
			log.Errorf("error to consume:%v", err)
			prome.NewStatUnit("group_Consume_err").MarkERR().End()
		}
		if ctx.Err() != nil {
			log.Errorf("error to ctx.Err:%v", ctx.Err())
			return
		}
	}
}

type ConsumerGroupHandler struct {
	cs ConsumerService
}

func (CGH *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (CGH *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (CGH *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Infof("Message claimed:  topic = %s, partitions = %d", claim.Topic(), claim.Partition())
	lastLagTime := time.Now().Unix()
	for message := range claim.Messages() {
		//fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s, partions = %d, offset = %d", string(message.Value), message.Timestamp, message.Topic, message.Partition, message.Offset)
		nowTs := time.Now().Unix()
		if nowTs-lastLagTime > 15 {
			lag := claim.HighWaterMarkOffset() - message.Offset
			prome.NewStatUnit("lag").SetCounter(int(lag)).End()
			lastLagTime = nowTs
		}
		err := CGH.cs.ReceiveMessages(message)
		if err != nil {
			session.MarkMessage(message, "")
		}
	}

	return nil
}

type ConsumerService interface {
	ReceiveMessages(message *sarama.ConsumerMessage) (err error)
}
