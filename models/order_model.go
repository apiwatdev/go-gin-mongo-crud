package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	Customer string               `bson:"customer" json:"customer"`
	Total    float64              `bson:"total" json:"total"`
	Items    []primitive.ObjectID `bson:"items,omitempty" json:"items,omitempty"`
}
