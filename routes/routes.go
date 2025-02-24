package routes

import (
	"github.com/djohanmirza/go-struct/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, productController *controller.ProductController,
	customerController *controller.CustomerController,
	employeeController *controller.EmployeeController,
	inventoryController *controller.InventoryController,
	orderController *controller.OrderController,
	paymentController *controller.PaymentController,
	receiptController *controller.ReceiptController,
	supplierController *controller.SupplierController,
	taxController *controller.TaxController) {
	// Product Routes
	productGroup := app.Group("/products")
	productGroup.Post("/", productController.CreateProduct)
	productGroup.Get("/", productController.GetAllProducts)
	productGroup.Get("/:id", productController.GetProductByID)
	productGroup.Put("/:id", productController.UpdateProduct)
	productGroup.Delete("/:id", productController.DeleteProduct)

	// Customer Routes
	customerGroup := app.Group("/customers")
	customerGroup.Post("/", customerController.CreateCustomer)
	customerGroup.Get("/", customerController.GetAllCustomers)
	customerGroup.Get("/:id", customerController.GetCustomerByID)
	customerGroup.Put("/:id", customerController.UpdateCustomer)
	customerGroup.Delete("/:id", customerController.DeleteCustomer)

	// Employee Routes
	employeeGroup := app.Group("/employees")
	employeeGroup.Post("/", employeeController.CreateEmployee)
	employeeGroup.Get("/", employeeController.GetAllEmployees)
	employeeGroup.Get("/:id", employeeController.GetEmployeeByID)
	employeeGroup.Put("/:id", employeeController.UpdateEmployee)
	employeeGroup.Delete("/:id", employeeController.DeleteEmployee)

	// Inventory Routes
	inventoryGroup := app.Group("/inventory")
	inventoryGroup.Post("/", inventoryController.CreateInventory)
	inventoryGroup.Put("/:id", inventoryController.UpdateInventory)
	inventoryGroup.Delete("/:id", inventoryController.DeleteInventory)

	// Order Routes
	orderGroup := app.Group("/orders")
	orderGroup.Post("/", orderController.CreateOrder)
	orderGroup.Get("/", orderController.GetAllOrders)
	orderGroup.Get("/:id", orderController.GetOrderByID)
	orderGroup.Delete("/:id", orderController.DeleteOrder)

	// Payment Routes
	paymentGroup := app.Group("/payments")
	paymentGroup.Post("/", paymentController.CreatePayment)
	paymentGroup.Get("/", paymentController.GetAllPayments)
	paymentGroup.Get("/:id", paymentController.GetPaymentByID)
	paymentGroup.Delete("/:id", paymentController.DeletePayment)

	// Receipt Routes
	receiptGroup := app.Group("/receipts")
	receiptGroup.Post("/", receiptController.CreateReceipt)
	receiptGroup.Get("/", receiptController.GetAllReceipts)
	receiptGroup.Get("/:id", receiptController.GetReceiptByID)
	receiptGroup.Delete("/:id", receiptController.DeleteReceipt)

	// Supplier Routes
	supplierGroup := app.Group("/suppliers")
	supplierGroup.Post("/", supplierController.CreateSupplier)
	supplierGroup.Get("/", supplierController.GetAllSuppliers)
	supplierGroup.Get("/:id", supplierController.GetSupplierByID)
	supplierGroup.Put("/:id", supplierController.UpdateSupplier)
	supplierGroup.Delete("/:id", supplierController.DeleteSupplier)

	// Tax Routes
	taxGroup := app.Group("/taxes")
	taxGroup.Post("/", taxController.CreateTax)
	taxGroup.Get("/", taxController.GetAllTaxes)
	taxGroup.Get("/:id", taxController.GetTaxByID)
	taxGroup.Put("/:id", taxController.UpdateTax)
	taxGroup.Delete("/:id", taxController.DeleteTax)
}
