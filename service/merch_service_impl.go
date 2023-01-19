package service

import (
	"context"
	"database/sql"
	"e-commerce-api/exception"
	"e-commerce-api/model/domain"
	"e-commerce-api/model/web"
	"e-commerce-api/repository"
	"e-commerce-api/utils"
)

type MerchServiceImpl struct {
	MerchRepository repository.MerchRepository
	DB *sql.DB
}

func NewMerchService(merchRespository repository.MerchRepository, DB *sql.DB) MerchService {
	return &MerchServiceImpl{
		MerchRepository: merchRespository,
		DB: DB,
	}
}

func (service *MerchServiceImpl) FindAll(ctx context.Context) []web.MerchResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	merchs := service.MerchRepository.FindAll(ctx, tx)

	return utils.ToMerchResponses(merchs)
}

func (service *MerchServiceImpl) 	FindById(ctx context.Context, merchId int) web.MerchResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	merch, err := service.MerchRepository.FindById(ctx, tx, merchId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	return utils.ToMerchResponse(merch)
}

func (service *MerchServiceImpl) 	Create(ctx context.Context, req web.MerchCreateRequest) web.MerchResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	merch := domain.Merch{
		Name: req.Name,
		Description: req.Description,
		Images: req.Images,
		Price: req.Price,
		Quantity: req.Quantity,
	}

	merch = service.MerchRepository.Create(ctx, tx, merch)

	return utils.ToMerchResponse(merch)
}

func (service *MerchServiceImpl) 	Update(ctx context.Context, req web.MerchUpdateRequest) web.MerchResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	merch, err := service.MerchRepository.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	merch = domain.Merch{
		Name: req.Name,
		Description: req.Description,
		Images: req.Images,
		Price: req.Price,
		Quantity: req.Quantity,
	}

	merch = service.MerchRepository.Update(ctx, tx, merch)

	return utils.ToMerchResponse(merch)
}

func (service *MerchServiceImpl) 	Delete(ctx context.Context, merchId int) {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	merch, err := service.MerchRepository.FindById(ctx, tx, merchId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MerchRepository.Delete(ctx, tx, merch)
}