package helper

import (
	"enigma-laundry/internal/model/domain"
	"enigma-laundry/internal/model/dto"
	"strconv"
)

func CustomerToResponse(customer domain.Customer) dto.CustomerResponse {
	customerId := strconv.Itoa(customer.Id)
	return dto.CustomerResponse{
		Id:          customerId,
		Name:        customer.Name,
		PhoneNumber: customer.PhoneNumber,
		Address:     customer.Address,
	}
}

func CustomerToResponses(customers []domain.Customer) []dto.CustomerResponse {

	var customerResponses []dto.CustomerResponse
	for _, customer := range customers {
		customerResponse := CustomerToResponse(customer)
		customerResponses = append(customerResponses, customerResponse)
	}

	return customerResponses
}

func EmployeeToResponse(employee domain.Employee) dto.EmployeeResponse {
	employeeId := strconv.Itoa(employee.Id)
	return dto.EmployeeResponse{
		Id:          employeeId,
		Name:        employee.Name,
		PhoneNumber: employee.PhoneNumber,
		Address:     employee.Address,
	}
}

func EmployeeToResponses(employees []domain.Employee) []dto.EmployeeResponse {
	var employeeResponses []dto.EmployeeResponse
	for _, employee := range employees {
		employeeResponse := EmployeeToResponse(employee)
		employeeResponses = append(employeeResponses, employeeResponse)
	}
	return employeeResponses
}

func ProductToResponse(product domain.Product) dto.ProductResponse {
	productId := strconv.Itoa(product.Id)
	return dto.ProductResponse{
		Id:    productId,
		Name:  product.Name,
		Unit:  product.Unit,
		Price: product.Price,
	}

}

func ProductToResponses(products []domain.Product) []dto.ProductResponse {
	var productResponses []dto.ProductResponse
	for _, product := range products {
		productResponse := ProductToResponse(product)
		productResponses = append(productResponses, productResponse)
	}
	return productResponses
}
