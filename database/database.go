package database

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Connection to db
var Connection *mongo.Database

// Connect to db
func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("DB URI"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(nil)
	if err != nil {
		log.Fatal(err)
	}

	Connection = client.Database("linkapi-golang")
	log.Println("mongodb connected")
}
