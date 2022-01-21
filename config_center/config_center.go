package config_center

import (
	"context"
)

const DefaultConfigCenterType = "etcd"

var clientMap map[string]Client

func RegistryClient(name string, client Client) {
	clientMap[name] = client
}

var defaultClient Client

func Init(name string) {
	if name != "" {
		defaultClient = clientMap[name]
	}
	defaultClient = clientMap[DefaultConfigCenterType]
}

func Close() error {
	return defaultClient.Close()
}

func Set(ctx context.Context, key, val string) error {
	return defaultClient.Set(ctx, key, val)
}

func Get(ctx context.Context, key string) (*Item, error) {
	return defaultClient.Get(ctx, key)
}

func Watch(ctx context.Context, item *Item, callback func(item *Item)) {
	defaultClient.Watch(ctx, item, callback)
}
