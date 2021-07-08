package redis

import (
	"context"
	"github.com/azdbaaaaaa/util/log"
	"github.com/go-redis/redis/v8"
)

//var (
//	RedisPool *redis.Pool
//	RedisOnce sync.Once
//)

type Config struct {
	Addr string `json:"addr"`
	Auth string `json:"auth"`
	DB   int    `json:"db"`
}

func New(cfg *Config) (rdb *redis.Client, err error) {
	option := &redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Auth, // no password set
		DB:       cfg.DB,   // use default DB
	}
	rdb = redis.NewClient(option)

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("New redis error,option:%#v, %v", *option, err)
		return
	}
	// tracing
	// https://github.com/open-telemetry/opentelemetry-go/blob/master/README.md
	//rdb.AddHook(redisext.OpenTelemetryHook{})
	return
}
