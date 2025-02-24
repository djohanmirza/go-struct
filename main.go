package main

import (
	"fmt"
	"github.com/djohanmirza/go-struct/controller"
	"github.com/djohanmirza/go-struct/database"
	"github.com/djohanmirza/go-struct/repository"
	"github.com/djohanmirza/go-struct/routes"
	"github.com/djohanmirza/go-struct/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDatabase()

	//Dependency Injection
	productRepo := repository.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	//Customer
	customerRepo := repository.NewCustomerRepository(database.DB)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controller.NewCustomerController(customerService)

	//Employee
	employeeRepo := repository.NewEmployeeRepository(database.DB)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	//Inventory
	inventoryRepo := repository.NewInventoryRepository(database.DB)
	inventoryService := services.NewInventoryService(inventoryRepo)
	inventoryController := controller.NewInventoryController(inventoryService)

	//Order
	orderRepo := repository.NewOrderRepository(database.DB)
	orderService := services.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService)

	//Payment
	paymentRepo := repository.NewPaymentRepository(database.DB)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentController := controller.NewPaymentController(paymentService)

	//Receipt
	receiptRepo := repository.NewReceiptRepository(database.DB)
	receiptService := services.NewReceiptService(receiptRepo)
	receiptController := controller.NewReceiptController(receiptService)

	//Supplier
	supplierRepo := repository.NewSupplierRepository(database.DB)
	supplierService := services.NewSupplierService(supplierRepo)
	supplierController := controller.NewSupplierController(supplierService)

	//Tax
	taxRepo := repository.NewTaxRepository(database.DB)
	taxService := services.NewTaxService(taxRepo)
	taxController := controller.NewTaxController(taxService)

	//Setup Routes
	routes.SetupRoutes(app, productController, customerController, employeeController, inventoryController, orderController, paymentController, receiptController, supplierController, taxController)

	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
