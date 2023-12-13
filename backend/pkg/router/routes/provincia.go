package routes

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupProvinciaRoutes(router fiber.Router) {
	service := services.NewProvinciaService(database.Db)
	provincia := handlers.Provincia{Service: service}

	group := router.Group("/provincia")
	group.Get("/", provincia.List)
	group.Post("/", provincia.Insert)
	group.Post("/:id", provincia.Update)
	group.Delete("/:id", provincia.Delete)
}
