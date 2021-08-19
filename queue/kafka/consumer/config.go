package consumer

type Config struct {
	Broker     string `json:"broker"` // 以,分割的列表
	Topic      string `json:"topic"`
	Group      string `json:"group"`
	Offset     string `json:"offset"`                                 //  earliest or latest  default is latest
	AutoCommit bool   `json:"auto_commit" mapstructure:"auto_commit"` // 是否自动提交偏移量
}
