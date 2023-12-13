package routes

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupCiudadRoutes(router fiber.Router) {
	service := services.NewCiudadService(database.Db)
	ciudad := handlers.Ciudad{Service: service}

	group := router.Group("/ciudad")
	group.Get("/", ciudad.List)
	group.Post("/", ciudad.Insert)
	group.Post("/:id", ciudad.Update)
	group.Delete("/:id", ciudad.Delete)
}
