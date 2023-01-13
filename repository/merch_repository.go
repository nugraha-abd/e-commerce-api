package repository

import (
	"context"
	"database/sql"
	"e-commerce-api/model/domain"
)

type MerchRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Merch
	FindById(ctx context.Context, tx *sql.Tx, merchId int) (domain.Merch, error)
	Create(ctx context.Context, tx *sql.Tx, merch domain.Merch) domain.Merch
	Update(ctx context.Context, tx *sql.Tx, merch domain.Merch) domain.Merch
	Delete(ctx context.Context, tx *sql.Tx, merch domain.Merch) 
	Order(ctx context.Context, tx *sql.Tx, merch domain.Merch) domain.Merch
}