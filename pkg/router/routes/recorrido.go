package routes

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRecorridoRoutes(router fiber.Router) {
	service := services.NewRecorridoService(database.Db)
	recorrido := handlers.Recorrido{Service: service}

	group := router.Group("/empresa/:id_empresa/recorrido")
	group.Get("/", recorrido.List)
	group.Post("/", recorrido.Insert)
	group.Post("/:id", recorrido.Update)
	group.Delete("/:id", recorrido.Delete)
}
