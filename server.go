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
		h := getHeaders(c)
		r, err := crawl.SendGet("https://baseball-reference.com", h, &http.Client{})
		if err != nil {
			panic(err)
		}
		return c.JSON(r)
	})

	// fangraphs
	app.Get("/fg", func(c *fiber.Ctx) error {
		h := getHeaders(c)
		r, err := crawl.SendGet("https://baseball-reference.com", h, &http.Client{})
		if err != nil {
			panic(err)
		}
		return c.JSON(r)
	})

	// retro sheet
	app.Get("/rs", func(c *fiber.Ctx) error {
		h := getHeaders(c)
		r, err := crawl.SendGet("https://baseball-reference.com", h, &http.Client{})
		if err != nil {
			panic(err)
		}
		return c.JSON(r)
	})

	// baseball savant
	app.Get("/sv", func(c *fiber.Ctx) error {
		h := getHeaders(c)
		r, err := crawl.SendGet("https://baseball-reference.com", h, &http.Client{})
		if err != nil {
			panic(err)
		}
		return c.JSON(r)
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func getHeaders(c *fiber.Ctx) map[string]string {
	h := make(map[string]string)
	c.Request().Header.VisitAll(func(key, value []byte) {
		h[string(key)] = string(value)
	})
	return h
}
