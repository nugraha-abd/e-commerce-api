package utils

import (
	"e-commerce-api/model/domain"
	"e-commerce-api/model/web"
)

func ToMerchResponse(Merch domain.Merch) web.MerchResponse {
	return web.MerchResponse{
		Id: Merch.Id,
		Name: Merch.Name,
		Description: Merch.Description,
		Images: Merch.Images,
		Price: Merch.Price,
		Quantity: Merch.Quantity,
	}
}

func ToMerchResponses(merchs []domain.Merch) []web.MerchResponse {
	var MerchResponses []web.MerchResponse
	for _, merch := range merchs {
		MerchResponses = append(MerchResponses, ToMerchResponse(merch))
	}
	return MerchResponses
}