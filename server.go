package main

import (
	"github.com/Jensen-holm/bsbl-api/rf"
	"github.com/Jensen-holm/bsbl-api/user"

	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	// baseball reference
	app.Get("/bsbl-api", func(c *fiber.Ctx) error {
		h := getHeaders(c)
		usr := user.NewUser(h)
		bbref := rf.NewRef(usr)

		err := bbref.Main()
		if err != nil {
			return fiber.NewError(
				fiber.StatusBadRequest,
				err.Error(),
			)
		}

		return c.JSON(bbref.Results)
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("error listening to server: %v", err)
	}
}

func getHeaders(c *fiber.Ctx) map[string]string {
	h := make(map[string]string)
	c.Request().Header.VisitAll(func(key, value []byte) {
		h[string(key)] = string(value)
	})
	return h
}
