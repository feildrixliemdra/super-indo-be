package payload

import (
	"super-indo-be/internal/model"
	"time"
)

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint64 `json:"price" binding:"required"`
	CategoryID  uint64 `json:"category_id" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Stock       uint64 `json:"stock" binding:"required"`
}

func (c *CreateProductRequest) ToModel() model.Product {
	return model.Product{
		Name:        c.Name,
		Description: c.Description,
		Price:       c.Price,
		CategoryID:  c.CategoryID,
		Image:       c.Image,
		Stock:       c.Stock,
	}
}

type CreateProductResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
	CategoryID  uint64 `json:"category_id"`
	Image       string `json:"image"`
	Stock       uint64 `json:"stock"`
}

type GetProductListRequest struct {
	Limit uint64 `form:"limit"`
	Page  uint64 `form:"page"`
}

type GetProductListResponse struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Price      uint64    `json:"price"`
	CategoryID uint64    `json:"category_id"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetProductDetailResponse struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       uint64    `json:"price"`
	CategoryID  uint64    `json:"category_id"`
	Image       string    `json:"image"`
	Stock       uint64    `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
