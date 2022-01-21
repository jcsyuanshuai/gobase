package etcd

import "context"

func (e *etcd) Set(ctx context.Context, key, val string) error {
	checkAndSet(e)
	cli := e.client
	_, err := cli.Put(ctx, key, val)
	if err != nil {
		return err
	}
	return nil
}
