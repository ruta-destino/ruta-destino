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
			"error": "Id de empresa es un entero positivo",
		})
	}
	recorridos, err := h.Service.List(uint(idEmpresa))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible obtener los recorridos",
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
			"error": "No fue posible procesar el cuerpo de petición",
		})
	}
	idParam := c.Params("id_empresa")
	idEmpresa, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de empresa es un entero positivo",
		})
	}
	model := models.Recorrido(serializer)

	// TODO: Esto lo hago por faltas de tiempo.
	// Comprobar que los terminales estén vinculados con la empresa, sino
	// devolver un error. Esto debería trabajar desde base de datos porque aquí
	// puede haber cambios entre la comprobación y agregar el recorrido, lo
	// correcto es usar una transacción en este caso.
	empresaService := services.NewEmpresaService(h.Service.Db)
	terminales, err := empresaService.ListTerminales(uint(idEmpresa))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible obtener los terminales",
		})
	}
	origen := false
	for _, t := range terminales {
		if model.IdTerminalOrigen == t.Id {
			origen = true
		}
	}
	if !origen {
		return c.Status(400).JSON(fiber.Map{
			"error": "La empresa no trabaja con el terminal de origen",
		})
	}
	destino := false
	for _, t := range terminales {
		if model.IdTerminalDestino == t.Id {
			destino = true
		}
	}
	if !destino {
		return c.Status(400).JSON(fiber.Map{
			"error": "La empresa no trabaja con el terminal de destino",
		})
	}

	err = h.Service.Insert(uint(idEmpresa), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	serializer.Id = model.Id
	serializer.IdEmpresa = uint(idEmpresa)
	return c.JSON(serializer)
}

func (h *Recorrido) Get(c *fiber.Ctx) error {
	idParam := c.Params("id_empresa")
	idEmpresa, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de empresa es un entero positivo",
		})
	}
	idParam = c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de recorrido es un entero positivo",
		})
	}
	model, err := h.Service.Get(uint(idEmpresa), uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible obtener el recorrido",
		})
	}
	serializer := serializers.Recorrido(*model)
	return c.JSON(serializer)
}

func (h *Recorrido) Update(c *fiber.Ctx) error {
	serializer := serializers.Recorrido{}
	err := c.BodyParser(&serializer)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible procesar el cuerpo de petición",
		})
	}
	idParam := c.Params("id_empresa")
	idEmpresa, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de empresa es un entero positivo",
		})
	}
	idParam = c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de recorrido es un entero positivo",
		})
	}
	model := models.Recorrido(serializer)

	// TODO: Esto lo hago por faltas de tiempo.
	// Comprobar que los terminales estén vinculados con la empresa, sino
	// devolver un error. Esto debería trabajar desde base de datos porque aquí
	// puede haber cambios entre la comprobación y agregar el recorrido, lo
	// correcto es usar una transacción en este caso.
	empresaService := services.NewEmpresaService(h.Service.Db)
	terminales, err := empresaService.ListTerminales(uint(idEmpresa))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No fue posible obtener los terminales",
		})
	}
	origen := false
	for _, t := range terminales {
		if model.IdTerminalOrigen == t.Id {
			origen = true
		}
	}
	if !origen {
		return c.Status(400).JSON(fiber.Map{
			"error": "La empresa no trabaja con el terminal de origen",
		})
	}
	destino := false
	for _, t := range terminales {
		if model.IdTerminalDestino == t.Id {
			destino = true
		}
	}
	if !destino {
		return c.Status(400).JSON(fiber.Map{
			"error": "La empresa no trabaja con el terminal de destino",
		})
	}

	err = h.Service.Update(uint(idEmpresa), uint(id), &model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
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
			"error": "Id de empresa es un entero positivo",
		})
	}
	idParam = c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Id de recorrido es un entero positivo",
		})
	}
	err = h.Service.Delete(uint(idEmpresa), uint(id))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}
