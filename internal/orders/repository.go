package orders

import (
	"errors"
	"time"

	"github.com/albinofreitas/linkapi-golang/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *Order) store() error {
	time.Sleep(10 * time.Second)

	collection := database.Connection.Collection("orders")

	o.Status = OrderStatus.PENDING
	_, err := collection.InsertOne(nil, o)

	if err != nil {
		return errors.New("Failed to insert order")
	}

	return nil
}

func (o *Order) getByID(id string) error {
	collection := database.Connection.Collection("orders")

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errors.New("Invalid Object ID")
	}

	err = collection.FindOne(nil, bson.M{"_id": _id}).Decode(&o)

	if err != nil {
		return errors.New("Order not found")
	}

	return nil
}

func (o *Order) update() error {
	collection := database.Connection.Collection("orders")

	_, err := collection.UpdateOne(
		nil,
		bson.M{"_id": o.ID},
		bson.D{
			{
				Key:   "$set",
				Value: bson.D{{Key: "status", Value: o.Status}},
			},
		},
	)

	if err != nil {
		return errors.New("Failed to update order")
	}

	return nil
}
