package controller

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/services"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{service}
}

func (c *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	var category entity.Category
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.categoryService.CreateCategory(&category); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(category)
}

func (c *CategoryController) GetAllCategories(ctx *fiber.Ctx) error {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(categories)
}

func (c *CategoryController) GetCategoryByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	category, err := c.categoryService.GetCategoryByID(uint64(id))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Category not found"})
	}
	return ctx.JSON(category)
}

func (c *CategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	var category entity.Category
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.categoryService.UpdateCategory(uint64(id), &category); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(category)
}

func (c *CategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	if err := c.categoryService.DeleteCategory(uint64(id)); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
