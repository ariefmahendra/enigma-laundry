package repository

import (
	"context"
	"database/sql"
	"enigma-laundry/internal/model/domain"
	"enigma-laundry/internal/model/dto"
	"fmt"
	"strconv"
)

type TransactionRepository interface {
	InitTransaction() (*sql.Tx, error)
	InsertBill(ctx context.Context, tx *sql.Tx, bill domain.TxBill) (domain.TxBill, error)
	InsertBillDetail(ctx context.Context, tx *sql.Tx, billDetail []domain.TxBillDetail) ([]domain.TxBillDetail, error)
	GetById(ctx context.Context, tx *sql.Tx, id string) (dto.GetBillByIdResponse, error)
	GetAll(ctx context.Context, tx *sql.Tx) ([]domain.Bill, error)
}

type TransactionRepositoryImpl struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		db: db,
	}
}

func (repository *TransactionRepositoryImpl) InitTransaction() (*sql.Tx, error) {
	return repository.db.Begin()
}

func (repository *TransactionRepositoryImpl) InsertBill(ctx context.Context, tx *sql.Tx, bill domain.TxBill) (domain.TxBill, error) {
	fmt.Println("insert bill", bill)

	SQL := "INSERT INTO tx_bill (bill_date, entry_date, finish_date, employee_id, customer_id, total_bill) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	err := tx.QueryRowContext(ctx, SQL, bill.BillDate, bill.EntryDate, bill.FinishDate, bill.EmployeeId, bill.CustomerId, bill.TotalBill).Scan(&bill.Id)
	if err != nil {
		return domain.TxBill{}, err
	}

	fmt.Println("result insert bill", bill)
	return bill, nil
}

func (repository *TransactionRepositoryImpl) InsertBillDetail(ctx context.Context, tx *sql.Tx, billDetail []domain.TxBillDetail) ([]domain.TxBillDetail, error) {
	var billDetails []domain.TxBillDetail
	for _, item := range billDetail {
		SQL := "insert into tx_bill_detail(bill_id, product_id, quantity, product_price) values ($1, $2, $3, $4) returning id"

		fmt.Println("billId: ", item.BillId)
		err := tx.QueryRowContext(ctx, SQL, item.BillId, item.ProductId, item.Quantity, item.ProductPrice).Scan(&item.Id)
		fmt.Println(err)

		if err != nil {
			return nil, err
		}

		billDetails = append(billDetails, item)
	}

	return billDetails, nil
}

func (repository *TransactionRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id string) (dto.GetBillByIdResponse, error) {
	var (
		bill     domain.TxBill
		customer dto.CustomerResponse
		employee dto.EmployeeResponse
	)
	sqlFindBill := `
		select b.id, b.bill_date, b.entry_date, b.finish_date, b.total_bill, c.id, c.name, c.phone_number, c.address, 
		e.id, e.name, e.phone_number, e.address
		
		from tx_bill as b 
		join 
			mst_customer as c on b.customer_id = c.id
		join 
			mst_employee as e on b.employee_id = e.id
		where 
			b.id = $1;
	`

	err := tx.QueryRowContext(ctx, sqlFindBill, id).Scan(&bill.Id, &bill.BillDate, &bill.EntryDate, &bill.FinishDate, &bill.TotalBill, &customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address, &employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
	if err != nil {
		return dto.GetBillByIdResponse{}, err
	}

	sqlFindBillDetail := `
		select bd.id, bd.bill_id, bd.quantity, bd.product_price, p.id, p.name, p.unit, p.price
		from 
			tx_bill_detail as bd
		join 
			mst_product as p on bd.product_id = p.id
		where bd.bill_id = $1
	`

	rows, err := tx.QueryContext(ctx, sqlFindBillDetail, id)
	if err != nil {
		return dto.GetBillByIdResponse{}, err
	}

	var billDetail []dto.GetBillDetailByIdResponse
	for rows.Next() {
		var (
			detail  dto.GetBillDetailByIdResponse
			product domain.Product
		)

		err := rows.Scan(&detail.Id, &detail.BillId, &detail.Quantity, &detail.ProductPrice, &product.Id, &product.Name, &product.Unit, &product.Price)
		if err != nil {
			return dto.GetBillByIdResponse{}, err
		}

		productId := strconv.Itoa(product.Id)
		detail.Product = dto.ProductResponse{
			Id:    productId,
			Name:  product.Name,
			Unit:  product.Unit,
			Price: product.Price,
		}

		billDetail = append(billDetail, detail)
	}

	billId := strconv.Itoa(bill.Id)
	billDate := bill.BillDate.Format("02-01-2006")
	entryDate := bill.EntryDate.Format("02-01-2006")
	finishDate := bill.FinishDate.Format("02-01-2006")

	billResponse := dto.GetBillByIdResponse{
		Id:         billId,
		BillDate:   billDate,
		EntryDate:  entryDate,
		FinishDate: finishDate,
		Employee:   employee,
		Customer:   customer,
		BillDetail: billDetail,
		TotalBill:  bill.TotalBill,
	}

	return billResponse, nil

}

func (repository *TransactionRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) ([]domain.Bill, error) {
	sqlGetAllBill := `
	select b.id, b.bill_date, b.entry_date, b.finish_date, b.total_bill, c.id, c.name, c.phone_number, c.address, 
	e.id, e.name, e.phone_number, e.address
	
	from tx_bill as b 
	join 
		mst_customer as c on b.customer_id = c.id
	join 
		mst_employee as e on b.employee_id = e.id
	`

	rowsBill, err := tx.QueryContext(ctx, sqlGetAllBill)
	if err != nil {
		fmt.Println("error query rowsBill: ", err)
		return nil, err
	}

	var bills []domain.Bill
	for rowsBill.Next() {
		var bill domain.Bill
		var customer domain.Customer
		var employee domain.Employee

		err := rowsBill.Scan(&bill.Id, &bill.BillDate, &bill.EntryDate, &bill.FinishDate, &bill.TotalBill, &customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address, &employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)

		if err != nil {
			return nil, err
		}

		bill = domain.Bill{
			Id:         bill.Id,
			BillDate:   bill.BillDate,
			EntryDate:  bill.EntryDate,
			FinishDate: bill.FinishDate,
			Employee:   employee,
			Customer:   customer,
			TotalBill:  bill.TotalBill,
		}

		bills = append(bills, bill)
	}

	for i, bill := range bills {
		sqlBillDetail := `
					select bd.id, bd.bill_id, bd.quantity, bd.product_price, p.id, p.name, p.unit, p.price
					from 
						tx_bill_detail as bd 
					join 
						mst_product as p on bd.product_id = p.id
					where bd.bill_id = $1
				`

		rowsBillDetail, err := tx.QueryContext(ctx, sqlBillDetail, bill.Id)
		if err != nil {
			return nil, err
		}

		var billDetails []domain.BillDetail
		for rowsBillDetail.Next() {
			var (
				billDetail domain.BillDetail
				product    domain.Product
			)

			err := rowsBillDetail.Scan(&billDetail.Id, &billDetail.BillId, &billDetail.Quantity, &billDetail.ProductPrice, &product.Id, &product.Name, &product.Unit, &product.Price)
			if err != nil {
				return nil, err
			}

			billDetail.Product = product

			billDetails = append(billDetails, billDetail)
		}

		bills[i].BillDetail = billDetails
	}

	return bills, nil
}
