package grpc

const (
	DefaultClientDialTimeoutInSec = 10
)

type ServerConfig struct {
	Addr string `json:"addr"`
}

type ClientConfig struct {
	Addr        string `json:"addr"`
	DialTimeout int    `json:"dial_timeout" mapstructure:"dial_timeout"`
}
