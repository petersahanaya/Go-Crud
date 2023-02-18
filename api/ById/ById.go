package ById

import (
	"context"
	"fmt"
	"learn-mongo/pkg/JSON"
	"learn-mongo/pkg/models/Todos"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ById(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		ctx := context.TODO()
		Id := req.URL.Query().Get("id")
		fmt.Println("This is Id : ", Id)

		todos := Todos.TodosCollection()
		hexId, errs := primitive.ObjectIDFromHex(Id)

		if errs != nil {
			fmt.Println(errs.Error())
			http.Error(res, errs.Error(), 400)
			return
		}
		fmt.Println("This is Hex ID : ", hexId)
		var result Todos.Todo
		todos.FindOne(ctx, bson.M{"_id": hexId}).Decode(&result)

		res.Header().Set("Content-Type", "application/json")

		resultString, err := JSON.STRINGIFY(result)

		if err != nil {
			panic(err)
		}

		res.Write(resultString)
		return
	}

	if req.Method == "DELETE" {
		Id := req.URL.Query().Get("id")

		todos := Todos.TodosCollection()

		HexId, _ := primitive.ObjectIDFromHex(Id)

		_, err := todos.DeleteOne(context.TODO(), bson.M{"_id": HexId})

		if err != nil {
			res.WriteHeader(400)
			res.Write([]byte("Something Error When Make A POST.."))
			return
		}

		res.Write([]byte("Successfully Deleted..."))
		return
	}
}
