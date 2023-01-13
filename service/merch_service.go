package service

import (
	"context"
	"e-commerce-api/model/web"
)

type MerchService interface {
	FindAll(ctx context.Context) []web.MerchResponse
	FindById(ctx context.Context, merchId int) web.MerchResponse
	Create(ctx context.Context, req web.MerchCreateRequest) web.MerchResponse
	Update(ctx context.Context, req web.MerchUpdateRequest) web.MerchResponse
	Delete(ctx context.Context, merchId int)
}