package GetAll

import (
	"context"
	"learn-mongo/pkg/JSON"
	"learn-mongo/pkg/models/Todos"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		todos := Todos.TodosCollection()

		var results []Todos.Todo

		cursor, err := todos.Find(context.TODO(), bson.D{})

		if err != nil {
			panic(err)
		}

		for cursor.Next(context.TODO()) {
			var todo Todos.Todo

			cursor.Decode(&todo)

			results = append(results, todo)
		}

		resultString, errs := JSON.STRINGIFY(results)

		if errs != nil {
			panic(errs)
		}

		res.Header().Set("Content-Type", "application/json")

		res.Write(resultString)
		return
	}
}
