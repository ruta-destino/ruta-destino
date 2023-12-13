package handlers

import (
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/database/services"
	"ruta-destino/pkg/router/serializers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Empresa struct {
	Service *services.Empresa
}

func (h *Empresa) List(c *fiber.Ctx) error {
	empresas, err := h.Service.List()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't get empresa entries from db",
		})
	}
	serializer := []serializers.Empresa{}
	for _, e := range empresas {
		s := serializers.Empresa(e)
		serializer = append(serializer, s)
	}
	return c.JSON(serializer)
}

func (h *Empresa) Get(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	empresa, err := h.Service.Get(uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't get empresa entry",
		})
	}
	serializer := serializers.Empresa(*empresa)
	return c.JSON(serializer)
}

func (h *Empresa) Insert(c *fiber.Ctx) error {
	serializer := serializers.Empresa{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	model := models.Empresa(serializer)
	err = h.Service.Insert(&model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't insert `empresa`",
		})
	}
	serializer.Id = model.Id
	return c.JSON(serializer)
}

func (h *Empresa) Update(c *fiber.Ctx) error {
	serializer := serializers.Empresa{}
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
	model := models.Empresa(serializer)
	err = h.Service.Update(uint(id), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't update empresa entry",
		})
	}
	serializer.Id = uint(id)
	return c.JSON(serializer)
}

func (h *Empresa) Delete(c *fiber.Ctx) error {
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
			"error": "Couldn't delete empresa entry",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}

func (h *Empresa) ListTerminales(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	terminales, err := h.Service.ListTerminales(uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't get terminal entries from db",
		})
	}
	serializer := []serializers.Terminal{}
	for _, t := range terminales {
		s := serializers.Terminal(t)
		serializer = append(serializer, s)
	}
	return c.JSON(serializer)
}

func (h *Empresa) LinkTerminal(c *fiber.Ctx) error {
	serializer := serializers.EmpresaLinkTerminal{}
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
	err = h.Service.LinkTerminal(uint(id), serializer.Id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't link to terminal",
		})
	}
	return c.JSON(fiber.Map{})
}

func (h *Empresa) UnlinkTerminal(c *fiber.Ctx) error {
	serializer := serializers.EmpresaLinkTerminal{}
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
	err = h.Service.UnlinkTerminal(uint(id), serializer.Id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't unlink from terminal",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}
