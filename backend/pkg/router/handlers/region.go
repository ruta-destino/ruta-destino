package handlers

import (
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/serializers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Region struct {
	Service *services.Region
}

func (h *Region) List(c *fiber.Ctx) error {
	regiones, err := h.Service.List()
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

func (h *Region) Get(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	model, err := h.Service.Get(uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't get region entry",
		})
	}
	serializer := serializers.Region(*model)
	return c.JSON(serializer)
}

func (h *Region) Insert(c *fiber.Ctx) error {
	serializer := serializers.Region{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	model := models.Region(serializer)
	err = h.Service.Insert(&model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't insert `region`",
		})
	}
	serializer.Id = model.Id
	return c.JSON(serializer)
}

func (h *Region) Update(c *fiber.Ctx) error {
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
	model := models.Region(serializer)
	err = h.Service.Update(uint(id), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't update region entry",
		})
	}
	serializer.Id = uint(id)
	return c.JSON(serializer)
}

func (h *Region) Delete(c *fiber.Ctx) error {
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
			"error": "Couldn't delete region entry",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}
