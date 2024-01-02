package service

import (
	"context"
	"enigma-laundry/helper"
	"enigma-laundry/internal/model/domain"
	"enigma-laundry/internal/model/dto"
	"enigma-laundry/internal/repository"
	"fmt"
	"strconv"
	"time"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, request dto.CreateBillRequest) (dto.CreateBillResponse, error)
	GetTransactionById(ctx context.Context, id string) (dto.GetBillByIdResponse, error)
	GetAllTransaction(ctx context.Context) ([]dto.GetBillByIdResponse, error)
}

type TransactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
	productRepository     repository.ProductRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository, productRepository repository.ProductRepository) TransactionService {
	return &TransactionServiceImpl{transactionRepository: transactionRepository, productRepository: productRepository}
}

func (service *TransactionServiceImpl) CreateTransaction(ctx context.Context, request dto.CreateBillRequest) (dto.CreateBillResponse, error) {
	tx, err := service.transactionRepository.InitTransaction()
	if err != nil {
		return dto.CreateBillResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	BillDate, err := time.Parse("2006-01-02", request.BillDate)
	if err != nil {
		return dto.CreateBillResponse{}, err
	}

	entryDate, err := time.Parse("2006-01-02", request.EntryDate)
	if err != nil {
		return dto.CreateBillResponse{}, err
	}

	finishDate, err := time.Parse("2006-01-02", request.FinishDate)
	if err != nil {
		return dto.CreateBillResponse{}, err
	}

	employeeId, err := strconv.Atoi(request.EmployeeId)
	if err != nil {
		return dto.CreateBillResponse{}, err
	}

	customerId, err := strconv.Atoi(request.CustomerId)
	if err != nil {
		return dto.CreateBillResponse{}, err
	}

	bill := domain.TxBill{
		BillDate:   BillDate,
		EntryDate:  entryDate,
		FinishDate: finishDate,
		EmployeeId: employeeId,
		CustomerId: customerId,
	}

	var billDetail []domain.TxBillDetail
	totalBill := 0
	for _, item := range request.BillDetail {
		productId, err := strconv.Atoi(item.ProductId)
		if err != nil {
			return dto.CreateBillResponse{}, err
		}

		product, err := service.productRepository.FindById(ctx, productId)
		if err != nil {
			return dto.CreateBillResponse{}, err
		}

		billDetailDomain := domain.TxBillDetail{
			BillId:       bill.Id,
			ProductId:    productId,
			Quantity:     item.Qty,
			ProductPrice: product.Price,
		}

		billDetail = append(billDetail, billDetailDomain)

		totalBill += product.Price * item.Qty
	}

	bill.TotalBill = totalBill

	billResponse, err := service.transactionRepository.InsertBill(ctx, tx, bill)
	if err != nil {
		return dto.CreateBillResponse{}, err
	}

	for i := range billDetail {
		billDetail[i].BillId = billResponse.Id
	}

	billDetails, err := service.transactionRepository.InsertBillDetail(ctx, tx, billDetail)
	if err != nil {
		return dto.CreateBillResponse{}, err
	}

	var billDetailResponse []dto.CreateBillDetailResponse
	for _, item := range billDetails {
		billDetailId := strconv.Itoa(item.Id)
		billDetailBillId := strconv.Itoa(item.BillId)
		billDetailProductId := strconv.Itoa(item.ProductId)

		billDetailResponse = append(billDetailResponse, dto.CreateBillDetailResponse{
			Id:           billDetailId,
			BillId:       billDetailBillId,
			ProductId:    billDetailProductId,
			ProductPrice: item.ProductPrice,
			Quantity:     item.Quantity,
		})
	}

	// convert data type
	billIdStr := strconv.Itoa(billResponse.Id)
	employeeIdStr := strconv.Itoa(billResponse.EmployeeId)
	customerIdStr := strconv.Itoa(billResponse.CustomerId)
	billDateStr := BillDate.Format("2006-01-02")
	entryDateStr := entryDate.Format("2006-01-02")
	finishDateStr := finishDate.Format("2006-01-02")

	response := dto.CreateBillResponse{
		Id:         billIdStr,
		BillDate:   billDateStr,
		EntryDate:  entryDateStr,
		FinishDate: finishDateStr,
		EmployeeId: employeeIdStr,
		CustomerId: customerIdStr,
		BillDetail: billDetailResponse,
	}

	return response, nil
}

func (service *TransactionServiceImpl) GetTransactionById(ctx context.Context, id string) (dto.GetBillByIdResponse, error) {
	tx, err := service.transactionRepository.InitTransaction()
	if err != nil {
		return dto.GetBillByIdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	billResponse, err := service.transactionRepository.GetById(ctx, tx, id)
	if err != nil {
		return dto.GetBillByIdResponse{}, err
	}

	return billResponse, nil
}

func (service *TransactionServiceImpl) GetAllTransaction(ctx context.Context) ([]dto.GetBillByIdResponse, error) {
	tx, err := service.transactionRepository.InitTransaction()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	bills, err := service.transactionRepository.GetAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	var billResponses []dto.GetBillByIdResponse
	for _, bill := range bills {
		var billDetails []dto.GetBillDetailByIdResponse

		for _, billDetail := range bill.BillDetail {
			productId := strconv.Itoa(billDetail.Product.Id)
			product := dto.ProductResponse{
				Id:    productId,
				Name:  billDetail.Product.Name,
				Unit:  billDetail.Product.Unit,
				Price: billDetail.Product.Price,
			}

			billDetail := dto.GetBillDetailByIdResponse{
				Id:           strconv.Itoa(billDetail.Id),
				BillId:       strconv.Itoa(billDetail.BillId),
				ProductPrice: billDetail.ProductPrice,
				Quantity:     billDetail.Quantity,
				Product:      product,
			}

			billDetails = append(billDetails, billDetail)
		}

		billId := strconv.Itoa(bill.Id)
		billDate := bill.BillDate.Format("02-01-2006")
		entryDate := bill.EntryDate.Format("02-01-2006")
		finishDate := bill.FinishDate.Format("02-01-2006")
		employee := dto.EmployeeResponse{
			Id:          strconv.Itoa(bill.Employee.Id),
			Name:        bill.Employee.Name,
			PhoneNumber: bill.Employee.PhoneNumber,
			Address:     bill.Employee.Address,
		}

		customer := dto.CustomerResponse{
			Id:          strconv.Itoa(bill.Customer.Id),
			Name:        bill.Customer.Name,
			PhoneNumber: bill.Customer.PhoneNumber,
			Address:     bill.Customer.Address,
		}

		billResponse := dto.GetBillByIdResponse{
			Id:         billId,
			BillDate:   billDate,
			EntryDate:  entryDate,
			FinishDate: finishDate,
			Employee:   employee,
			Customer:   customer,
			BillDetail: billDetails,
			TotalBill:  bill.TotalBill,
		}

		billResponses = append(billResponses, billResponse)
	}

	fmt.Println("billResponse :", billResponses)

	return billResponses, nil
}
