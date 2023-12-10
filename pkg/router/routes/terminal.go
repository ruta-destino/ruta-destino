package routes

import (
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupTerminalRoutes(router fiber.Router) {
	group := router.Group("/terminal")

	terminal := handlers.Terminal{}
	group.Get("/", terminal.List)
	group.Post("/", terminal.Insert)
	group.Post("/:id", terminal.Update)
	group.Delete("/:id", terminal.Delete)
}
