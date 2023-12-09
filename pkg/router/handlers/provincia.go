package handlers

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/router/serializers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Provincia struct{}

func (p *Provincia) List(c *fiber.Ctx) error {
	model := models.Provincia{}
	provincias, err := model.List(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't get provincia entries from db",
		})
	}
	serializer := []serializers.Provincia{}
	for _, r := range provincias {
		s := serializers.Provincia(r)
		serializer = append(serializer, s)
	}
	return c.JSON(serializer)
}

func (p *Provincia) Insert(c *fiber.Ctx) error {
	serializer := serializers.Provincia{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	model := models.Provincia(serializer)
	err = model.Insert(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't insert `provincia`",
		})
	}
	return c.JSON(serializer)
}

func (p *Provincia) Update(c *fiber.Ctx) error {
	serializer := serializers.Provincia{}
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
	model := models.Provincia(serializer)
	err = model.Update(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't update provincia entry",
		})
	}
	return c.JSON(serializer)
}

func (p *Provincia) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	model := models.Provincia{Id: uint(id)}
	err = model.Delete(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't delete provincia entry",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}
