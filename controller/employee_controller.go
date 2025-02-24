package controller

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/services"
	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	employeeService services.EmployeeService
}

func NewEmployeeController(service services.EmployeeService) *EmployeeController {
	return &EmployeeController{service}
}

func (c *EmployeeController) CreateEmployee(ctx *fiber.Ctx) error {
	var employee entity.Employee
	if err := ctx.BodyParser(&employee); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.employeeService.CreateEmployee(&employee); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(employee)
}

func (c *EmployeeController) GetAllEmployees(ctx *fiber.Ctx) error {
	employees, err := c.employeeService.GetAllEmployees()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(employees)
}

func (c *EmployeeController) GetEmployeeByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	employee, err := c.employeeService.GetEmployeeByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}
	return ctx.JSON(employee)
}

func (c *EmployeeController) UpdateEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var employee entity.Employee
	if err := ctx.BodyParser(&employee); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.employeeService.UpdateEmployee(id, &employee); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(employee)
}

func (c *EmployeeController) DeleteEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.employeeService.DeleteEmployee(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
