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
	regiones := model.List(database.Db)
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
		return c.JSON(err)
	}
	model := models.Region(serializer)
	err = model.Insert(database.Db)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(serializer)
}
