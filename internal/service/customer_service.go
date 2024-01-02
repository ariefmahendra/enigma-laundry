package service

import (
	"context"
	"enigma-laundry/helper"
	"enigma-laundry/internal/model/domain"
	"enigma-laundry/internal/model/dto"
	"enigma-laundry/internal/repository"
	"strconv"
)

type CustomerService interface {
	Create(ctx context.Context, customerRequest dto.CustomerRequest) (dto.CustomerResponse, error)
	Update(ctx context.Context, customerRequest dto.CustomerUpdateRequest) (dto.CustomerResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, customerId int) (dto.CustomerResponse, error)
	FindAll(ctx context.Context) ([]dto.CustomerResponse, error)
}

type CustomerServiceImpl struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{customerRepository: customerRepository}
}

func (service *CustomerServiceImpl) Create(ctx context.Context, customerRequest dto.CustomerRequest) (dto.CustomerResponse, error) {

	customer := domain.Customer{
		Name:        customerRequest.Name,
		PhoneNumber: customerRequest.PhoneNumber,
		Address:     customerRequest.Address,
	}

	customer, err := service.customerRepository.Insert(ctx, customer)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	return helper.CustomerToResponse(customer), nil
}

func (service *CustomerServiceImpl) Update(ctx context.Context, customerRequest dto.CustomerUpdateRequest) (dto.CustomerResponse, error) {
	customerRequestId, err := strconv.Atoi(customerRequest.Id)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	customer := domain.Customer{
		Id:          customerRequestId,
		Name:        customerRequest.Name,
		PhoneNumber: customerRequest.PhoneNumber,
		Address:     customerRequest.Address,
	}

	customer, err = service.customerRepository.Update(ctx, customer)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	return helper.CustomerToResponse(customer), nil
}

func (service *CustomerServiceImpl) Delete(ctx context.Context, id int) error {
	err := service.customerRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (service *CustomerServiceImpl) FindById(ctx context.Context, customerId int) (dto.CustomerResponse, error) {

	customer, err := service.customerRepository.FindById(ctx, customerId)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	return helper.CustomerToResponse(customer), nil
}

func (service *CustomerServiceImpl) FindAll(ctx context.Context) ([]dto.CustomerResponse, error) {
	customers, err := service.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.CustomerToResponses(customers), nil
}
