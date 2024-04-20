package bootstraps

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client       *mongo.Client
	databaseName string
)

func InitMongo(uri string, dbname string) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	databaseName = dbname
	log.Println("Connected to MongoDB")
	return nil
}

func GetMongoClient() *mongo.Client {
	return client
}

func GetDatabase() *mongo.Database {
	return client.Database(databaseName)
}
