package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes := map[string][]string{
		"santa_rita": {
			"09:55 - Valroa - Santa Bárbara",
			"10:30 - Valroa - Santa Bárbara",
			"13:10 - NCN - Santa Bárbara",
		},
		"isla_jacob": {
			"11:10 - Aránguiz - Santa Bárbara",
		},
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bienvenido a ruta destino.")
	})

	app.Get("/routes/", func(c *fiber.Ctx) error {
		return c.JSON(routes)
	})

	app.Get("/routes/:nombre", func(c *fiber.Ctx) error {
		nombre := c.Params("nombre")
		if nombre == "" {
			return c.JSON(map[string][]string{})
		}
		return c.JSON(routes[nombre])
	})

	app.Listen(":3000")
}
