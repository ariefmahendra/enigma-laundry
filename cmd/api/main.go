package main

import (
	"enigma-laundry/helper"
	"enigma-laundry/internal/config"
	"enigma-laundry/internal/controller"
	"enigma-laundry/internal/repository"
	"enigma-laundry/internal/router"
	"enigma-laundry/internal/service"
	"net/http"
)

func main() {
	initDB := config.InitDB()
	defer func() {
		err := initDB.Close()
		helper.PanicIfError(err)
	}()

	// initialize customer
	customerRepository := repository.NewCustomerRepository(initDB)
	customerService := service.NewCustomerService(customerRepository)
	CustomerController := controller.NewCustomerController(customerService)

	// initialize employee
	employeeRepository := repository.NewEmployeeRepository(initDB)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controller.NewEmployeeController(employeeService)

	// initialize product
	productRepository := repository.NewProductRepository(initDB)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	// initialize transaction
	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository, productRepository, initDB)
	transactionController := controller.NewTransactionController(transactionService)

	handler := router.NewRouter(CustomerController, employeeController, productController, transactionController)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
