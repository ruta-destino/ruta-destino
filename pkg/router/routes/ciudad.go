package routes

import (
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupCiudadRoutes(router fiber.Router) {
	group := router.Group("/ciudad")

	ciudad := handlers.Ciudad{}
	group.Get("/", ciudad.List)
	group.Post("/", ciudad.Insert)
	group.Post("/:id", ciudad.Update)
	group.Delete("/:id", ciudad.Delete)
}
