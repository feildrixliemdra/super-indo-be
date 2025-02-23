package service

import (
	"context"
	"super-indo-be/internal/errorcustom"
	"super-indo-be/internal/model"
	"super-indo-be/internal/payload"
	"super-indo-be/internal/repository"
)

type ICategoryService interface {
	Create(ctx context.Context, p payload.CreateCategoryRequest) (res payload.CreateCategoryResponse, err error)
	GetAll(ctx context.Context) ([]payload.GetCategoryListResponse, error)
}

type category struct {
	CategoryRepository repository.ICategoryRepository
}

func NewCategoryService(categoryRepository repository.ICategoryRepository) ICategoryService {
	return &category{
		CategoryRepository: categoryRepository,
	}
}

// Create implements ICategoryService.
func (c *category) Create(ctx context.Context, p payload.CreateCategoryRequest) (res payload.CreateCategoryResponse, err error) {

	existCategory, err := c.CategoryRepository.GetBy(ctx, model.Category{Code: p.Code})
	if err != nil {
		return res, err
	}

	if existCategory != nil {
		return res, errorcustom.ErrCategoryAlreadyExists
	}

	category := model.Category{
		Name:        p.Name,
		Code:        p.Code,
		Description: p.Description,
	}

	id, err := c.CategoryRepository.Create(ctx, category)
	if err != nil {
		return res, err
	}

	res = payload.CreateCategoryResponse{
		ID:          id,
		Name:        p.Name,
		Code:        p.Code,
		Description: p.Description,
	}

	return res, nil
}

// GetAll implements ICategoryService.
func (c *category) GetAll(ctx context.Context) ([]payload.GetCategoryListResponse, error) {
	res := []payload.GetCategoryListResponse{}
	category, err := c.CategoryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, c := range category {
		res = append(res, payload.GetCategoryListResponse{
			ID:          c.ID,
			Name:        c.Name,
			Code:        c.Code,
			Description: c.Description,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
		})
	}

	return res, nil
}
