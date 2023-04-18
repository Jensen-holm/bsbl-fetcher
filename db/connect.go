package db

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Connection struct {
	Uri    string
	Client *mongo.Client
}

func Connect() *Connection {
	var uri = os.Getenv("LOCAL_MONGO")
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalf("error connecting to mongodb client: %v", err)
	}

	return &Connection{
		Uri:    uri,
		Client: client,
	}
}

func (con *Connection) Find(database, collection string, filter bson.M) ([]byte, error) {
	dbCollection := con.Client.Database(database).Collection(collection)

	cur, err := dbCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error applying filter to db -> %v", err)
	}

	data, err := con.ParseCursor(cur)
	if err != nil {
		return nil, fmt.Errorf("error parsing mongodb cursor -> %v", err)
	}

	jsonData, err := bson2json(data)
	if err != nil {
		return nil, fmt.Errorf("error converting bson 2 json -> %v", err)
	}

	return jsonData, nil
}

func (con *Connection) ParseCursor(cur *mongo.Cursor) (bson.M, error) {
	defer cur.Close(context.Background())

	data := make(bson.M, 0)
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			fmt.Printf("parse cursor warning: %v", err)
		}

		for k, v := range result {
			data[k] = v
		}
	}

	return data, nil
}

func bson2json(b bson.M) ([]byte, error) {
	j, err := json.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("error marshaling bson to json: %v", err)
	}
	return j, nil
}
