package consumer

import (
	"github.com/Shopify/sarama"
	"github.com/azdbaaaaaa/util/log"
	"strings"
	"time"
)

func NewConsumerGroup(conf *Config) sarama.ConsumerGroup {
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

	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Group.Session.Timeout = 20 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 6 * time.Second
	config.Consumer.IsolationLevel = sarama.ReadCommitted
	config.Consumer.Offsets.AutoCommit.Enable = conf.AutoCommit
	config.Consumer.Offsets.Initial = offset
	consumerGroup, err := sarama.NewConsumerGroup(brokers, conf.Group, config)
	if err != nil {
		log.Panicf("error to NewConsumerGroup %v", err)
		return nil
	}
	return consumerGroup
}
