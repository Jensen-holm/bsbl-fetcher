package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		h := c.Request().Header.String()
		fmt.Println(h)

		//r, err := crawl.SendGet("https://baseball-reference.com", h, &http.Client{})
		//if err != nil {
		//	panic(err)
		//}
		return c.JSON(h)
	})

	app.Listen(":3000")
}
