package etcd

import "github.com/xx/gobase/registry"

type Factory struct {
}

func (f Factory) Create() registry.Client {
	return NewEtcd()
}

var _ registry.Factory = &Factory{}
