package sdk

import "context"

type Engine struct {
	Client interface{}
}

type Base interface {
	Load(ctx context.Context, name string) *Engine
	Init(name string, opts interface{}) error
}
