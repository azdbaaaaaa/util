package grpc

type Config struct {
	Addr        string `json:"addr"`
	DialTimeout int    `json:"dial_timeout" mapstructure:"dial_timeout"`
}
