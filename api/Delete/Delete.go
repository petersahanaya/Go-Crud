package Delete

import (
	"context"
	"learn-mongo/pkg/models/Todos"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOne(res http.ResponseWriter, req *http.Request) {
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
