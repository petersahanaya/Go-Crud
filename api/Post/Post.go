package Post

import (
	"context"
	"encoding/json"
	"learn-mongo/pkg/models/Todos"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostTodo(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var data Todos.Todo
		todos := Todos.TodosCollection()
		err := json.NewDecoder(req.Body).Decode(&data)

		if err != nil {
			panic(err)
		}

		todo := Todos.Todo{
			ID:        primitive.NewObjectID(),
			Title:     data.Title,
			Body:      data.Body,
			CreatedAt: time.Now().String(),
		}

		_, errs := todos.InsertOne(context.TODO(), todo)

		if errs != nil {
			panic(errs)
		}

		res.Write([]byte("Successfully Created.."))
		return
	}
}
