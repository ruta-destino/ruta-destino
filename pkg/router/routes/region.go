package routes

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRegionRoutes(router fiber.Router) {
	service := services.NewRegionService(database.Db)
	region := handlers.Region{Service: service}

	group := router.Group("/region")
	group.Get("/", region.List)
	group.Post("/", region.Insert)
	group.Post("/:id", region.Update)
	group.Delete("/:id", region.Delete)
}
