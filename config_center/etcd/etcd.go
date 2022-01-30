package etcd

import (
	"github.com/xx/gobase/config_center"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type etcd struct {
	opts    Options
	client  *clientv3.Client
	closeCh chan struct{}
}

var _ config_center.Client = &etcd{}

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

func New(opts ...Option) *etcd {
	etcd := new(etcd)
	if len(opts) == 0 {
		etcd.opts = DefaultOptions()
	} else {
		etcd.opts = NewOptions(opts...)
	}
	return etcd
}
