package orders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/albinofreitas/linkapi-golang/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	database.Connect()
	collection = database.Connection.Collection("orders")
}

// HandlePendingOrders procces pending orders
func HandlePendingOrders() {
	cursor, err := collection.Find(nil, bson.M{"status": OrderStatus.PENDING})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(nil)

	for cursor.Next(nil) {
		var order Order

		if err = cursor.Decode(&order); err != nil {
			log.Println(err)
			continue
		}

		go handleOrder(order)
	}
}

func handleOrder(order Order) {
	reqBody, _ := json.Marshal(order)

	response, err := http.Post(
		"API TO INTEGRATE ORDER",
		"application/json",
		bytes.NewBuffer(reqBody),
	)

	if err != nil {
		log.Println(err)
		return
	}

	var result map[string]string

	json.NewDecoder(response.Body).Decode(&result)

	order.Status = OrderStatus.INTEGRATED

	if response.StatusCode != http.StatusOK {
		order.Status = OrderStatus.FAILED
		go sendSlackErrorNotification(order.ID, result["message"])
	}

	order.update()
	fmt.Println(order)
}

func sendSlackErrorNotification(id primitive.ObjectID, message string) {
	formatedMessage := fmt.Sprintf("Integration from order: %v failed because %v", id.Hex(), message)

	reqBody, err := json.Marshal(map[string]string{
		"message": formatedMessage,
	})

	if err != nil {
		log.Println(err)
	}

	_, err = http.Post(
		"SLACK NOTIFICATION URL",
		"application/json",
		bytes.NewBuffer(reqBody),
	)

	if err != nil {
		log.Println(err)
	}
}
