package controller

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/services"
	"github.com/gofiber/fiber/v2"
)

type InventoryController struct {
	inventoryService services.InventoryService
}

func NewInventoryController(service services.InventoryService) *InventoryController {
	return &InventoryController{service}
}

func (c *InventoryController) CreateInventory(ctx *fiber.Ctx) error {
	var inventory entity.Inventory
	if err := ctx.BodyParser(&inventory); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.inventoryService.CreateInventory(&inventory); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(inventory)
}

func (c *InventoryController) GetAllInventories(ctx *fiber.Ctx) error {
	inventories, err := c.inventoryService.GetAllInventories()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(inventories)
}

func (c *InventoryController) GetInventoryByProductID(ctx *fiber.Ctx) error {
	productID := ctx.Params("product_id")
	inventory, err := c.inventoryService.GetInventoryByProductID(productID)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Inventory not found"})
	}
	return ctx.JSON(inventory)
}

func (c *InventoryController) UpdateInventory(ctx *fiber.Ctx) error {
	productID := ctx.Params("product_id")
	var inventory entity.Inventory
	if err := ctx.BodyParser(&inventory); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.inventoryService.UpdateInventory(productID, &inventory); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(inventory)
}

func (c *InventoryController) DeleteInventory(ctx *fiber.Ctx) error {
	productID := ctx.Params("product_id")
	if err := c.inventoryService.DeleteInventory(productID); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
