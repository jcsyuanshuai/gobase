package etcd

import (
	"context"
	"fmt"
	"github.com/xx/gobase/config_center"
	"testing"
	"time"
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

	go func() {
		cli1 := New()
		defer cli1.Close()

		for i := 0; i < 8; i++ {
			time.Sleep(1000)
			cli1.Set(context.TODO(), "test-key", fmt.Sprintf("test-val-%d", i))
		}
	}()

	cli := New()
	defer cli.Close()
	ctx := context.TODO()

	item, err := cli.Get(ctx, "test-key")
	if err != nil {
		return
	}
	cli.Watch(ctx, item, func(item *config_center.Item) {
		fmt.Println(item.GetValue() == "test-val-9")
	})

	for i := 0; i < 10; i++ {
		time.Sleep(1)
		break
	}
}
