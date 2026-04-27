package http

const (
	DefaultClientDialTimeoutInSec = 10
)

type ServerConfig struct {
	Addr            string `json:"addr"`
	ShutdownTimeout int    `json:"shutdown_timeout" mapstructure:"shutdown_timeout"`
	Pprof           bool   `json:"pprof" mapstructure:"pprof"`
}

type ClientConfig struct {
	Addr        string `json:"addr"`
	DialTimeout int    `json:"dial_timeout" mapstructure:"dial_timeout"`
}
