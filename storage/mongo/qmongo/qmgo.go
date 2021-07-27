package qmongo

import (
	"context"
	"github.com/qiniu/qmgo"
)

func New(conf *qmgo.Config) (*qmgo.Database, error) {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: conf.Uri})
	if err != nil {
		return nil, err
	}
	db := client.Database(conf.Database)
	return db, nil
}
