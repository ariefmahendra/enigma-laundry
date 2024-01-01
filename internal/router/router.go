package router

import (
	"enigma-laundry/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(customerController controller.CustomerController, employeeController controller.EmployeeController, productController controller.ProductController, transactionController controller.TransactionController) *gin.Engine {
	r := gin.Default()

	// management customer
	r.POST("/customers", customerController.CreateCustomer)
	r.PUT("/customers/:id", customerController.UpdateCustomer)
	r.DELETE("/customers/:id", customerController.DeleteCustomer)
	r.GET("/customers/:id", customerController.GetCustomerById)
	r.GET("/customers", customerController.GetCustomers)

	// management employee
	r.POST("/employees", employeeController.CreateEmployee)
	r.PUT("/employees/:id", employeeController.UpdateEmployee)
	r.DELETE("/employees/:id", employeeController.DeleteEmployee)
	r.GET("/employees/:id", employeeController.GetEmployeeById)
	r.GET("/employees", employeeController.GetEmployees)

	// management product
	r.POST("/products", productController.CreateProduct)
	r.PUT("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)
	r.GET("/products/:id", productController.GetById)
	r.GET("/products", productController.GetAll)

	// transaction
	r.POST("/transactions", transactionController.CreateTransaction)
	r.GET("/transactions/:id", transactionController.GetTransactionById)
	r.GET("/transactions", transactionController.GetAllTransaction)

	return r
}
