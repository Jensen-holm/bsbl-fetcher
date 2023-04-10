package main

import (
	scrape "github.com/Jensen-holm/bsbl-api/crawl"
	"github.com/Jensen-holm/bsbl-api/session"
	"github.com/Jensen-holm/bsbl-api/user"

	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/bsbl-api", func(c *fiber.Ctx) error {
		h := scrape.GetHeaders(c)
		usr := user.NewUser(h)

		bsbl, err := session.WebPage(h["Web-Page"], usr)
		if err != nil {
			return fiber.NewError(
				fiber.StatusBadRequest,
				err.Error(),
			)
		}

		err = bsbl.Main()
		if err != nil {
			return fiber.NewError(
				fiber.StatusBadRequest,
				err.Error(),
			)
		}

		return c.JSON(bsbl.Results())
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("error listening to server: %v", err)
	}
}
