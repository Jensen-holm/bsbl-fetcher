package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	uri := os.Getenv("LOCAL_MONGO")
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalf("error connecting to mongodb client: %v", err)
	}

	usrs := client.Database("IOTPB").Collection("users")
	filter := bson.M{"name": "John Doe"}
	cursor, err := usrs.Find(context.Background(), filter)
	if err != nil {
		log.Fatalf("error finding John Doe in the db: %v", err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			// Handle error
		}
		fmt.Println(result)
	}

	// api listening for requests from the ui

}
