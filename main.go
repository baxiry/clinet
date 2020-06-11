package main

import (
	"github.com/gofiber/fiber"
)

func Index(c *fiber.Ctx) {
	c.JSON("this is index page")
}

func main() {
	app := fiber.New()

	app.Get("/", Index)

	app.Listen(9000)
}
