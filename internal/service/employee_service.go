package service

import (
	"context"
	"enigma-laundry/helper"
	"enigma-laundry/internal/model/domain"
	"enigma-laundry/internal/model/dto"
	"enigma-laundry/internal/repository"
	"fmt"
	"strconv"
)

type EmployeeService interface {
	Create(ctx context.Context, employeeRequest dto.EmployeeRequest) (dto.EmployeeResponse, error)
	Update(ctx context.Context, employeeRequest dto.EmployeeUpdateRequest) (dto.EmployeeResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (dto.EmployeeResponse, error)
	FindAll(ctx context.Context) ([]dto.EmployeeResponse, error)
}

type EmployeeServiceImpl struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{employeeRepository: employeeRepository}
}

func (service *EmployeeServiceImpl) Create(ctx context.Context, employeeRequest dto.EmployeeRequest) (dto.EmployeeResponse, error) {
	employee := domain.Employee{
		Name:        employeeRequest.Name,
		PhoneNumber: employeeRequest.PhoneNumber,
		Address:     employeeRequest.Address,
	}

	employeeResponse, err := service.employeeRepository.Insert(ctx, employee)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	return helper.EmployeeToResponse(employeeResponse), nil
}

func (service *EmployeeServiceImpl) Update(ctx context.Context, employeeRequest dto.EmployeeUpdateRequest) (dto.EmployeeResponse, error) {
	employeeId, err := strconv.Atoi(employeeRequest.Id)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	employee := domain.Employee{
		Id:          employeeId,
		Name:        employeeRequest.Name,
		PhoneNumber: employeeRequest.PhoneNumber,
		Address:     employeeRequest.Address,
	}

	employeeResponse, err := service.employeeRepository.Update(ctx, employee)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	return helper.EmployeeToResponse(employeeResponse), nil
}

func (service *EmployeeServiceImpl) Delete(ctx context.Context, id int) error {
	fmt.Println(id)
	employee, err := service.employeeRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = service.employeeRepository.Delete(ctx, employee.Id)
	if err != nil {
		return err
	}

	return nil
}

func (service *EmployeeServiceImpl) FindById(ctx context.Context, id int) (dto.EmployeeResponse, error) {
	employee, err := service.employeeRepository.FindById(ctx, id)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	return helper.EmployeeToResponse(employee), nil
}

func (service *EmployeeServiceImpl) FindAll(ctx context.Context) ([]dto.EmployeeResponse, error) {
	employees, err := service.employeeRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return helper.EmployeeToResponses(employees), nil
}
