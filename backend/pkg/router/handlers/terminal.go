package handlers

import (
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/serializers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Terminal struct {
	Service *services.Terminal
}

func (h *Terminal) List(c *fiber.Ctx) error {
	terminales, err := h.Service.List()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible obtener los terminales",
		})
	}
	serializer := []serializers.Terminal{}
	for _, t := range terminales {
		s := serializers.Terminal(t)
		serializer = append(serializer, s)
	}
	return c.JSON(serializer)
}

func (h *Terminal) Get(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de terminal es un entero positivo",
		})
	}
	terminal, err := h.Service.Get(uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible obtener el terminal",
		})
	}
	serializer := serializers.Terminal(*terminal)
	return c.JSON(serializer)
}

func (h *Terminal) Insert(c *fiber.Ctx) error {
	serializer := serializers.Terminal{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible procesar el cuerpo de petición",
		})
	}
	model := models.Terminal(serializer)
	err = h.Service.Insert(&model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	serializer.Id = model.Id
	return c.JSON(serializer)
}

func (h *Terminal) Update(c *fiber.Ctx) error {
	serializer := serializers.Terminal{}
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
			"error": "Id de terminal es un entero positivo",
		})
	}
	model := models.Terminal(serializer)
	err = h.Service.Update(uint(id), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	serializer.Id = uint(id)
	return c.JSON(serializer)
}

func (h *Terminal) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de terminal es un entero positivo",
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

func (h *Terminal) ListEmpresa(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de terminal es un entero positivo",
		})
	}
	empresas, err := h.Service.ListEmpresa(uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible obtener las empresas",
		})
	}
	serializer := []serializers.Empresa{}
	for _, e := range empresas {
		s := serializers.Empresa(e)
		serializer = append(serializer, s)
	}
	return c.JSON(serializer)
}
