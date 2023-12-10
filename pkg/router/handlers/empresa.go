package handlers

import (
	"ruta-destino/pkg/database"
	"ruta-destino/pkg/database/models"
	"ruta-destino/pkg/router/serializers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Empresa struct{}

func (*Empresa) List(c *fiber.Ctx) error {
	model := models.Empresa{}
	empresas, err := model.List(database.Db)
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

func (*Empresa) Get(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	model := models.Empresa{Id: uint(id)}
	empresa, err := model.Get(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't get empresa entry",
		})
	}
	serializer := serializers.Empresa(*empresa)
	return c.JSON(serializer)
}

func (*Empresa) Insert(c *fiber.Ctx) error {
	serializer := serializers.Empresa{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't parse request body",
		})
	}
	model := models.Empresa(serializer)
	err = model.Insert(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't insert `empresa`",
		})
	}
	serializer.Id = model.Id
	return c.JSON(serializer)
}

func (*Empresa) Update(c *fiber.Ctx) error {
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
	serializer.Id = uint(id)
	model := models.Empresa(serializer)
	err = model.Update(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't update empresa entry",
		})
	}
	return c.JSON(serializer)
}

func (*Empresa) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	model := models.Empresa{Id: uint(id)}
	err = model.Delete(database.Db)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't delete empresa entry",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}

func (*Empresa) ListTerminales(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id, not a number",
		})
	}
	model := models.Empresa{Id: uint(id)}
	terminales, err := model.ListTerminales(database.Db)
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

func (*Empresa) LinkTerminal(c *fiber.Ctx) error {
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
	model := models.Empresa{Id: uint(id)}
	err = model.LinkTerminal(database.Db, serializer.Id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't link to terminal",
		})
	}
	return c.JSON(fiber.Map{})
}

func (*Empresa) UnlinkTerminal(c *fiber.Ctx) error {
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
	model := models.Empresa{Id: uint(id)}
	err = model.UnlinkTerminal(database.Db, serializer.Id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Couldn't unlink from terminal",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}