package controller

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ShipController struct {
	shipService services.ShipService
}

func NewShipController(service services.ShipService) *ShipController {
	return &ShipController{service}
}

func (c *ShipController) CreateShip(ctx *fiber.Ctx) error {
	var ship entity.Ship
	if err := ctx.BodyParser(&ship); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.shipService.CreateShip(&ship); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(ship)
}

func (c *ShipController) GetAllShips(ctx *fiber.Ctx) error {
	ships, err := c.shipService.GetAllShips()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(ships)
}

func (c *ShipController) GetShipByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	ship, err := c.shipService.GetShipByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Ship not found"})
	}
	return ctx.JSON(ship)
}

func (c *ShipController) UpdateShip(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	var ship entity.Ship
	if err := ctx.BodyParser(&ship); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.shipService.UpdateShip(id, &ship); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(ship)
}

func (c *ShipController) DeleteShip(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	if err := c.shipService.DeleteShip(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
