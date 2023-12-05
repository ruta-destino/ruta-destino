package router

import (
	"ruta-destino/pkg/router/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	routes.SetupRegionRoutes(api)
}
