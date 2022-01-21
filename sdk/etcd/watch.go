package etcd

import (
	"context"
	"github.com/xx/gobase/config"
)

func (e *etcd) Watch(ctx context.Context, item *config.Item, callback func(item *config.Item)) {
	if callback == nil {
		return
	}
	watchCh := e.client.Watch(ctx, item.Key)

	go func() {
		for {
			select {
			case resp := <-watchCh:
				if len(resp.Events) > 0 {
					event := resp.Events[len(resp.Events)-1]
					item.SetValue(string(event.Kv.Value))
					callback(item)
				}
			case _, ok := <-e.closeCh:
				if !ok {
					return
				}
			}
		}
	}()
}
