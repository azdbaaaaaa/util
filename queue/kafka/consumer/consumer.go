package consumer

import (
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

func NewConsumerGroup(conf *Config) sarama.ConsumerGroup {
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
	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Group.Session.Timeout = 20 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 6 * time.Second
	config.Consumer.IsolationLevel = sarama.ReadCommitted
	config.Consumer.Offsets.AutoCommit.Enable = autoCommit
	config.Consumer.Offsets.Initial = offset
	consumerGroup, err := sarama.NewConsumerGroup(brokers, conf.Group, config)
	if err != nil {
		log.Panicf("error to NewConsumerGroup %v", err)
		return nil
	}
	return consumerGroup
}