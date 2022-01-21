package etcd

import (
	"context"
	"github.com/xx/gobase/config"
)

func (e *etcd) Get(ctx context.Context, key string) (*config.Item, error) {
	cli := e.client
	get, err := cli.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	ret := &config.Item{
		Key: key,
	}
	ret.SetValue(string(get.Kvs[0].Value))
	return ret, nil
}
