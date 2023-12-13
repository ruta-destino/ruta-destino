package handlers

import (
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/serializers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Recorrido struct {
	Service *services.Recorrido
}

func (h *Recorrido) List(c *fiber.Ctx) error {
	idParam := c.Params("id_empresa")
	idEmpresa, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	recorridos, err := h.Service.List(uint(idEmpresa))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't get recorrido entries from db",
		})
	}
	serializer := []serializers.Recorrido{}
	for _, c := range recorridos {
		s := serializers.Recorrido(c)
		serializer = append(serializer, s)
	}
	return c.JSON(serializer)
}

func (h *Recorrido) Insert(c *fiber.Ctx) error {
	serializer := serializers.Recorrido{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	idParam := c.Params("id_empresa")
	idEmpresa, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	model := models.Recorrido(serializer)
	err = h.Service.Insert(uint(idEmpresa), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't insert `recorrido`",
		})
	}
	serializer.Id = model.Id
	serializer.IdEmpresa = uint(idEmpresa)
	return c.JSON(serializer)
}

func (h *Recorrido) Update(c *fiber.Ctx) error {
	serializer := serializers.Recorrido{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	idParam := c.Params("id_empresa")
	idEmpresa, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	idParam = c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	model := models.Recorrido(serializer)
	err = h.Service.Update(uint(idEmpresa), uint(id), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't update recorrido entry",
		})
	}
	serializer.Id = uint(id)
	serializer.IdEmpresa = uint(idEmpresa)
	return c.JSON(serializer)
}

func (h *Recorrido) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id_empresa")
	idEmpresa, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	idParam = c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	err = h.Service.Delete(uint(idEmpresa), uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't delete ciudad entry",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}
