package handlers

import (
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/serializers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Ciudad struct {
	Service *services.Ciudad
}

func (h *Ciudad) List(c *fiber.Ctx) error {
	ciudades, err := h.Service.List()
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

func (h *Ciudad) Insert(c *fiber.Ctx) error {
	serializer := serializers.Ciudad{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	model := models.Ciudad(serializer)
	err = h.Service.Insert(&model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't insert `ciudad`",
		})
	}
	serializer.Id = model.Id
	return c.JSON(serializer)
}

func (h *Ciudad) Update(c *fiber.Ctx) error {
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
	model := models.Ciudad(serializer)
	err = h.Service.Update(uint(id), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't update ciudad entry",
		})
	}
	serializer.Id = uint(id)
	return c.JSON(serializer)
}

func (h *Ciudad) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	err = h.Service.Delete(uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't delete ciudad entry",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}
