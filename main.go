package main

import (
	"log"
	"os"
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conn, ok := os.LookupEnv("CONNECTION_STRING")
	if !ok {
		log.Fatal("CONNECTION_STRING no está definida")
	}

	_, err := database.Open(conn)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
