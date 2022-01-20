package etcd

import (
	"context"
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	cli := New()
	defer cli.Close()

	ctx := context.TODO()
	err := cli.Set(ctx, "test-key", "test-val")
	if err != nil {
		return
	}
	item, err := cli.Get(ctx, "test-key")
	if err != nil {
		return
	}
	fmt.Print(item.GetValue() == "test-val")
}

func TestWatch(t *testing.T) {

}
