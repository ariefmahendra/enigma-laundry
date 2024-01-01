package repository

import (
	"context"
	"database/sql"
	"enigma-laundry/internal/model/domain"
)

type ProductRepository interface {
	Insert(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (domain.Product, error)
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByName(ctx context.Context, name string) ([]domain.Product, error)
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (repository *ProductRepositoryImpl) Insert(ctx context.Context, product domain.Product) (domain.Product, error) {
	SQL := "insert into mst_product (name, unit, price) values ($1, $2, $3) returning id"

	err := repository.db.QueryRowContext(ctx, SQL, product.Name, product.Unit, product.Price).Scan(&product.Id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	SQL := "update mst_product set name = $1, unit = $2, price = $3 where id = $4"

	_, err := repository.db.ExecContext(ctx, SQL, product.Name, product.Unit, product.Price, product.Id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, id int) error {
	SQL := "delete from mst_product where id = $1"

	_, err := repository.db.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, id int) (domain.Product, error) {
	SQL := "select id, name, unit, price from mst_product where id = $1"

	var product domain.Product
	err := repository.db.QueryRowContext(ctx, SQL, id).Scan(&product.Id, &product.Name, &product.Unit, &product.Price)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) ([]domain.Product, error) {
	SQL := "select id, name, unit, price from mst_product"

	rows, err := repository.db.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Unit, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) FindByName(ctx context.Context, name string) ([]domain.Product, error) {
	SQL := "select id, name, unit, price from mst_product where name = $1"

	rows, err := repository.db.QueryContext(ctx, SQL, name)
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Unit, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
