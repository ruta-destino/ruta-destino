package handlers

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/router/serializers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Region struct{}

func (*Region) List(c *fiber.Ctx) error {
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

func (*Region) Insert(c *fiber.Ctx) error {
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
	serializer.Id = model.Id
	return c.JSON(serializer)
}

func (*Region) Update(c *fiber.Ctx) error {
	serializer := serializers.Region{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	serializer.Id = uint(id)
	model := models.Region(serializer)
	err = model.Update(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't update region entry",
		})
	}
	return c.JSON(serializer)
}

func (*Region) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	model := models.Region{Id: uint(id)}
	err = model.Delete(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't delete region entry",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}
