package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Config struct {
	URI            string `json:"uri" toml:"uri" yaml:"uri"`
	UserName       string `json:"username" toml:"username" yaml:"username"`
	Password       string `json:"password" toml:"password" yaml:"password"`
	DB             string `json:"db" toml:"db" yaml:"db"`
	CollectionName string `json:"collection_name" toml:"collection_name" yaml:"collection_name"`
	Timeout        int    `json:"timeout" toml:"timeout" yaml:"timeout"`
	ReadTimeout    int    `json:"read_timeout" toml:"read_timeout" yaml:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout" toml:"write_timeout" yaml:"write_timeout"`
}

func New(cfg *Config) (*mongo.Database, error) {
	if cfg.Timeout == 0 {
		cfg.Timeout = 10
	}
	// 建立mongodb连接
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	// 选择数据库
	mdb := client.Database(cfg.DB)
	return mdb, nil
}
