package conf

import (
	"context"
	"sync"
)

var clientMap = map[string]Client{}

func RegisterClient(name string, client Client) {
	clientMap[name] = client
}

type Client interface {
	Close() error

	Set(ctx context.Context, key, val string) error

	Get(ctx context.Context, key string) (*Item, error)

	Watch(ctx context.Context, item *Item, f func(item *Item))
}

type Entry struct {
	Key   string
	Value string
}

type Item struct {
	Action    int64
	Namespace string
	Key       string
	value     string
	Entries   []*Entry
	IsDefault bool
	sync.RWMutex
}

func (v *Item) SetValue(val string) {
	v.Lock()
	defer v.Unlock()
	v.value = val
}

func (v *Item) GetValue() string {
	v.RLock()
	defer v.RUnlock()
	return v.value
}

type Auth struct {
	Username string
	Password string
}

type Center struct {
	Namespace     string
	Authorization Auth
	Endpoint      string
}
