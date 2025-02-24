package controller

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/services"
	"github.com/gofiber/fiber/v2"
)

type ReceiptController struct {
	receiptService services.ReceiptService
}

func NewReceiptController(service services.ReceiptService) *ReceiptController {
	return &ReceiptController{service}
}

func (c *ReceiptController) CreateReceipt(ctx *fiber.Ctx) error {
	var receipt entity.Receipt
	if err := ctx.BodyParser(&receipt); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.receiptService.CreateReceipt(&receipt); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(receipt)
}

func (c *ReceiptController) GetAllReceipts(ctx *fiber.Ctx) error {
	receipts, err := c.receiptService.GetAllReceipts()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(receipts)
}

func (c *ReceiptController) GetReceiptByID(ctx *fiber.Ctx) error {
	receiptID := ctx.Params("receipt_id")
	receipt, err := c.receiptService.GetReceiptByID(receiptID)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Receipt not found"})
	}
	return ctx.JSON(receipt)
}

func (c *ReceiptController) DeleteReceipt(ctx *fiber.Ctx) error {
	receiptID := ctx.Params("receipt_id")
	if err := c.receiptService.DeleteReceipt(receiptID); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
