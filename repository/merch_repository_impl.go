package repository

import (
	"context"
	"database/sql"
	"e-commerce-api/model/domain"
	"e-commerce-api/utils"
	"errors"
)

type MerchRepositoryImpl struct {
}

func NewMerchRepository() MerchRepository {
	return &MerchRepositoryImpl{}
}

func (repository *MerchRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Merch {
	SQL := "select seller_id, seller_name, description, images, price, quantity from merch inner join seller on merch.seller_id=seller.seller_id"
	rows, err := tx.QueryContext(ctx, SQL)
	utils.PanicIfError(err)

	var merchs []domain.Merch
	for rows.Next() {
		merch := domain.Merch{}
		err := rows.Scan(&merch.Id, &merch.Name, &merch.Description, &merch.Images, &merch.Price, &merch.Quantity)
		utils.PanicIfError(err)
		merchs = append(merchs, merch)
	}

	return merchs
}

func (repository *MerchRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, merchId int) (domain.Merch, error) {
	SQL := "select seller_id, seller_name, description, images, price, quantity from merch where seller_id = ? inner join seller on merch.seller_id=seller.seller_id"
	rows, err := tx.QueryContext(ctx, SQL, merchId)
	utils.PanicIfError(err)

	merch := domain.Merch{}
	if rows.Next() {
		err := rows.Scan(&merch.Id, &merch.Name, &merch.Description, &merch.Images, &merch.Price, &merch.Quantity)
		utils.PanicIfError(err)
		return merch, nil
	} else {
		return merch, errors.New("merch not found")
	}
}

func (repository *MerchRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, merch domain.Merch) domain.Merch {
	SQL := "insert into merch(merch_name, description, images, price, quantity) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, merch.Name, merch.Description, merch.Images, merch.Price, merch.Quantity)
	utils.PanicIfError(err)

	id, err := result.LastInsertId()
	utils.PanicIfError(err)

	merch.Id = int(id)
	return merch
}

func (repository *MerchRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, merch domain.Merch) domain.Merch {
	SQL := "update merch set merch_name = ?, description = ?, images = ?, price = ?, quantity = ? where merch_id = ?"
	_, err := tx.ExecContext(ctx, SQL, merch.Name, merch.Description, merch.Images, merch.Price, merch.Quantity, merch.Id)
	utils.PanicIfError(err)

	return merch
}

func (repository *MerchRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, merch domain.Merch) {
	SQL := "delete from merch where merch_id = ?"
	_, err := tx.ExecContext(ctx, SQL, merch.Id)
	utils.PanicIfError(err)
}

func (repository *MerchRepositoryImpl) Order(ctx context.Context, tx *sql.Tx, merch domain.Merch) domain.Merch {
	// SQL := "delete from merch where merch_id = ?"
	// _, err := tx.ExecContext(ctx, SQL, merch.Id)
	// utils.PanicIfError(err)
	return domain.Merch{} // placeholder
}