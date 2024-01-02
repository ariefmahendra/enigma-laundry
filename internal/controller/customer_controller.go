package controller

import (
	"enigma-laundry/helper"
	"enigma-laundry/internal/model/dto"
	"enigma-laundry/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CustomerController interface {
	CreateCustomer(ctx *gin.Context)
	UpdateCustomer(ctx *gin.Context)
	DeleteCustomer(ctx *gin.Context)
	GetCustomerById(ctx *gin.Context)
	GetCustomers(ctx *gin.Context)
}

type CustomerControllerImpl struct {
	customerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{customerService: customerService}
}

func (controller *CustomerControllerImpl) CreateCustomer(ctx *gin.Context) {
	var customerCreated dto.CustomerRequest

	err := ctx.BindJSON(&customerCreated)
	helper.PanicIfError(err)

	customerResponse, err := controller.customerService.Create(ctx, customerCreated)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *CustomerControllerImpl) UpdateCustomer(ctx *gin.Context) {
	var customerUpdated dto.CustomerUpdateRequest

	customerId := ctx.Param("id")

	customerUpdated.Id = customerId

	err := ctx.BindJSON(&customerUpdated)
	helper.PanicIfError(err)

	customerResponse, err := controller.customerService.Update(ctx, customerUpdated)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *CustomerControllerImpl) DeleteCustomer(ctx *gin.Context) {
	customerId := ctx.Param("id")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		log.Fatal(err)
	}

	err = controller.customerService.Delete(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *CustomerControllerImpl) GetCustomerById(ctx *gin.Context) {
	customerId := ctx.Param("id")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		log.Fatal(err)
	}

	customerResponse, _ := controller.customerService.FindById(ctx, id)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *CustomerControllerImpl) GetCustomers(ctx *gin.Context) {
	customerResponses, err := controller.customerService.FindAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customerResponses,
	}

	ctx.JSON(http.StatusOK, response)
}
