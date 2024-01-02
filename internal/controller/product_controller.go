package controller

import (
	"enigma-laundry/helper"
	"enigma-laundry/internal/model/dto"
	"enigma-laundry/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type ProductControllerImpl struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{productService: productService}
}

func (controller *ProductControllerImpl) CreateProduct(ctx *gin.Context) {
	var productCreated dto.ProductRequest
	err := ctx.BindJSON(&productCreated)
	helper.PanicIfError(err)

	productResponse, err := controller.productService.Create(ctx, productCreated)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *ProductControllerImpl) UpdateProduct(ctx *gin.Context) {
	var productUpdate dto.ProductRequest

	productId := ctx.Param("id")

	productUpdate.Id = productId

	err := ctx.BindJSON(&productUpdate)
	helper.PanicIfError(err)

	productResponse, err := controller.productService.Update(ctx, productUpdate)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *ProductControllerImpl) DeleteProduct(ctx *gin.Context) {
	idProduct := ctx.Param("id")
	id, err := strconv.Atoi(idProduct)
	helper.PanicIfError(err)

	err = controller.productService.Delete(ctx, id)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *ProductControllerImpl) GetById(ctx *gin.Context) {
	idProduct := ctx.Param("id")
	id, err := strconv.Atoi(idProduct)
	helper.PanicIfError(err)

	productResponse, err := controller.productService.FindById(ctx, id)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *ProductControllerImpl) GetAll(ctx *gin.Context) {
	name := ctx.Query("name")
	if name != "" {
		productResponse, err := controller.productService.FindByName(ctx, name)
		helper.PanicIfError(err)

		response := dto.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   productResponse,
		}

		ctx.JSON(http.StatusOK, response)
		return
	}

	productResponse, err := controller.productService.FindAll(ctx)
	helper.PanicIfError(err)

	response := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	ctx.JSON(http.StatusOK, response)
}
