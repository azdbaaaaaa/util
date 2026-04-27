package kafka

type Message struct {
	ID        string      `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}
