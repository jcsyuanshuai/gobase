package etcd

import (
	"context"
	"github.com/xx/gobase/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type etcd struct {
	opts    Options
	client  *clientv3.Client
	closeCh chan struct{}
}

func (e *etcd) Close() error {
	return e.client.Close()
}

func (e *etcd) Set(ctx context.Context, key, val string) error {
	checkAndSet(e)
	cli := e.client
	_, err := cli.Put(ctx, key, val)
	if err != nil {
		return err
	}
	return nil
}

func checkAndSet(e *etcd) {
	if e.client == nil {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints: []string{e.opts.Endpoint},
			Username:  e.opts.Username,
			Password:  e.opts.Password,
		})
		if err != nil {
			return
		}
		e.client = cli
	}
}

func (e *etcd) Get(ctx context.Context, key string) (*conf.Item, error) {
	cli := e.client
	get, err := cli.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	ret := &conf.Item{
		Key: key,
	}
	ret.SetValue(string(get.Kvs[0].Value))
	return ret, nil
}

func (e *etcd) Watch(ctx context.Context, item *conf.Item, callback func(item *conf.Item)) {
	if callback == nil {
		return
	}
	watchCh := e.client.Watch(ctx, item.Key)

	go func() {
		for {
			select {
			case resp := <-watchCh:
				if len(resp.Events) > 0 {
					event := resp.Events[len(resp.Events)-1]
					item.SetValue(string(event.Kv.Value))
					callback(item)
				}
			case _, ok := <-e.closeCh:
				if !ok {
					return
				}
			}
		}
	}()
}

func New(opts ...Option) *etcd {
	etcd := new(etcd)
	if len(opts) == 0 {
		etcd.opts = DefaultOptions()
	} else {
		etcd.opts = NewOptions(opts...)
	}
	return etcd
}

var _ conf.Client = &etcd{}
