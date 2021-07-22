package consumer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/azdbaaaaaa/util/log"
	"strings"
	"time"
)

type Config struct {
	Broker     string `json:"broker"` // 以,分割的列表
	Topic      string `json:"topic"`
	Group      string `json:"group"`
	Offset     string `json:"offset"`                                 // one of `latest` and `earliest`
	AutoCommit int32  `json:"auto_commit" mapstructure:"auto_commit"` //是否自动提交偏移量, 0:true, 1:false
}

type ConsumerGroup struct {
	group sarama.ConsumerGroup
}

func NewConsumerGroup(brokers []string, groupId string, offset int64, autoCommit bool) *ConsumerGroup {
	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Group.Session.Timeout = 20 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 6 * time.Second
	config.Consumer.IsolationLevel = sarama.ReadCommitted
	config.Consumer.Offsets.AutoCommit.Enable = autoCommit
	config.Consumer.Offsets.Initial = offset
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupId, config)
	if err != nil {
		log.Panicf("error to NewConsumerGroup %v", err)
		return nil
	}
	return &ConsumerGroup{group: consumerGroup}
}

func (CG *ConsumerGroup) Consume(topics []string, ctx context.Context, handler ConsumerGroupHandler) {
	defer CG.group.Close()
	for {
		if err := CG.group.Consume(ctx, topics, handler); err != nil {
			log.Errorw("ConsumerGroup.Consume", "err", err)
		}
		if ctx.Err() != nil {
			log.Errorw("ctx.Err", ctx.Err())
			return
		}
	}
}

/*type ConsumerGroupHandler struct {
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
	for message := range claim.Messages() {
		err := CGH.cs.ReceiveMessages(message)
		if err != nil {
			log.Errorf("ConsumerGroupHandler.ReceiveMessages: %s, error: %v", string(message.Value), err)
		} else {
			session.MarkMessage(message, "")
		}
	}
	return nil
}
*/
type ConsumerGroupHandler interface {
	//ReceiveMessages(message *sarama.ConsumerMessage) (err error)
	Setup(sarama.ConsumerGroupSession) error
	Cleanup(sarama.ConsumerGroupSession) error
	ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error
}

func Run(ctx context.Context, conf *Config, handler ConsumerGroupHandler) {
	var (
		offset     int64
		autoCommit bool
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

	switch conf.AutoCommit {
	case 0:
		autoCommit = true
	case 1:
		autoCommit = false
	default:
		autoCommit = true
	}
	cg := NewConsumerGroup(brokers, conf.Group, offset, autoCommit)
	go cg.Consume([]string{conf.Topic}, ctx, handler)
	return
}
