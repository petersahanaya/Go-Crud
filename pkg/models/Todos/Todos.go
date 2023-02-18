package Todos

import (
	"fmt"
	"learn-mongo/pkg/mongoDB"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Title     string             `bson:"title" json:"title"`
	Body      string             `bson:"body" json:"body"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
}

func TodosCollection() *mongo.Collection {
	client, err := mongoDB.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	todos := client.Collection("todos")

	return todos
}
