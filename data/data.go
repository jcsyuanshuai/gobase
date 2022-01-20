package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"
import "go.mongodb.org/mongo-driver/mongo"
import "github.com/go-redis/redis/v8"

type Buildable interface {
	BuildMysqlEngine(opts *MysqlOpts) (*Database, error)
	BuildRedisEngine(opts *RedisOpts) (*Database, error)
	BuildMongoEngine(opts *MongoOpts) (*Database, error)
}

type Builder struct {
}

func (b *Builder) BuildMysqlEngine(opts *MysqlOpts) (*Database, error) {
	db, err := gorm.Open(mysql.Open(""))
	if err != nil {
		return &Database{}, err
	}
	return &Database{
		engine: db,
	}, nil
}

func (b *Builder) BuildRedisEngine(opts *RedisOpts) (*Database, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     opts.Host,
		Password: opts.Password,
		DB:       opts.Database,
	})

	_, err := db.Ping(context.Background()).Result()
	if err != nil {
		return &Database{}, err
	}
	return &Database{
		engine: db,
	}, nil
}

func (b *Builder) BuildMongoEngine(opts *MongoOpts) (*Database, error) {
	db, err := mongo.Connect(context.Background(), options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s", opts.AppUrl)),
	)
	if err != nil {
		return &Database{}, nil
	}
	return &Database{
		engine: db,
	}, nil
}

type Database struct {
	engine interface{}
}

type Newable interface {
	New(b Buildable) (*Database, error)
}
