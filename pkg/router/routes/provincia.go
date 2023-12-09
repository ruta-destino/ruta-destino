package routes

import (
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupProvinciaRoutes(router fiber.Router) {
	group := router.Group("/provincia")

	provincia := handlers.Provincia{}
	group.Get("/", provincia.List)
	group.Post("/", provincia.Insert)
	group.Post("/:id", provincia.Update)
	group.Delete("/:id", provincia.Delete)
}
