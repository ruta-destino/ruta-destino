package handlers

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/router/serializers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Ciudad struct{}

func (*Ciudad) List(c *fiber.Ctx) error {
	model := models.Ciudad{}
	ciudades, err := model.List(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't get ciudad entries from db",
		})
	}
	serializer := []serializers.Ciudad{}
	for _, c := range ciudades {
		s := serializers.Ciudad(c)
		serializer = append(serializer, s)
	}
	return c.JSON(serializer)
}

func (*Ciudad) Insert(c *fiber.Ctx) error {
	serializer := serializers.Ciudad{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	model := models.Ciudad(serializer)
	err = model.Insert(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't insert `ciudad`",
		})
	}
	serializer.Id = model.Id
	return c.JSON(serializer)
}

func (*Ciudad) Update(c *fiber.Ctx) error {
	serializer := serializers.Ciudad{}
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
	model := models.Ciudad(serializer)
	err = model.Update(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't update ciudad entry",
		})
	}
	return c.JSON(serializer)
}

func (*Ciudad) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	model := models.Ciudad{Id: uint(id)}
	err = model.Delete(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't delete ciudad entry",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}
