package routes

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupEmpresaRoutes(router fiber.Router) {
	service := services.NewEmpresaService(database.Db)
	empresa := handlers.Empresa{Service: service}

	group := router.Group("/empresa")
	group.Get("/", empresa.List)
	group.Post("/", empresa.Insert)
	group.Post("/:id", empresa.Update)
	group.Delete("/:id", empresa.Delete)
	group.Get("/:id", empresa.Get)
	group.Get("/:id/terminal", empresa.ListTerminales)
	group.Post("/:id/terminal", empresa.LinkTerminal)
	group.Delete("/:id/terminal", empresa.UnlinkTerminal)
}
