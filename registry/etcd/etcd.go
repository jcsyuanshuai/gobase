package etcd

import "github.com/xx/gobase/registry"

type Etcd struct {
}

func NewEtcd() *Etcd {
	return &Etcd{}
}

func (e *Etcd) Discover() error {
	panic("implement me")
}

func (e *Etcd) Register(metas ...registry.Metadata) error {
	panic("implement me")
}

func (e *Etcd) Unregister(meta registry.Metadata) error {
	panic("implement me")
}

var _ registry.Client = &Etcd{}

func SetRegistryClient() {
	registry.SetClient("etcd", NewEtcd())
}
