package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	uri := os.Getenv("LOCAL_MONGO")
	opts := options.Client().ApplyURI(uri)
	_, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalf("error connecting to mongodb client: %v", err)
	}

	// listen for requests from the ui
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome to the baseball fetcher")
	})

	app.Listen(":3000")
}
