package service

import (
	"context"
	"super-indo-be/internal/dto"
	"super-indo-be/internal/errorcustom"
	"super-indo-be/internal/payload"
	"super-indo-be/internal/repository"

	log "github.com/sirupsen/logrus"
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

func (s *product) Create(ctx context.Context, req payload.CreateProductRequest) (res payload.CreateProductResponse, err error) {
	productModel := req.ToModel()

	id, err := s.ProductRepository.Create(ctx, productModel)
	if err != nil {
		log.Errorf("error create product: %v", err)
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
	}

	return res, nil
}

func (s *product) GetAll(ctx context.Context, req payload.GetProductListRequest) (res []payload.GetProductListResponse, totalData int64, err error) {
	products, totalData, err := s.ProductRepository.GetAll(ctx, req)
	if err != nil {
		log.Errorf("error get all product: %v", err)
		return res, 0, err
	}

	res = dto.ProductModelListToProductListResponse(products)

	return res, totalData, nil
}

func (s *product) GetByID(ctx context.Context, id uint64) (res payload.GetProductDetailResponse, err error) {
	product, err := s.ProductRepository.GetByID(ctx, id)
	if err != nil {
		log.Errorf("error get product by id: %v", err)
		return res, err
	}

	if product == nil {
		return res, errorcustom.ErrProductNotFound
	}

	res = dto.ProductModelToProductDetailResponse(*product)

	return res, nil
}
