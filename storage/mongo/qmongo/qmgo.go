package qmongo

import (
	"context"
	"github.com/qiniu/qmgo"
)

func New(conf *qmgo.Config) (*qmgo.QmgoClient, error) {
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, conf)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
