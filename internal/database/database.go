package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credentials := options.Credential{
		Username: "root",
		Password: "12345678",
	}

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials),
	)
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	collection := client.Database("library").Collection("books")

	return collection
}
