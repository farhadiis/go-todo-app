package datastore

import (
	"context"
	"farhadiis/todo/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func Connect() *mongo.Client {
	uri := utils.GetEnv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	log.Println("MongoDB connection successfully completed.")
	return client
}

func Disconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
