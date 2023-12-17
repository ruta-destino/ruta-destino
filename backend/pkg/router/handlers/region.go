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
			"error": "No fue posible obtener las regiones",
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
			"error": "Id de región es un entero positivo",
		})
	}
	model, err := h.Service.Get(uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible obtener la region",
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
			"error": "No fue posible procesar el cuerpo de petición",
		})
	}
	model := models.Region(serializer)
	err = h.Service.Insert(&model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
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
			"error": "No fue posible procesar el cuerpo de petición",
		})
	}
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de región es un entero positivo",
		})
	}
	model := models.Region(serializer)
	err = h.Service.Update(uint(id), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
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
			"error": "Id de región es un entero positivo",
		})
	}
	err = h.Service.Delete(uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}
