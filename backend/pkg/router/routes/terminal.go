package routes

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupTerminalRoutes(router fiber.Router) {
	service := services.NewTerminalService(database.Db)
	terminal := handlers.Terminal{Service: service}

	group := router.Group("/terminal")
	group.Get("/", terminal.List)
	group.Post("/", terminal.Insert)
	group.Post("/:id", terminal.Update)
	group.Delete("/:id", terminal.Delete)
	group.Get("/:id", terminal.Get)
	group.Get("/:id/empresa", terminal.ListEmpresa)
}
