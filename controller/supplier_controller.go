package controller

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type SupplierController struct {
	supplierService services.SupplierService
}

func NewSupplierController(service services.SupplierService) *SupplierController {
	return &SupplierController{service}
}

func (c *SupplierController) CreateSupplier(ctx *fiber.Ctx) error {
	var supplier entity.Supplier
	if err := ctx.BodyParser(&supplier); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.supplierService.CreateSupplier(&supplier); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(supplier)
}

func (c *SupplierController) GetAllSuppliers(ctx *fiber.Ctx) error {
	suppliers, err := c.supplierService.GetAllSuppliers()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(suppliers)
}

func (c *SupplierController) GetSupplierByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	supplier, err := c.supplierService.GetSupplierByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Supplier not found"})
	}
	return ctx.JSON(supplier)
}

func (c *SupplierController) UpdateSupplier(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	var supplier entity.Supplier
	if err := ctx.BodyParser(&supplier); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.supplierService.UpdateSupplier(id, &supplier); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(supplier)
}

func (c *SupplierController) DeleteSupplier(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	if err := c.supplierService.DeleteSupplier(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
