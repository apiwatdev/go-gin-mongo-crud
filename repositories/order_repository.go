package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/apiwatdev/go-gin-mongo-crud/constants"
	"github.com/apiwatdev/go-gin-mongo-crud/dao"
	"github.com/apiwatdev/go-gin-mongo-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	FindAll(filter interface{}) ([]models.Order, error)
	FindOneById(id string) (models.Order, error)
	FindOneByIdWithItems(forderId string) (dao.OrderWithItemDao, error)
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

func (r *OrderRepositoryImp) FindAll(filter interface{}) ([]models.Order, error) {

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []models.Order
	if err = cursor.All(r.ctx, &results); err != nil {
		return nil, err
	}
	cursor.Close(r.ctx)
	return results, nil

}

func (r *OrderRepositoryImp) FindOneById(orderId string) (models.Order, error) {
	id, err := primitive.ObjectIDFromHex(orderId)

	if err != nil {
		return models.Order{}, err
	}
	var result models.Order

	err = r.collection.FindOne(r.ctx, bson.D{{Key: "_id", Value: id}}).Decode(&result)
	if err != nil {
		return models.Order{}, err
	}

	return result, nil
}

func (r *OrderRepositoryImp) FindOneByIdWithItems(orderID string) (dao.OrderWithItemDao, error) {
	objectID, err := primitive.ObjectIDFromHex(orderID)

	fmt.Println(objectID)
	if err != nil {
		return dao.OrderWithItemDao{}, fmt.Errorf("invalid ObjectID: %v", err)
	}

	pipeline := mongo.Pipeline{
		bson.D{
			{
				Key: "$match", Value: bson.D{
					{Key: "_id", Value: primitive.ObjectID(objectID)},
				},
			},
		},
		bson.D{
			{
				Key: "$lookup", Value: bson.D{
					{Key: "from", Value: "order_items"},
					{Key: "localField", Value: "_id"},
					{Key: "foreignField", Value: "order_id"},
					{Key: "as", Value: "items"},
				},
			},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "customer", Value: 1},
				{Key: "total", Value: 1},
				{Key: "items._id", Value: 1},
				{Key: "items.product", Value: 1},
				{Key: "items.quantity", Value: 1},
				{Key: "items.price", Value: 1},
				{Key: "items.order_id", Value: 1},
			},
			},
		},
	}

	cursor, err := r.collection.Aggregate(r.ctx, pipeline)
	if err != nil {
		return dao.OrderWithItemDao{}, err
	}
	defer cursor.Close(r.ctx)

	var result dao.OrderWithItemDao

	if cursor.Next(r.ctx) {
		if err := cursor.Decode(&result); err != nil {
			return dao.OrderWithItemDao{}, err
		}
	} else {
		return dao.OrderWithItemDao{}, fmt.Errorf("order not found")
	}

	return result, err

}
