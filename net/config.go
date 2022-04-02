package net

import (
	"github.com/azdbaaaaaa/util/log"
	"github.com/azdbaaaaaa/util/net/grpc"
	"github.com/azdbaaaaaa/util/net/http"
)

type Config struct {
	Author  string            `json:"author" mapstructure:"author"`
	Service string            `json:"service" mapstructure:"service"`
	Log     log.LoggerOption  `json:"log" mapstructure:"log"`
	GRPC    grpc.ServerConfig `json:"grpc" mapstructure:"grpc"`
	HTTP    http.ServerConfig `json:"http" mapstructure:"http"`
}
