package controller

import (
	"enigma-laundry/helper"
	"enigma-laundry/internal/model/dto"
	"enigma-laundry/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController interface {
	CreateEmployee(ctx *gin.Context)
	UpdateEmployee(ctx *gin.Context)
	DeleteEmployee(ctx *gin.Context)
	GetEmployeeById(ctx *gin.Context)
	GetEmployees(ctx *gin.Context)
}

type EmployeeControllerImpl struct {
	employeeService service.EmployeeService
}

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{employeeService: employeeService}
}

func (controller *EmployeeControllerImpl) CreateEmployee(ctx *gin.Context) {
	var employeeCreated dto.EmployeeRequest
	err := ctx.BindJSON(&employeeCreated)
	helper.PanicIfError(err)

	employee, err := controller.employeeService.Create(ctx, employeeCreated)
	helper.PanicIfError(err)
	fmt.Println(employee)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   employee,
	}

	ctx.JSON(http.StatusOK, response)

}

func (controller *EmployeeControllerImpl) UpdateEmployee(ctx *gin.Context) {
	var (
		employeeUpdated dto.EmployeeUpdateRequest
		id              int
	)

	idEmployee := ctx.Param("id")
	id, err := strconv.Atoi(idEmployee)
	helper.PanicIfError(err)

	err = ctx.BindJSON(&employeeUpdated)
	helper.PanicIfError(err)

	employeeUpdated.Id = id

	employeeResponse, err := controller.employeeService.Update(ctx, employeeUpdated)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   employeeResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *EmployeeControllerImpl) DeleteEmployee(ctx *gin.Context) {
	idEmployee := ctx.Param("id")

	id, err := strconv.Atoi(idEmployee)
	helper.PanicIfError(err)

	err = controller.employeeService.Delete(ctx, id)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, response)

}

func (controller *EmployeeControllerImpl) GetEmployeeById(ctx *gin.Context) {
	idEmployee := ctx.Param("id")
	id, err := strconv.Atoi(idEmployee)
	helper.PanicIfError(err)

	employeeResponse, err := controller.employeeService.FindById(ctx, id)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   employeeResponse,
	}

	ctx.JSON(http.StatusOK, response)

}

func (controller *EmployeeControllerImpl) GetEmployees(ctx *gin.Context) {
	employees, err := controller.employeeService.FindAll(ctx)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   employees,
	}

	ctx.JSON(http.StatusOK, response)
}
