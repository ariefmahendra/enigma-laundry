package repository

import (
	"context"
	"database/sql"
	"enigma-laundry/internal/model/domain"
)

type EmployeeRepository interface {
	Insert(ctx context.Context, employee domain.Employee) (domain.Employee, error)
	Update(ctx context.Context, employee domain.Employee) (domain.Employee, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (domain.Employee, error)
	FindAll(ctx context.Context) ([]domain.Employee, error)
}

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &EmployeeRepositoryImpl{db: db}
}

func (repository *EmployeeRepositoryImpl) Insert(ctx context.Context, employee domain.Employee) (domain.Employee, error) {
	SQL := "insert into mst_employee (name, phone_number, address) values ($1, $2, $3) returning id"

	err := repository.db.QueryRowContext(ctx, SQL, employee.Name, employee.PhoneNumber, employee.Address).Scan(&employee.Id)
	if err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}

func (repository *EmployeeRepositoryImpl) Update(ctx context.Context, employee domain.Employee) (domain.Employee, error) {
	SQL := "update mst_employee set name = $1, phone_number = $2, address = $3 where id = $4"

	_, err := repository.db.ExecContext(ctx, SQL, employee.Name, employee.PhoneNumber, employee.Address, employee.Id)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

func (repository *EmployeeRepositoryImpl) Delete(ctx context.Context, id int) error {
	SQL := "delete from mst_employee where id = $1"

	_, err := repository.db.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *EmployeeRepositoryImpl) FindById(ctx context.Context, id int) (domain.Employee, error) {
	SQL := "select id, name, phone_number, address from mst_employee where id = $1"

	var employee domain.Employee
	err := repository.db.QueryRowContext(ctx, SQL, id).Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (repository *EmployeeRepositoryImpl) FindAll(ctx context.Context) ([]domain.Employee, error) {
	SQL := "select id, name, phone_number, address from mst_employee"

	rows, err := repository.db.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}

	var employees []domain.Employee
	for rows.Next() {
		var employee domain.Employee
		err := rows.Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
