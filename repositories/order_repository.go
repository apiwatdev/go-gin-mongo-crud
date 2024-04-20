package repositories

import (
	"context"
	"log"

	"github.com/apiwatdev/go-gin-mongo-crud/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	InsertOne(order interface{}) (*mongo.InsertOneResult, error)
}

type OrderRepositoryImp struct {
	collectionName string
	ctx            context.Context
	collection     *mongo.Collection
}

func NewOrderRepository(database *mongo.Database, ctx context.Context) OrderRepository {
	collectionName := constants.COLLECTION_ORDER
	return &OrderRepositoryImp{
		ctx:            ctx,
		collection:     database.Collection(collectionName),
		collectionName: collectionName,
	}
}

func (r *OrderRepositoryImp) InsertOne(order interface{}) (*mongo.InsertOneResult, error) {
	result, err := r.collection.InsertOne(r.ctx, order)
	if err != nil {
		log.Println("Error creating order:", err)
		return nil, err
	}
	return result, nil
}
