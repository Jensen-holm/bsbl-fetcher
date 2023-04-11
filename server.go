package main

import (
	"github.com/Jensen-holm/bsbl-api/crawl"
	"github.com/Jensen-holm/bsbl-api/session"
	"github.com/Jensen-holm/bsbl-api/user"
	"time"

	"github.com/gofiber/fiber/v2"
	"log"
)

// tehe

func main() {
	app := fiber.New()

	app.Get("/bsbl-api", func(c *fiber.Ctx) error {

		st := time.Now()

		h := crawl.GetHeaders(c)
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

		et := float64(time.Since(st))

		result := session.NewResult(et, bsbl.Results())

		return c.JSON(result.Unpack())
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("error listening to server: %v", err)
	}
}
