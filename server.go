package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	opts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.8.0")
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalf("error connecting to mongodb client: %v", err)
	}

	fmt.Println(client)
}
