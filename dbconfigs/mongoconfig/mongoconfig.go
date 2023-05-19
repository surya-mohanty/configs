package mongoconfig

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

var client *mongo.Client

func GetConnection() (*mongo.Client, error) {
	log.Println("Entering mongo connection")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		log.Println("Error connecting:", err)
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return client, err
}
