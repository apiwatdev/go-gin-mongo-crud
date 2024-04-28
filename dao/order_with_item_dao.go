package dao

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItemDao struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	OrderID  primitive.ObjectID `bson:"order_id"  json:"orderId,omitempty"`
	Product  string             `bson:"product" json:"product"`
	Quantity int                `bson:"quantity" json:"quantity"`
	Price    float64            `bson:"price" json:"price"`
}

type OrderWithItemDao struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Customer string             `bson:"customer" json:"customer"`
	Total    float64            `bson:"total" json:"total"`
	Items    []OrderItemDao     `bson:"items" json:"items" `
}

func (o *OrderWithItemDao) MarshalJSON() ([]byte, error) {
	type Alias OrderWithItemDao
	if o.Items == nil {
		o.Items = []OrderItemDao{}
	}
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(o),
	})
}
