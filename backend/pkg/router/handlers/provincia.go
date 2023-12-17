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
			"error": "No fue posible obtener las provincias",
		})
	}
	serializer := []serializers.Provincia{}
	for _, p := range provincias {
		s := serializers.Provincia(p)
		serializer = append(serializer, s)
	}
	return c.JSON(serializer)
}

func (h *Provincia) Get(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de provincia es un entero positivo",
		})
	}
	model, err := h.Service.Get(uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible obtener la provincia",
		})
	}
	serializer := serializers.Provincia(*model)
	return c.JSON(serializer)
}

func (h *Provincia) Insert(c *fiber.Ctx) error {
	serializer := serializers.Provincia{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible procesar el cuerpo de petición",
		})
	}
	model := models.Provincia(serializer)
	err = h.Service.Insert(&model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
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
			"error": "No fue posible procesar el cuerpo de petición",
		})
	}
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de provincia es un entero positivo",
		})
	}
	model := models.Provincia(serializer)
	err = h.Service.Update(uint(id), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
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
			"error": "Id de provincia es un entero positivo",
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
