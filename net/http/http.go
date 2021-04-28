package http

type ServerConfig struct {
	Addr            string `json:"addr"`
	ShutdownTimeout int    `json:"shutdown_timeout" mapstructure:"shutdown_timeout"`
}

type ClientConfig struct {
	Addr        string `json:"addr"`
	DialTimeout int    `json:"dial_timeout" mapstructure:"dial_timeout"`
}
