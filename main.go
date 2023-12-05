package main

import (
	"log"
	"os"
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conn, ok := os.LookupEnv("CONNECTION_STRING")
	if !ok {
		log.Fatal("CONNECTION_STRING no está definida")
	}

	db, err := database.Open(conn)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// TODO: mover esto a un nuevo paquete
	type RegionSerializer struct {
		Id     uint   `json:"id"`
		Nombre string `json:"nombre"`
	}

	// TODO: mover a un nuevo paquete
	app.Get("/region", func(c *fiber.Ctx) error {
		r := models.Region{}
		regiones := r.List(db)
		return c.JSON(regiones)
	})

	app.Post("/region", func(c *fiber.Ctx) error {
		region := RegionSerializer{}
		err := c.BodyParser(&region)
		if err != nil {
			c.Status(400)
			return c.SendString("Error al agregar la región")
		}
		r := models.Region{Nombre: region.Nombre}
		err = r.Insert(db)
		if err != nil {
			c.Status(400)
			return c.SendString("Error al agregar la región")
		}
		return c.JSON(r)
	})

	app.Listen(":3000")
}
