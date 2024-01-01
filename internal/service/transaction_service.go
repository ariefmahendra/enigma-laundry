package service

import (
	"context"
	"database/sql"
	"enigma-laundry/helper"
	"enigma-laundry/internal/model/domain"
	"enigma-laundry/internal/model/dto"
	"enigma-laundry/internal/repository"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, request dto.CreateBillRequest) (dto.CreateBillResponse, error)
	GetTransactionById(ctx context.Context, id string) (dto.GetBillByIdResponse, error)
	GetAllTransaction(ctx context.Context) ([]dto.GetBillByIdResponse, error)
}

type TransactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
	productRepository     repository.ProductRepository
	db                    *sql.DB
}

func NewTransactionService(transactionRepository repository.TransactionRepository, productRepository repository.ProductRepository, db *sql.DB) TransactionService {
	return &TransactionServiceImpl{transactionRepository: transactionRepository, productRepository: productRepository, db: db}
}

func (service *TransactionServiceImpl) CreateTransaction(ctx context.Context, request dto.CreateBillRequest) (dto.CreateBillResponse, error) {
	tx, err := service.db.Begin()
	if err != nil {
		return dto.CreateBillResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	bill := domain.TxBill{
		BillDate:   request.BillDate,
		EntryDate:  request.EntryDate,
		FinishDate: request.FinishDate,
		EmployeeId: request.EmployeeId,
		CustomerId: request.CustomerId,
	}

	var billDetail []domain.TxBillDetail
	totalBill := 0
	for _, item := range request.BillDetail {
		product, err := service.productRepository.FindById(ctx, item.ProductId)
		if err != nil {
			return dto.CreateBillResponse{}, err
		}

		billDetailDomain := domain.TxBillDetail{
			BillId:       bill.Id,
			ProductId:    item.ProductId,
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
		billDetailResponse = append(billDetailResponse, dto.CreateBillDetailResponse{
			Id:           item.Id,
			BillId:       item.BillId,
			ProductId:    item.ProductId,
			ProductPrice: item.ProductPrice,
			Quantity:     item.Quantity,
		})
	}

	response := dto.CreateBillResponse{
		Id:         billResponse.Id,
		BillDate:   billResponse.BillDate,
		EntryDate:  billResponse.EntryDate,
		FinishDate: billResponse.FinishDate,
		Employee:   billResponse.EmployeeId,
		Customer:   billResponse.CustomerId,
		BillDetail: billDetailResponse,
		TotalBill:  billResponse.TotalBill,
	}

	return response, nil
}

func (service *TransactionServiceImpl) GetTransactionById(ctx context.Context, id string) (dto.GetBillByIdResponse, error) {
	tx, err := service.db.Begin()
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
	//"TODO implement me"
	panic("implement me")
}
