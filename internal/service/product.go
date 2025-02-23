package service

import (
	"context"
	"super-indo-be/internal/dto"
	"super-indo-be/internal/errorcustom"
	"super-indo-be/internal/payload"
	"super-indo-be/internal/repository"
)

type IProductService interface {
	Create(ctx context.Context, p payload.CreateProductRequest) (res payload.CreateProductResponse, err error)
	GetAll(ctx context.Context, p payload.GetProductListRequest) (res []payload.GetProductListResponse, totalData int64, err error)
	GetByID(ctx context.Context, id uint64) (res payload.GetProductDetailResponse, err error)
}

type product struct {
	ProductRepository repository.IProductRepository
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &product{
		ProductRepository: productRepository,
	}
}

// Create implements IProductService.
func (s *product) Create(ctx context.Context, req payload.CreateProductRequest) (res payload.CreateProductResponse, err error) {
	productModel := req.ToModel()

	id, err := s.ProductRepository.Create(ctx, productModel)
	if err != nil {
		return res, err
	}

	res = payload.CreateProductResponse{
		ID:          id,
		Name:        productModel.Name,
		Description: productModel.Description,
		Price:       productModel.Price,
		CategoryID:  productModel.CategoryID,
		Image:       productModel.Image,
		Stock:       productModel.Stock,
		CreatedAt:   productModel.CreatedAt,
		UpdatedAt:   productModel.UpdatedAt,
	}

	return res, nil
}

// GetAll implements IProductService.
func (s *product) GetAll(ctx context.Context, req payload.GetProductListRequest) (res []payload.GetProductListResponse, totalData int64, err error) {
	products, totalData, err := s.ProductRepository.GetAll(ctx, req)
	if err != nil {
		return res, 0, err
	}

	res = dto.ProductModelListToProductListResponse(products)

	return res, totalData, nil
}

// GetByID implements IProductService.
func (s *product) GetByID(ctx context.Context, id uint64) (res payload.GetProductDetailResponse, err error) {
	product, err := s.ProductRepository.GetByID(ctx, id)
	if err != nil {
		return res, err
	}

	if product == nil {
		return res, errorcustom.ErrProductNotFound
	}

	res = dto.ProductModelToProductDetailResponse(*product)

	return res, nil
}
