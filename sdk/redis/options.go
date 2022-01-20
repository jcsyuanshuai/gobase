package redis

import "github.com/go-redis/redis/v8"

type Options struct {
}

func (o *Options) Connect() *redis.Client {
	return nil
}
