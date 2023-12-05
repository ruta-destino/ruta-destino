package handlers

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/router/serializers"

	"github.com/gofiber/fiber/v2"
)

type Region struct{}

func (r *Region) List(c *fiber.Ctx) error {
	model := models.Region{}
	regiones, err := model.List(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't get region entries from db",
		})
	}
	serializer := []serializers.Region{}
	for _, r := range regiones {
		s := serializers.Region(r)
		serializer = append(serializer, s)
	}
	return c.JSON(serializer)
}

func (r *Region) Insert(c *fiber.Ctx) error {
	serializer := serializers.Region{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	model := models.Region(serializer)
	err = model.Insert(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't insert `region`",
		})
	}
	return c.JSON(serializer)
}
