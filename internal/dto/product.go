package dto

import (
	"super-indo-be/internal/model"
	"super-indo-be/internal/payload"
)

func ProductModelListToProductListResponse(products []model.Product) []payload.GetProductListResponse {
	res := []payload.GetProductListResponse{}
	for _, product := range products {
		res = append(res, payload.GetProductListResponse{
			ID:         product.ID,
			Name:       product.Name,
			Price:      product.Price,
			CategoryID: product.CategoryID,
			Image:      product.Image,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		})
	}
	return res
}

func ProductModelToProductDetailResponse(product model.Product) payload.GetProductDetailResponse {
	return payload.GetProductDetailResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		Image:       product.Image,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
