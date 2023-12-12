package handlers

import (
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/serializers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Provincia struct {
	Service *services.Provincia
}

func (h *Provincia) List(c *fiber.Ctx) error {
	provincias, err := h.Service.List()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't get provincia entries from db",
		})
	}
	serializer := []serializers.Provincia{}
	for _, p := range provincias {
		s := serializers.Provincia(p)
		serializer = append(serializer, s)
	}
	return c.JSON(serializer)
}

func (h *Provincia) Insert(c *fiber.Ctx) error {
	serializer := serializers.Provincia{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	model := models.Provincia(serializer)
	err = h.Service.Insert(&model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't insert `provincia`",
		})
	}
	serializer.Id = model.Id
	return c.JSON(serializer)
}

func (h *Provincia) Update(c *fiber.Ctx) error {
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
	model := models.Provincia(serializer)
	err = h.Service.Update(uint(id), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't update provincia entry",
		})
	}
	serializer.Id = uint(id)
	return c.JSON(serializer)
}

func (h *Provincia) Delete(c *fiber.Ctx) error {
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
			"error": "Couldn't delete provincia entry",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}
