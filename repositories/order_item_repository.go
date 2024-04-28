package repositories

import (
	"context"
	"log"

	"github.com/apiwatdev/go-gin-mongo-crud/constants"
	"github.com/apiwatdev/go-gin-mongo-crud/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItemRepository interface {
	InsertOne(order interface{}) (*mongo.InsertOneResult, error)
	InsertMany(orderItems []models.OrderItem) (*mongo.InsertManyResult, error)
}

type OrderItemRepositoryImp struct {
	ctx            context.Context
	collection     *mongo.Collection
	collectionName string
}

func NewOrderItemRepository(database *mongo.Database, ctx context.Context) OrderItemRepository {
	collectionName := constants.COLLECTION_ORDER_ITEM
	return &OrderItemRepositoryImp{
		ctx:            ctx,
		collection:     database.Collection(collectionName),
		collectionName: collectionName,
	}

}
func (r *OrderItemRepositoryImp) InsertOne(order interface{}) (*mongo.InsertOneResult, error) {
	result, err := r.collection.InsertOne(r.ctx, order)
	if err != nil {
		log.Println("Error creating order:", err)
		return nil, err
	}
	return result, nil
}

func (r *OrderItemRepositoryImp) InsertMany(orderItems []models.OrderItem) (*mongo.InsertManyResult, error) {
	var itemsInterface []interface{}
	for _, item := range orderItems {
		itemsInterface = append(itemsInterface, item)
	}

	result, err := r.collection.InsertMany(r.ctx, itemsInterface)
	if err != nil {
		log.Println("Error creating order:", err)
		return nil, err
	}
	return result, nil
}
