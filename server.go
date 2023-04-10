package main

import (
	"github.com/Jensen-holm/bsbl-api/crawl"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	app := fiber.New()

	// baseball reference
	app.Get("/rf", func(c *fiber.Ctx) error {
		h := make(map[string]string)
		c.Request().Header.VisitAll(func(key, value []byte) {
			h[string(key)] = string(value)
		})

		r, err := crawl.SendGet("https://baseball-reference.com", h, &http.Client{})
		if err != nil {
			panic(err)
		}
		return c.JSON(r)
	})

	app.Listen(":3000")
}
