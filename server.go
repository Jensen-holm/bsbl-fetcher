package main

import (
	"github.com/Jensen-holm/bsbl-fetcher/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func main() {

	var con = db.Connect()

	app := fiber.New()

	app.Get("/find", func(c *fiber.Ctx) error {
		name := c.Query("name")
		filt := bson.M{"name": name}
		result, err := con.Find("IOTPB", "users", filt)
		if err != nil {
			panic(err)
		}

		c.Set("Content-Type", "application/json")
		_, err = c.Write(result)
		if err != nil {
			return err
		}

		err = c.SendStatus(fiber.StatusOK)
		if err != nil {
			return err
		}

		return nil
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("server error: %v", err)
	}

}
