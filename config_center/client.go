package config_center

import "context"

type Client interface {
	Close() error

	Set(ctx context.Context, key, val string) error

	Get(ctx context.Context, key string) (*Item, error)

	Watch(ctx context.Context, item *Item, callback func(item *Item))
}
