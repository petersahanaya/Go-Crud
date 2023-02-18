package mongoDB

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*20)
	defer cancel()
	optionsClient := options.Client()
	optionsClient.ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.NewClient(optionsClient)

	if err != nil {
		return nil, err
	}

	errs := client.Connect(ctx)

	if errs != nil {
		return nil, errs
	}

	return client.Database("go-testing"), nil
}
