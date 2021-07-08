package kafka

type ConsumerConfig struct {
	Broker string `json:"broker"` // 以,分割的列表
	Topic  string `json:"topic"`
	Group  string `json:"group"`
	Offset string `json:"offset"` // one of `latest` and `earliest`
}

type ProducerConfig struct {
	Broker string `json:"broker"` // 以,分割的列表
	Topic  string `json:"topic"`
	//ReturnSuccess bool   `json:"return_success"`
	//ReturnError   bool   `json:"return_error"`
}
