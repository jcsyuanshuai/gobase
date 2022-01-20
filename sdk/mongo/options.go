package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Options struct {
	Database string
	AppUrl   string
}

func (o *Options) Connect() *mongo.Database {
	dsn := fmt.Sprintf("")
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		return nil
	}

	err = db.Ping(context.TODO(), nil)
	if err != nil {
		return nil
	}
	database := db.Database(o.Database)

	return database
}
