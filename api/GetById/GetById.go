package GetById

import (
	"context"
	"learn-mongo/pkg/JSON"
	"learn-mongo/pkg/models/Todos"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetById(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		ctx := context.TODO()
		Id := req.URL.Query().Get("id")

		todos := Todos.TodosCollection()
		hexId, _ := primitive.ObjectIDFromHex(Id)
		var result Todos.Todo
		todos.FindOne(ctx, bson.M{"_id": hexId}).Decode(&result)

		res.Header().Set("Content-Type", "application/json")

		resultString, err := JSON.STRINGIFY(result)

		if err != nil {
			panic(err)
		}

		res.Write(resultString)
	}
}
