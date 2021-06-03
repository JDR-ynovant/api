package repository

import (
	"context"
	"fmt"
	"github.com/JDR-ynovant/api/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type Repository interface {
	FindById(id string) interface{}
}

var connection *mongo.Client
var connectionError error

func init() {
	connection, connectionError = getMongoConnection()

	if connectionError != nil {
		panic(connectionError)
	}
}

func GetMongoConnection() *mongo.Client {
	return connection
}

func CloseConnection() {
	_ = connection.Disconnect(context.Background())
}

func getMongoConnection() (*mongo.Client, error) {
	config := internal.GetConfig()
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", config.DbUser, config.DbPass, config.DbHost, config.DbPort)))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB.")
	return client, err
}

func GetMongoDbCollection(CollectionName string) (*mongo.Collection, error) {
	config := internal.GetConfig()
	client := GetMongoConnection()

	collection := client.Database(config.DbName).Collection(CollectionName)

	return collection, nil
}
