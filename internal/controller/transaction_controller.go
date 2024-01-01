package controller

import (
	"enigma-laundry/helper"
	"enigma-laundry/internal/model/dto"
	"enigma-laundry/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionController interface {
	CreateTransaction(ctx *gin.Context)
	GetTransactionById(ctx *gin.Context)
	GetAllTransaction(ctx *gin.Context)
}

type TransactionControllerImpl struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{transactionService: transactionService}
}

func (controller TransactionControllerImpl) CreateTransaction(ctx *gin.Context) {
	var transactionRequest dto.CreateBillRequest
	err := ctx.BindJSON(&transactionRequest)
	helper.PanicIfError(err)

	transaction, err := controller.transactionService.CreateTransaction(ctx, transactionRequest)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   transaction,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller TransactionControllerImpl) GetTransactionById(ctx *gin.Context) {
	id := ctx.Param("id")

	transaction, err := controller.transactionService.GetTransactionById(ctx, id)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   transaction,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller TransactionControllerImpl) GetAllTransaction(ctx *gin.Context) {
	billsResponse, err := controller.transactionService.GetAllTransaction(ctx)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   billsResponse,
	}

	ctx.JSON(http.StatusOK, response)
}
