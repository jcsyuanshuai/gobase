package mongo

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"sync"
)

var poolMap sync.Map

func init() {
	structCodec, _ := bsoncodec.NewStructCodec(bsoncodec.JSONFallbackStructTagParser)
	bson.DefaultRegistry = bson.NewRegistryBuilder().RegisterDefaultEncoder(reflect.Struct,
		structCodec).RegisterDefaultDecoder(reflect.Struct,
		structCodec).Build()
}

func Load(ctx context.Context, name string) (*mongo.Database, error) {
	if cli, ok := poolMap.Load(name); ok {
		if db, ok := cli.(*mongo.Database); ok {
			return db, nil
		}
	}
	return nil, nil
}

func Init(ctx context.Context, name, config string) error {
	opts := new(Options)
	err := json.Unmarshal([]byte(config), &opts)
	if err != nil {
		return err
	}
	db := opts.Connect()
	poolMap.Store(name, db)
	return nil
}
