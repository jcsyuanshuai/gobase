package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

var poolMap sync.Map

func Load(ctx context.Context, name string) (*redis.Client, error) {
	if cli, ok := poolMap.Load(name); ok {
		if db, ok := cli.(*redis.Client); ok {
			return db, nil
		}
	}
	return nil, nil
}

func Init(name string, opts *Options) {
	engine := opts.Connect()
	poolMap.Store(name, engine)
}
