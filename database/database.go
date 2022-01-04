package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Sanjungliu/assesment-user-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance(ctx context.Context, config *config.Config) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DBConnectionString()))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

func OpenCollection(client *mongo.Client, dbName, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)

	return collection
}
