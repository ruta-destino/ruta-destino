package routes

import (
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRegionRoutes(router fiber.Router) {
	regionGroup := router.Group("/region")

	region := handlers.Region{}
	regionGroup.Get("/", region.List)
	regionGroup.Post("/", region.Insert)
}
