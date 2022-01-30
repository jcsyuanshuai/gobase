package config_center

import (
	"context"
	"github.com/xx/gobase/config"
)

const DefaultConfigCenterType = "etcd"

var clientMap map[string]Client

var defaultClient Client

func Init() {
	configType := config.GlobalConfig.ConfigCenter.Type

	if configType == "etcd" {
		defaultClient = clientMap[configType]
	}
}

func SetClient(name string, cli Client) {
	clientMap[name] = cli
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
