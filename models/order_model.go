package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Customer string             `bson:"customer"`
	Total    float64            `bson:"total"`
}
