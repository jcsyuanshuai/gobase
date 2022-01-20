package demo

import (
	"context"
	"github.com/xx/gobase/conf"
	"github.com/xx/gobase/sdk/etcd"
	"github.com/xx/gobase/sdk/mongo"
)

var mongoNames = []string{
	"mongo1", "mongo2",
}

func InitMysqlEngine(ctx context.Context) {

	//mysql.Init(item.Key, item.GetValue())
	//cli.Watch(ctx, item, func(item *conf.Item) {
	//	mysql.Init(ctx, item.Key, item.GetValue())
	//})
}

func InitRedisEngine(ctx context.Context) {

	cli := etcd.New()

	for _, name := range mongoNames {
		item, err := cli.Get(ctx, name)
		err = mongo.Init(ctx, item.Key, item.GetValue())
		if err != nil {
			return
		}
		cli.Watch(ctx, item, func(item *conf.Item) {
			err = mongo.Init(ctx, item.Key, item.GetValue())
			if err != nil {
				return
			}
		})
	}
}
