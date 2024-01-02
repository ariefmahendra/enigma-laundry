package service

import (
	"context"
	"enigma-laundry/helper"
	"enigma-laundry/internal/model/domain"
	"enigma-laundry/internal/model/dto"
	"enigma-laundry/internal/repository"
	"strconv"
)

type ProductService interface {
	Create(ctx context.Context, request dto.ProductRequest) (dto.ProductResponse, error)
	Update(ctx context.Context, request dto.ProductRequest) (dto.ProductResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (dto.ProductResponse, error)
	FindAll(ctx context.Context) ([]dto.ProductResponse, error)
	FindByName(ctx context.Context, name string) ([]dto.ProductResponse, error)
}

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository: productRepository}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request dto.ProductRequest) (dto.ProductResponse, error) {

	product := domain.Product{
		Name:  request.Name,
		Unit:  request.Unit,
		Price: request.Price,
	}
	productResponse, err := service.productRepository.Insert(ctx, product)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return helper.ProductToResponse(productResponse), nil
}

func (service *ProductServiceImpl) Update(ctx context.Context, request dto.ProductRequest) (dto.ProductResponse, error) {
	productId, err := strconv.Atoi(request.Id)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	product := domain.Product{
		Id:    productId,
		Name:  request.Name,
		Unit:  request.Unit,
		Price: request.Price,
	}

	productResponse, err := service.productRepository.Update(ctx, product)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return helper.ProductToResponse(productResponse), nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := service.productRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = service.productRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductServiceImpl) FindById(ctx context.Context, id int) (dto.ProductResponse, error) {
	product, err := service.productRepository.FindById(ctx, id)
	if err != nil {
		return dto.ProductResponse{}, nil
	}

	return helper.ProductToResponse(product), nil
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := service.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.ProductToResponses(products), nil
}

func (service *ProductServiceImpl) FindByName(ctx context.Context, name string) ([]dto.ProductResponse, error) {
	product, err := service.productRepository.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return helper.ProductToResponses(product), nil
}
