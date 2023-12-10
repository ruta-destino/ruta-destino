package routes

import (
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupEmpresaRoutes(router fiber.Router) {
	group := router.Group("/empresa")

	empresa := handlers.Empresa{}
	group.Get("/", empresa.List)
	group.Post("/", empresa.Insert)
	group.Post("/:id", empresa.Update)
	group.Delete("/:id", empresa.Delete)
}
