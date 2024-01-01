package repository

import (
	"context"
	"database/sql"
	"enigma-laundry/helper"
	"enigma-laundry/internal/model/domain"
)

type CustomerRepository interface {
	Insert(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	Update(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, customerId int) (domain.Customer, error)
	FindAll(ctx context.Context) ([]domain.Customer, error)
}

type CustomerRepositoryImpl struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &CustomerRepositoryImpl{
		db: db,
	}
}

func (repository *CustomerRepositoryImpl) Insert(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	SQL := "insert into mst_customer (name, phone_number, address) values ($1, $2, $3) returning id"

	err := repository.db.QueryRowContext(ctx, SQL, customer.Name, customer.PhoneNumber, customer.Address).Scan(&customer.Id)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	SQL := "update mst_customer set name = $1, phone_number = $2, address = $3 where id = $4"

	result, err := repository.db.ExecContext(ctx, SQL, customer.Name, customer.PhoneNumber, customer.Address, customer.Id)
	helper.PanicIfError(err)

	_, err = result.RowsAffected()
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, id int) error {
	SQL := "delete from mst_customer where id = $1"

	result, err := repository.db.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, customerId int) (domain.Customer, error) {
	SQL := "select id, name, phone_number, address from mst_customer where id = $1"

	var customer domain.Customer
	err := repository.db.QueryRowContext(ctx, SQL, customerId).Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context) ([]domain.Customer, error) {
	SQL := "select id, name, phone_number, address from mst_customer"

	row, err := repository.db.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := row.Close()
		helper.PanicIfError(err)
	}()

	var customers []domain.Customer
	for row.Next() {
		var customer domain.Customer
		err := row.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}
