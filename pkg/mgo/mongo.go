package mgo

import (
	"context"
	"sync"
	"time"

	"github.com/caicloud/nirvana/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mongo mongdbConnection
var Mongo *mongo.Client
var once sync.Once

var (
	uri      = "mongodb://localhost:27017"
	userName = "mongo"
	password = "root@123"
)

func init() {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var err error
		opt := options.Client().ApplyURI(uri)

		opt.SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-1",
			AuthSource:    "admin",
			Username:      "mongo",
			Password:      password,
		})
		opt.SetMaxPoolSize(20)
		Mongo, err = mongo.Connect(ctx, opt)
		if err != nil {
			log.Error(err)
		}
	})
}
