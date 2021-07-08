package kafka

import (
	"context"
	"github.com/azdbaaaaaa/util/log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

type ConsumerGroup struct {
	group sarama.ConsumerGroup
}

func NewConsumerGroup(brokers []string, groupId string, offset int64) *ConsumerGroup {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Group.Session.Timeout = 20 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 6 * time.Second
	config.Consumer.IsolationLevel = sarama.ReadCommitted
	config.Consumer.Offsets.Initial = offset
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupId, config)
	if err != nil {
		log.Panicf("error to NewConsumerGroup", err)
		return nil
	}
	return &ConsumerGroup{group: consumerGroup}
}

func (CG *ConsumerGroup) Consume(topics []string, ctx context.Context, cs ConsumerService) {
	var consumer = ConsumerGroupHandler{cs: cs}
	defer CG.group.Close()
	for {
		if err := CG.group.Consume(ctx, topics, &consumer); err != nil {
			log.Errorw("error to consume", "err",err)
			//prome.NewStatUnit("group_Consume_err").MarkERR().End()
		}
		if ctx.Err() != nil {
			log.Errorw("ctx.Err", "err",ctx.Err())
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
			// TODO 延迟的prometheus 上报
			//lag := claim.HighWaterMarkOffset() - message.Offset
			//prome.NewStatUnit("lag").SetCounter(int(lag)).End()
			lastLagTime = nowTs
		}
		// TODO 消息ack
		//err := CGH.cs.ReceiveMessages(message)
		//if err != nil {
		//	session.MarkMessage(message, "")
		//}
		_ = CGH.cs.ReceiveMessages(message)
		session.MarkMessage(message, "")
	}

	return nil
}

type ConsumerService interface {
	ReceiveMessages(message *sarama.ConsumerMessage) (err error)
}

func Run(ctx context.Context, conf *ConsumerConfig, cs ConsumerService) {
	var (
		offset int64
	)
	brokers := strings.Split(conf.Broker, ",")
	switch conf.Offset {
	case "earliest":
		offset = sarama.OffsetOldest
	case "latest":
		offset = sarama.OffsetNewest
	default:
		offset = sarama.OffsetNewest
	}
	log.Infof("consumer brokers<%s>", conf.Broker)
	log.Infof("consumer group<%s>", conf.Group)
	log.Infof("consumer offset<%s>", conf.Offset)
	cg := NewConsumerGroup(brokers, conf.Group, offset)
	go cg.Consume([]string{conf.Topic}, ctx, cs)
	return
}
