package main

import (
	"log"
	"ruta-destino/pkg/config"
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	connectionString := config.GetConnectionString()
	_, err := database.Open(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
